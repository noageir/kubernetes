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

package v2alpha1

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJob) DeepCopyInto(out *CronJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJob.
func (in *CronJob) DeepCopy() *CronJob {
	if in == nil {
		return nil
	}
	out := new(CronJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CronJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobList) DeepCopyInto(out *CronJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CronJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobList.
func (in *CronJobList) DeepCopy() *CronJobList {
	if in == nil {
		return nil
	}
	out := new(CronJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CronJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobManualRequest) DeepCopyInto(out *CronJobManualRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.CreatedJob.DeepCopyInto(&out.CreatedJob)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobManualRequest.
func (in *CronJobManualRequest) DeepCopy() *CronJobManualRequest {
	if in == nil {
		return nil
	}
	out := new(CronJobManualRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CronJobManualRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobSpec) DeepCopyInto(out *CronJobSpec) {
	*out = *in
	if in.StartingDeadlineSeconds != nil {
		in, out := &in.StartingDeadlineSeconds, &out.StartingDeadlineSeconds
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	in.JobTemplate.DeepCopyInto(&out.JobTemplate)
	if in.SuccessfulJobsHistoryLimit != nil {
		in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.FailedJobsHistoryLimit != nil {
		in, out := &in.FailedJobsHistoryLimit, &out.FailedJobsHistoryLimit
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobSpec.
func (in *CronJobSpec) DeepCopy() *CronJobSpec {
	if in == nil {
		return nil
	}
	out := new(CronJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobStatus) DeepCopyInto(out *CronJobStatus) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = make([]v1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.LastScheduleTime != nil {
		in, out := &in.LastScheduleTime, &out.LastScheduleTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobStatus.
func (in *CronJobStatus) DeepCopy() *CronJobStatus {
	if in == nil {
		return nil
	}
	out := new(CronJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTemplate) DeepCopyInto(out *JobTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTemplate.
func (in *JobTemplate) DeepCopy() *JobTemplate {
	if in == nil {
		return nil
	}
	out := new(JobTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTemplateSpec) DeepCopyInto(out *JobTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTemplateSpec.
func (in *JobTemplateSpec) DeepCopy() *JobTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(JobTemplateSpec)
	in.DeepCopyInto(out)
	return out
}
