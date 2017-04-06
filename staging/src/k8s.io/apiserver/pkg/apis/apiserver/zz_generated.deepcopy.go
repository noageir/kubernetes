// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package apiserver

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *AdmissionConfiguration) DeepCopyInto(out *AdmissionConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]AdmissionPluginConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new AdmissionConfiguration.
func (x *AdmissionConfiguration) DeepCopy() *AdmissionConfiguration {
	if x == nil {
		return nil
	}
	out := new(AdmissionConfiguration)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *AdmissionConfiguration) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *AdmissionPluginConfiguration) DeepCopyInto(out *AdmissionPluginConfiguration) {
	*out = *in
	if in.Configuration == nil {
		out.Configuration = nil
	} else {
		out.Configuration = in.Configuration.DeepCopyObject()
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new AdmissionPluginConfiguration.
func (x *AdmissionPluginConfiguration) DeepCopy() *AdmissionPluginConfiguration {
	if x == nil {
		return nil
	}
	out := new(AdmissionPluginConfiguration)
	x.DeepCopyInto(out)
	return out
}
