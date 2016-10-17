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

package forgiveness

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/golang/glog"
	"k8s.io/kubernetes/pkg/admission"
	"k8s.io/kubernetes/pkg/api"
	apierrors "k8s.io/kubernetes/pkg/api/errors"
	"k8s.io/kubernetes/pkg/api/unversioned"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
)

func init() {
	admission.RegisterPlugin("DefaultForgivenessSeconds", func(client clientset.Interface, config io.Reader) (admission.Interface, error) {
		return NewDefaultForgivenessSeconds(client), nil
	})
}

// plugin contains the client used by the admission controller
// It will add default forgiveness tolerations for every pod
// that tolerate taints `notready:NoExecute` and `unreachable:NoExecute`,
// with forgiveness period of 5 minutes.
// If the pod already specifies a toleration for taint `notready:NoExecute`
// or `unreachable:NoExecute`, the plugin won't touch it.
type plugin struct {
	*admission.Handler
	client clientset.Interface
}

var defaultForgivenessSeconds int64 = 300

// NewDefaultForgivenessSeconds creates a new instance of the DefaultForgivenessSeconds admission controller
func NewDefaultForgivenessSeconds(client clientset.Interface) admission.Interface {
	return &plugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
		client:  client,
	}
}

func (p *plugin) Admit(attributes admission.Attributes) (err error) {
	if attributes.GetResource().GroupResource() != api.Resource("pods") {
		return nil
	}
	pod, ok := attributes.GetObject().(*api.Pod)
	if !ok {
		return apierrors.NewBadRequest("Resource was marked with kind Pod but was unable to be converted")
	}

	tolerations, err := api.GetTolerationsFromPodAnnotations(pod.Annotations)
	if err != nil {
		glog.V(5).Infof("Invalid Affinity detected, but we will leave handling of this to validation phase")
		return nil
	}

	forgiveNotready := false
	forgiveUnreachable := false
	for _, toleration := range tolerations {
		if toleration.Key == unversioned.TaintNodeNotReady && toleration.Effect == api.TaintEffectNoExecute {
			forgiveNotready = true
		}

		if toleration.Key == unversioned.TaintNodeUnreachable && toleration.Effect == api.TaintEffectNoExecute {
			forgiveUnreachable = true
		}
	}

	if !forgiveNotready {
		tolerations = append(tolerations, api.Toleration{
			Key:                unversioned.TaintNodeNotReady,
			Operator:           api.TolerationOpExists,
			Effect:             api.TaintEffectNoExecute,
			ForgivenessSeconds: &defaultForgivenessSeconds,
		})
	}

	if !forgiveUnreachable {
		tolerations = append(tolerations, api.Toleration{
			Key:                unversioned.TaintNodeUnreachable,
			Operator:           api.TolerationOpExists,
			Effect:             api.TaintEffectNoExecute,
			ForgivenessSeconds: &defaultForgivenessSeconds,
		})
	}

	if len(pod.Annotations) == 0 {
		pod.Annotations = map[string]string{}
	}

	tolerationsData, err := json.Marshal(tolerations)
	if err != nil {
		return apierrors.NewForbidden(attributes.GetResource().GroupResource(), pod.Name,
			fmt.Errorf("failed to add default forgiveness tolerations for taints `notready:NoExecute` and `unreachable:NoExecute`, err: %v", err))
	}

	pod.Annotations[api.TolerationsAnnotationKey] = string(tolerationsData)
	return nil
}
