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

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/kubernetes/pkg/api/v1"
)

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new Cluster.
func (x *Cluster) DeepCopy() *Cluster {
	if x == nil {
		return nil
	}
	out := new(Cluster)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *Cluster) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new ClusterCondition.
func (x *ClusterCondition) DeepCopy() *ClusterCondition {
	if x == nil {
		return nil
	}
	out := new(ClusterCondition)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new ClusterList.
func (x *ClusterList) DeepCopy() *ClusterList {
	if x == nil {
		return nil
	}
	out := new(ClusterList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *ClusterList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.ServerAddressByClientCIDRs != nil {
		in, out := &in.ServerAddressByClientCIDRs, &out.ServerAddressByClientCIDRs
		*out = make([]ServerAddressByClientCIDR, len(*in))
		copy(*out, *in)
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.LocalObjectReference)
			**out = **in
		}
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new ClusterSpec.
func (x *ClusterSpec) DeepCopy() *ClusterSpec {
	if x == nil {
		return nil
	}
	out := new(ClusterSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new ClusterStatus.
func (x *ClusterStatus) DeepCopy() *ClusterStatus {
	if x == nil {
		return nil
	}
	out := new(ClusterStatus)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *ServerAddressByClientCIDR) DeepCopyInto(out *ServerAddressByClientCIDR) {
	*out = *in
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new ServerAddressByClientCIDR.
func (x *ServerAddressByClientCIDR) DeepCopy() *ServerAddressByClientCIDR {
	if x == nil {
		return nil
	}
	out := new(ServerAddressByClientCIDR)
	x.DeepCopyInto(out)
	return out
}
