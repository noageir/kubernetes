// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package extensions

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedFlexVolume) DeepCopyInto(out *AllowedFlexVolume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedFlexVolume.
func (in *AllowedFlexVolume) DeepCopy() *AllowedFlexVolume {
	if in == nil {
		return nil
	}
	out := new(AllowedFlexVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedHostPath) DeepCopyInto(out *AllowedHostPath) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedHostPath.
func (in *AllowedHostPath) DeepCopy() *AllowedHostPath {
	if in == nil {
		return nil
	}
	out := new(AllowedHostPath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomMetricCurrentStatus) DeepCopyInto(out *CustomMetricCurrentStatus) {
	*out = *in
	out.CurrentValue = in.CurrentValue.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomMetricCurrentStatus.
func (in *CustomMetricCurrentStatus) DeepCopy() *CustomMetricCurrentStatus {
	if in == nil {
		return nil
	}
	out := new(CustomMetricCurrentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomMetricCurrentStatusList) DeepCopyInto(out *CustomMetricCurrentStatusList) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomMetricCurrentStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomMetricCurrentStatusList.
func (in *CustomMetricCurrentStatusList) DeepCopy() *CustomMetricCurrentStatusList {
	if in == nil {
		return nil
	}
	out := new(CustomMetricCurrentStatusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomMetricTarget) DeepCopyInto(out *CustomMetricTarget) {
	*out = *in
	out.TargetValue = in.TargetValue.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomMetricTarget.
func (in *CustomMetricTarget) DeepCopy() *CustomMetricTarget {
	if in == nil {
		return nil
	}
	out := new(CustomMetricTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomMetricTargetList) DeepCopyInto(out *CustomMetricTargetList) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomMetricTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomMetricTargetList.
func (in *CustomMetricTargetList) DeepCopy() *CustomMetricTargetList {
	if in == nil {
		return nil
	}
	out := new(CustomMetricTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSet) DeepCopyInto(out *DaemonSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSet.
func (in *DaemonSet) DeepCopy() *DaemonSet {
	if in == nil {
		return nil
	}
	out := new(DaemonSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DaemonSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetCondition) DeepCopyInto(out *DaemonSetCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetCondition.
func (in *DaemonSetCondition) DeepCopy() *DaemonSetCondition {
	if in == nil {
		return nil
	}
	out := new(DaemonSetCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetList) DeepCopyInto(out *DaemonSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DaemonSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetList.
func (in *DaemonSetList) DeepCopy() *DaemonSetList {
	if in == nil {
		return nil
	}
	out := new(DaemonSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DaemonSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetSpec) DeepCopyInto(out *DaemonSetSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	in.UpdateStrategy.DeepCopyInto(&out.UpdateStrategy)
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetSpec.
func (in *DaemonSetSpec) DeepCopy() *DaemonSetSpec {
	if in == nil {
		return nil
	}
	out := new(DaemonSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetStatus) DeepCopyInto(out *DaemonSetStatus) {
	*out = *in
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]DaemonSetCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetStatus.
func (in *DaemonSetStatus) DeepCopy() *DaemonSetStatus {
	if in == nil {
		return nil
	}
	out := new(DaemonSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetUpdateStrategy) DeepCopyInto(out *DaemonSetUpdateStrategy) {
	*out = *in
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		if *in == nil {
			*out = nil
		} else {
			*out = new(RollingUpdateDaemonSet)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetUpdateStrategy.
func (in *DaemonSetUpdateStrategy) DeepCopy() *DaemonSetUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(DaemonSetUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Deployment) DeepCopyInto(out *Deployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Deployment.
func (in *Deployment) DeepCopy() *Deployment {
	if in == nil {
		return nil
	}
	out := new(Deployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Deployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentCondition) DeepCopyInto(out *DeploymentCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentCondition.
func (in *DeploymentCondition) DeepCopy() *DeploymentCondition {
	if in == nil {
		return nil
	}
	out := new(DeploymentCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentList) DeepCopyInto(out *DeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Deployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentList.
func (in *DeploymentList) DeepCopy() *DeploymentList {
	if in == nil {
		return nil
	}
	out := new(DeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentRollback) DeepCopyInto(out *DeploymentRollback) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.UpdatedAnnotations != nil {
		in, out := &in.UpdatedAnnotations, &out.UpdatedAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.RollbackTo = in.RollbackTo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentRollback.
func (in *DeploymentRollback) DeepCopy() *DeploymentRollback {
	if in == nil {
		return nil
	}
	out := new(DeploymentRollback)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentRollback) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSpec) DeepCopyInto(out *DeploymentSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	in.Strategy.DeepCopyInto(&out.Strategy)
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.RollbackTo != nil {
		in, out := &in.RollbackTo, &out.RollbackTo
		if *in == nil {
			*out = nil
		} else {
			*out = new(RollbackConfig)
			**out = **in
		}
	}
	if in.ProgressDeadlineSeconds != nil {
		in, out := &in.ProgressDeadlineSeconds, &out.ProgressDeadlineSeconds
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSpec.
func (in *DeploymentSpec) DeepCopy() *DeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStatus) DeepCopyInto(out *DeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]DeploymentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStatus.
func (in *DeploymentStatus) DeepCopy() *DeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStrategy) DeepCopyInto(out *DeploymentStrategy) {
	*out = *in
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		if *in == nil {
			*out = nil
		} else {
			*out = new(RollingUpdateDeployment)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStrategy.
func (in *DeploymentStrategy) DeepCopy() *DeploymentStrategy {
	if in == nil {
		return nil
	}
	out := new(DeploymentStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FSGroupStrategyOptions) DeepCopyInto(out *FSGroupStrategyOptions) {
	*out = *in
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]GroupIDRange, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FSGroupStrategyOptions.
func (in *FSGroupStrategyOptions) DeepCopy() *FSGroupStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(FSGroupStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupIDRange) DeepCopyInto(out *GroupIDRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupIDRange.
func (in *GroupIDRange) DeepCopy() *GroupIDRange {
	if in == nil {
		return nil
	}
	out := new(GroupIDRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPIngressPath) DeepCopyInto(out *HTTPIngressPath) {
	*out = *in
	out.Backend = in.Backend
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPIngressPath.
func (in *HTTPIngressPath) DeepCopy() *HTTPIngressPath {
	if in == nil {
		return nil
	}
	out := new(HTTPIngressPath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPIngressRuleValue) DeepCopyInto(out *HTTPIngressRuleValue) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]HTTPIngressPath, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPIngressRuleValue.
func (in *HTTPIngressRuleValue) DeepCopy() *HTTPIngressRuleValue {
	if in == nil {
		return nil
	}
	out := new(HTTPIngressRuleValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPortRange) DeepCopyInto(out *HostPortRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPortRange.
func (in *HostPortRange) DeepCopy() *HostPortRange {
	if in == nil {
		return nil
	}
	out := new(HostPortRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingress) DeepCopyInto(out *Ingress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingress.
func (in *Ingress) DeepCopy() *Ingress {
	if in == nil {
		return nil
	}
	out := new(Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ingress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressBackend) DeepCopyInto(out *IngressBackend) {
	*out = *in
	out.ServicePort = in.ServicePort
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressBackend.
func (in *IngressBackend) DeepCopy() *IngressBackend {
	if in == nil {
		return nil
	}
	out := new(IngressBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressList) DeepCopyInto(out *IngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ingress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressList.
func (in *IngressList) DeepCopy() *IngressList {
	if in == nil {
		return nil
	}
	out := new(IngressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRule) DeepCopyInto(out *IngressRule) {
	*out = *in
	in.IngressRuleValue.DeepCopyInto(&out.IngressRuleValue)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRule.
func (in *IngressRule) DeepCopy() *IngressRule {
	if in == nil {
		return nil
	}
	out := new(IngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRuleValue) DeepCopyInto(out *IngressRuleValue) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		if *in == nil {
			*out = nil
		} else {
			*out = new(HTTPIngressRuleValue)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRuleValue.
func (in *IngressRuleValue) DeepCopy() *IngressRuleValue {
	if in == nil {
		return nil
	}
	out := new(IngressRuleValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressSpec) DeepCopyInto(out *IngressSpec) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		if *in == nil {
			*out = nil
		} else {
			*out = new(IngressBackend)
			**out = **in
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]IngressTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]IngressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressSpec.
func (in *IngressSpec) DeepCopy() *IngressSpec {
	if in == nil {
		return nil
	}
	out := new(IngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressStatus) DeepCopyInto(out *IngressStatus) {
	*out = *in
	in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressStatus.
func (in *IngressStatus) DeepCopy() *IngressStatus {
	if in == nil {
		return nil
	}
	out := new(IngressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressTLS) DeepCopyInto(out *IngressTLS) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressTLS.
func (in *IngressTLS) DeepCopy() *IngressTLS {
	if in == nil {
		return nil
	}
	out := new(IngressTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicy) DeepCopyInto(out *PodSecurityPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicy.
func (in *PodSecurityPolicy) DeepCopy() *PodSecurityPolicy {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSecurityPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicyList) DeepCopyInto(out *PodSecurityPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodSecurityPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicyList.
func (in *PodSecurityPolicyList) DeepCopy() *PodSecurityPolicyList {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSecurityPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicySpec) DeepCopyInto(out *PodSecurityPolicySpec) {
	*out = *in
	if in.DefaultAddCapabilities != nil {
		in, out := &in.DefaultAddCapabilities, &out.DefaultAddCapabilities
		*out = make([]core.Capability, len(*in))
		copy(*out, *in)
	}
	if in.RequiredDropCapabilities != nil {
		in, out := &in.RequiredDropCapabilities, &out.RequiredDropCapabilities
		*out = make([]core.Capability, len(*in))
		copy(*out, *in)
	}
	if in.AllowedCapabilities != nil {
		in, out := &in.AllowedCapabilities, &out.AllowedCapabilities
		*out = make([]core.Capability, len(*in))
		copy(*out, *in)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]FSType, len(*in))
		copy(*out, *in)
	}
	if in.HostPorts != nil {
		in, out := &in.HostPorts, &out.HostPorts
		*out = make([]HostPortRange, len(*in))
		copy(*out, *in)
	}
	in.SELinux.DeepCopyInto(&out.SELinux)
	in.RunAsUser.DeepCopyInto(&out.RunAsUser)
	in.SupplementalGroups.DeepCopyInto(&out.SupplementalGroups)
	in.FSGroup.DeepCopyInto(&out.FSGroup)
	if in.DefaultAllowPrivilegeEscalation != nil {
		in, out := &in.DefaultAllowPrivilegeEscalation, &out.DefaultAllowPrivilegeEscalation
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.AllowedHostPaths != nil {
		in, out := &in.AllowedHostPaths, &out.AllowedHostPaths
		*out = make([]AllowedHostPath, len(*in))
		copy(*out, *in)
	}
	if in.AllowedFlexVolumes != nil {
		in, out := &in.AllowedFlexVolumes, &out.AllowedFlexVolumes
		*out = make([]AllowedFlexVolume, len(*in))
		copy(*out, *in)
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicySpec.
func (in *PodSecurityPolicySpec) DeepCopy() *PodSecurityPolicySpec {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSet) DeepCopyInto(out *ReplicaSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSet.
func (in *ReplicaSet) DeepCopy() *ReplicaSet {
	if in == nil {
		return nil
	}
	out := new(ReplicaSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicaSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSetCondition) DeepCopyInto(out *ReplicaSetCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSetCondition.
func (in *ReplicaSetCondition) DeepCopy() *ReplicaSetCondition {
	if in == nil {
		return nil
	}
	out := new(ReplicaSetCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSetList) DeepCopyInto(out *ReplicaSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicaSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSetList.
func (in *ReplicaSetList) DeepCopy() *ReplicaSetList {
	if in == nil {
		return nil
	}
	out := new(ReplicaSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicaSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSetSpec) DeepCopyInto(out *ReplicaSetSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSetSpec.
func (in *ReplicaSetSpec) DeepCopy() *ReplicaSetSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicaSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSetStatus) DeepCopyInto(out *ReplicaSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ReplicaSetCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSetStatus.
func (in *ReplicaSetStatus) DeepCopy() *ReplicaSetStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicaSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationControllerDummy) DeepCopyInto(out *ReplicationControllerDummy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationControllerDummy.
func (in *ReplicationControllerDummy) DeepCopy() *ReplicationControllerDummy {
	if in == nil {
		return nil
	}
	out := new(ReplicationControllerDummy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationControllerDummy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollbackConfig) DeepCopyInto(out *RollbackConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollbackConfig.
func (in *RollbackConfig) DeepCopy() *RollbackConfig {
	if in == nil {
		return nil
	}
	out := new(RollbackConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateDaemonSet) DeepCopyInto(out *RollingUpdateDaemonSet) {
	*out = *in
	out.MaxUnavailable = in.MaxUnavailable
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateDaemonSet.
func (in *RollingUpdateDaemonSet) DeepCopy() *RollingUpdateDaemonSet {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateDaemonSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateDeployment) DeepCopyInto(out *RollingUpdateDeployment) {
	*out = *in
	out.MaxUnavailable = in.MaxUnavailable
	out.MaxSurge = in.MaxSurge
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateDeployment.
func (in *RollingUpdateDeployment) DeepCopy() *RollingUpdateDeployment {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsUserStrategyOptions) DeepCopyInto(out *RunAsUserStrategyOptions) {
	*out = *in
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]UserIDRange, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsUserStrategyOptions.
func (in *RunAsUserStrategyOptions) DeepCopy() *RunAsUserStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(RunAsUserStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SELinuxStrategyOptions) DeepCopyInto(out *SELinuxStrategyOptions) {
	*out = *in
	if in.SELinuxOptions != nil {
		in, out := &in.SELinuxOptions, &out.SELinuxOptions
		if *in == nil {
			*out = nil
		} else {
			*out = new(core.SELinuxOptions)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SELinuxStrategyOptions.
func (in *SELinuxStrategyOptions) DeepCopy() *SELinuxStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(SELinuxStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupplementalGroupsStrategyOptions) DeepCopyInto(out *SupplementalGroupsStrategyOptions) {
	*out = *in
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]GroupIDRange, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupplementalGroupsStrategyOptions.
func (in *SupplementalGroupsStrategyOptions) DeepCopy() *SupplementalGroupsStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(SupplementalGroupsStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserIDRange) DeepCopyInto(out *UserIDRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserIDRange.
func (in *UserIDRange) DeepCopy() *UserIDRange {
	if in == nil {
		return nil
	}
	out := new(UserIDRange)
	in.DeepCopyInto(out)
	return out
}
