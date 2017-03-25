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

// Package options contains flags and options for initializing federation-apiserver.
package options

import (
	"time"

	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	"k8s.io/kubernetes/pkg/api"
	kubeoptions "k8s.io/kubernetes/pkg/kubeapiserver/options"

	// add the kubernetes feature gates
	_ "k8s.io/kubernetes/pkg/features"

	"github.com/spf13/pflag"
	"k8s.io/kubernetes/pkg/apis/componentconfig"
)

// Runtime options for the federation-apiserver.
type ServerRunOptions struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions
	Etcd                    *genericoptions.EtcdOptions
	SecureServing           *genericoptions.SecureServingOptions
	InsecureServing         *genericoptions.ServingOptions
	Audit                   *genericoptions.AuditLogOptions
	Features                *genericoptions.FeatureOptions
	Authentication          *kubeoptions.BuiltInAuthenticationOptions
	Authorization           *kubeoptions.BuiltInAuthorizationOptions
	CloudProvider           *kubeoptions.CloudProviderOptions
	StorageSerialization    *kubeoptions.StorageSerializationOptions
	APIEnablement           *kubeoptions.APIEnablementOptions

	EventTTL time.Duration
}

func (s *ServerRunOptions) ComponentConfigz() componentconfig.APIServerConfiguration {
	componentconfig := componentconfig.APIServerConfiguration{
		AdmissionControl:            s.GenericServerRunOptions.AdmissionControl,
		AdmissionControlConfigFile:  s.GenericServerRunOptions.AdmissionControlConfigFile,
		AdvertiseAddress:            s.GenericServerRunOptions.AdvertiseAddress.String(),
		CorsAllowedOriginList:       s.GenericServerRunOptions.CorsAllowedOriginList,
		ExternalHost:                s.GenericServerRunOptions.ExternalHost,
		MaxRequestsInFlight:         s.GenericServerRunOptions.MaxRequestsInFlight,
		MaxMutatingRequestsInFlight: s.GenericServerRunOptions.MaxMutatingRequestsInFlight,
		MinRequestTimeout:           s.GenericServerRunOptions.MinRequestTimeout,
		TargetRAMMB:                 s.GenericServerRunOptions.TargetRAMMB,
		WatchCacheSizes:             s.GenericServerRunOptions.WatchCacheSizes,
		EnableProfiling:             s.Features.EnableProfiling,
		EnableContentionProfiling:   s.Features.EnableContentionProfiling,
		EnableSwaggerUI:             s.Features.EnableSwaggerUI,
		StorageConfig: componentconfig.APIServerEtcdConfiguration{
			Type:       s.Etcd.StorageConfig.Type,
			ServerSize: len(s.Etcd.StorageConfig.ServerList),
			Quorum:     s.Etcd.StorageConfig.Quorum,
			DeserializationCacheSize: s.Etcd.StorageConfig.DeserializationCacheSize,
		},
		DefaultStorageMediaType: s.Etcd.DefaultStorageMediaType,
		DeleteCollectionWorkers: s.Etcd.DeleteCollectionWorkers,
		EnableGarbageCollection: s.Etcd.EnableGarbageCollection,
		EnableWatchCache:        s.Etcd.EnableWatchCache,
		CloudConfigFile:         s.CloudProvider.CloudConfigFile,
		CloudProvider:           s.CloudProvider.CloudProvider,
		StorageSerialization: componentconfig.APIServerStorageSerializationOptions{
			StorageVersions:        s.StorageSerialization.StorageVersions,
			DefaultStorageVersions: s.StorageSerialization.DefaultStorageVersions,
		},
		AllowPrivileged:           nil,
		KubernetesServiceNodePort: nil,
		MasterCount:               nil,
		MaxConnectionBytesPerSec:  nil,
		ServiceClusterIPRange:     nil,
		ServiceNodePortRange:      nil,
	}

	return componentconfig
}

// NewServerRunOptions creates a new ServerRunOptions object with default values.
func NewServerRunOptions() *ServerRunOptions {
	s := ServerRunOptions{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		Etcd:                 genericoptions.NewEtcdOptions(storagebackend.NewDefaultConfig(kubeoptions.DefaultEtcdPathPrefix, api.Scheme, nil)),
		SecureServing:        kubeoptions.NewSecureServingOptions(),
		InsecureServing:      genericoptions.NewInsecureServingOptions(),
		Audit:                genericoptions.NewAuditLogOptions(),
		Features:             genericoptions.NewFeatureOptions(),
		Authentication:       kubeoptions.NewBuiltInAuthenticationOptions().WithAll(),
		Authorization:        kubeoptions.NewBuiltInAuthorizationOptions(),
		CloudProvider:        kubeoptions.NewCloudProviderOptions(),
		StorageSerialization: kubeoptions.NewStorageSerializationOptions(),
		APIEnablement:        kubeoptions.NewAPIEnablementOptions(),

		EventTTL: 1 * time.Hour,
	}
	// Overwrite the default for storage data format.
	s.Etcd.DefaultStorageMediaType = "application/vnd.kubernetes.protobuf"
	return &s
}

// AddFlags adds flags for ServerRunOptions fields to be specified via FlagSet.
func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	// Add the generic flags.
	s.GenericServerRunOptions.AddUniversalFlags(fs)
	s.Etcd.AddFlags(fs)
	s.SecureServing.AddFlags(fs)
	s.InsecureServing.AddFlags(fs)
	s.Audit.AddFlags(fs)
	s.Features.AddFlags(fs)
	s.Authentication.AddFlags(fs)
	s.Authorization.AddFlags(fs)
	s.CloudProvider.AddFlags(fs)
	s.StorageSerialization.AddFlags(fs)
	s.APIEnablement.AddFlags(fs)

	fs.DurationVar(&s.EventTTL, "event-ttl", s.EventTTL,
		"Amount of time to retain events. Default is 1h.")
}
