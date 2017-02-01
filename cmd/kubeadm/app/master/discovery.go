/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package master

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kuberuntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	certutil "k8s.io/client-go/util/cert"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/v1"
	extensions "k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
	"k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
)

const (
	kubeDiscoverySecretName = "clusterinfo"
	kubeDiscoveryName       = "kube-discovery"
)

// TODO: Remove this as soon as jbeda's token discovery refactoring PR has merged
const KubeDiscoveryDeployment = `
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    k8s-app: kube-discovery
    kubernetes.io/cluster-service: "true"
  name: kube-discovery
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: kube-discovery
      kubernetes.io/cluster-service: "true"
  template:
    metadata:
      labels:
        k8s-app: kube-discovery
        # TODO: I guess we can remove all these cluster-service labels...
        kubernetes.io/cluster-service: "true"
      annotations:
        # TODO: Move this to the beta tolerations field below as soon as the Tolerations field exists in PodSpec
        scheduler.alpha.kubernetes.io/tolerations: '[{"key":"dedicated","value":"master","effect":"NoSchedule"}]'
    spec:
      containers:
      - name: kube-discovery
        image: {{ .ImageRepository }}/kube-discovery-{{ .Arch }}:1.0
        imagePullPolicy: IfNotPresent
        command:
        - /usr/local/bin/kube-discovery
        ports:
        - containerPort: 9898
          hostPort: 9898
          name: http
        volumeMounts:
        - mountPath: /tmp/secret
          name: clusterinfo
          readOnly: true
      hostNetwork: true
      # tolerations:
      # - key: dedicated
      #   value: master
      #   effect: NoSchedule
      securityContext:
          seLinuxOptions:
            type: spc_t
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: beta.kubernetes.io/arch
                operator: In
                values:
                - {{ .Arch }}
      volumes:
      - name: clusterinfo
        secret:
          defaultMode: 420
          secretName: clusterinfo
`

func encodeKubeDiscoverySecretData(dcfg *kubeadmapi.TokenDiscovery, apicfg kubeadmapi.API, caCert *x509.Certificate) map[string][]byte {
	var (
		data         = map[string][]byte{}
		endpointList = []string{}
		tokenMap     = map[string]string{}
	)

	for _, addr := range apicfg.AdvertiseAddresses {
		endpointList = append(endpointList, fmt.Sprintf("https://%s:%d", addr, apicfg.Port))
	}

	tokenMap[dcfg.ID] = dcfg.Secret

	data["endpoint-list.json"], _ = json.Marshal(endpointList)
	data["token-map.json"], _ = json.Marshal(tokenMap)
	data["ca.pem"] = certutil.EncodeCertPEM(caCert)

	return data
}

func CreateDiscoveryDeploymentAndSecret(cfg *kubeadmapi.MasterConfiguration, client *clientset.Clientset) error {
	caCertificatePath := path.Join(kubeadmapi.GlobalEnvParams.HostPKIPath, kubeadmconstants.CACertName)
	caCerts, err := certutil.CertsFromFile(caCertificatePath)
	if err != nil {
		return fmt.Errorf("couldn't load the CA certificate file %s: %v", caCertificatePath, err)
	}

	// We are only putting one certificate in the certificate pem file, so it's safe to just pick the first one
	// TODO: Support multiple certs here in order to be able to rotate certs
	caCert := caCerts[0]

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: kubeDiscoverySecretName},
		Type:       v1.SecretTypeOpaque,
		Data:       encodeKubeDiscoverySecretData(cfg.Discovery.Token, cfg.API, caCert),
	}
	if _, err := client.Secrets(metav1.NamespaceSystem).Create(secret); err != nil {
		return fmt.Errorf("failed to create %q secret [%v]", kubeDiscoverySecretName, err)
	}

	if err := createDiscoveryDeployment(client); err != nil {
		return err
	}

	fmt.Println("[token-discovery] Created the kube-discovery deployment, waiting for it to become ready")

	start := time.Now()
	wait.PollInfinite(kubeadmconstants.APICallRetryInterval, func() (bool, error) {
		d, err := client.Extensions().Deployments(metav1.NamespaceSystem).Get(kubeDiscoveryName, metav1.GetOptions{})
		if err != nil {
			return false, nil
		}
		if d.Status.AvailableReplicas < 1 {
			return false, nil
		}
		return true, nil
	})
	fmt.Printf("[token-discovery] kube-discovery is ready after %f seconds\n", time.Since(start).Seconds())

	return nil
}

func createDiscoveryDeployment(client *clientset.Clientset) error {
	discoveryBytes, err := kubeadmutil.ParseTemplate(KubeDiscoveryDeployment, struct{ ImageRepository, Arch string }{
		ImageRepository: kubeadmapi.GlobalEnvParams.RepositoryPrefix,
		Arch:            runtime.GOARCH,
	})
	if err != nil {
		return fmt.Errorf("error when parsing kube-discovery template: %v", err)
	}

	discoveryDeployment := &extensions.Deployment{}
	if err := kuberuntime.DecodeInto(api.Codecs.UniversalDecoder(), discoveryBytes, discoveryDeployment); err != nil {
		return fmt.Errorf("unable to decode kube-discovery deployment %v", err)
	}
	if _, err := client.ExtensionsV1beta1().Deployments(metav1.NamespaceSystem).Create(discoveryDeployment); err != nil {
		return fmt.Errorf("unable to create a new discovery deployment: %v", err)
	}
	return nil
}
