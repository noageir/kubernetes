/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package api

// AUTO-GENERATED FUNCTIONS START HERE
import (
	resource "github.com/GoogleCloudPlatform/kubernetes/pkg/api/resource"
	conversion "github.com/GoogleCloudPlatform/kubernetes/pkg/conversion"
	fields "github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
	labels "github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	runtime "github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"
	util "github.com/GoogleCloudPlatform/kubernetes/pkg/util"
	inf "speter.net/go/exp/math/dec/inf"
	time "time"
)

func deepCopy_api_AWSElasticBlockStoreVolumeSource(in AWSElasticBlockStoreVolumeSource, out *AWSElasticBlockStoreVolumeSource, c *conversion.Cloner) error {
	out.VolumeID = in.VolumeID
	out.FSType = in.FSType
	out.Partition = in.Partition
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_api_AutoScaleIntentionThresholdConfig(in AutoScaleIntentionThresholdConfig, out *AutoScaleIntentionThresholdConfig, c *conversion.Cloner) error {
	out.Intent = in.Intent
	out.Value = in.Value
	out.Duration = in.Duration
	return nil
}

func deepCopy_api_AutoScaleThreshold(in AutoScaleThreshold, out *AutoScaleThreshold, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.Intentions != nil {
		out.Intentions = make([]AutoScaleIntentionThresholdConfig, len(in.Intentions))
		for i := range in.Intentions {
			if err := deepCopy_api_AutoScaleIntentionThresholdConfig(in.Intentions[i], &out.Intentions[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Intentions = nil
	}
	out.ActionType = in.ActionType
	out.ScaleBy = in.ScaleBy
	return nil
}

func deepCopy_api_AutoScaler(in AutoScaler, out *AutoScaler, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_AutoScalerSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_api_AutoScalerStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_AutoScalerList(in AutoScalerList, out *AutoScalerList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]AutoScaler, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_AutoScaler(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_AutoScalerSpec(in AutoScalerSpec, out *AutoScalerSpec, c *conversion.Cloner) error {
	if in.Thresholds != nil {
		out.Thresholds = make([]AutoScaleThreshold, len(in.Thresholds))
		for i := range in.Thresholds {
			if err := deepCopy_api_AutoScaleThreshold(in.Thresholds[i], &out.Thresholds[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Thresholds = nil
	}
	out.MaxAutoScaleCount = in.MaxAutoScaleCount
	out.MinAutoScaleCount = in.MinAutoScaleCount
	if in.TargetSelector != nil {
		out.TargetSelector = make(map[string]string)
		for key, val := range in.TargetSelector {
			out.TargetSelector[key] = val
		}
	} else {
		out.TargetSelector = nil
	}
	if in.Advisors != nil {
		out.Advisors = make([]string, len(in.Advisors))
		for i := range in.Advisors {
			out.Advisors[i] = in.Advisors[i]
		}
	} else {
		out.Advisors = nil
	}
	return nil
}

func deepCopy_api_AutoScalerStatus(in AutoScalerStatus, out *AutoScalerStatus, c *conversion.Cloner) error {
	if err := deepCopy_api_AutoScaleThreshold(in.LastActionTrigger, &out.LastActionTrigger, c); err != nil {
		return err
	}
	if err := deepCopy_util_Time(in.LastActionTimestamp, &out.LastActionTimestamp, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_Binding(in Binding, out *Binding, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectReference(in.Target, &out.Target, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_Capabilities(in Capabilities, out *Capabilities, c *conversion.Cloner) error {
	if in.Add != nil {
		out.Add = make([]Capability, len(in.Add))
		for i := range in.Add {
			out.Add[i] = in.Add[i]
		}
	} else {
		out.Add = nil
	}
	if in.Drop != nil {
		out.Drop = make([]Capability, len(in.Drop))
		for i := range in.Drop {
			out.Drop[i] = in.Drop[i]
		}
	} else {
		out.Drop = nil
	}
	return nil
}

func deepCopy_api_ComponentCondition(in ComponentCondition, out *ComponentCondition, c *conversion.Cloner) error {
	out.Type = in.Type
	out.Status = in.Status
	out.Message = in.Message
	out.Error = in.Error
	return nil
}

func deepCopy_api_ComponentStatus(in ComponentStatus, out *ComponentStatus, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if in.Conditions != nil {
		out.Conditions = make([]ComponentCondition, len(in.Conditions))
		for i := range in.Conditions {
			if err := deepCopy_api_ComponentCondition(in.Conditions[i], &out.Conditions[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

func deepCopy_api_ComponentStatusList(in ComponentStatusList, out *ComponentStatusList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]ComponentStatus, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_ComponentStatus(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_Container(in Container, out *Container, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Image = in.Image
	if in.Command != nil {
		out.Command = make([]string, len(in.Command))
		for i := range in.Command {
			out.Command[i] = in.Command[i]
		}
	} else {
		out.Command = nil
	}
	if in.Args != nil {
		out.Args = make([]string, len(in.Args))
		for i := range in.Args {
			out.Args[i] = in.Args[i]
		}
	} else {
		out.Args = nil
	}
	out.WorkingDir = in.WorkingDir
	if in.Ports != nil {
		out.Ports = make([]ContainerPort, len(in.Ports))
		for i := range in.Ports {
			if err := deepCopy_api_ContainerPort(in.Ports[i], &out.Ports[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Ports = nil
	}
	if in.Env != nil {
		out.Env = make([]EnvVar, len(in.Env))
		for i := range in.Env {
			if err := deepCopy_api_EnvVar(in.Env[i], &out.Env[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	if err := deepCopy_api_ResourceRequirements(in.Resources, &out.Resources, c); err != nil {
		return err
	}
	if in.VolumeMounts != nil {
		out.VolumeMounts = make([]VolumeMount, len(in.VolumeMounts))
		for i := range in.VolumeMounts {
			if err := deepCopy_api_VolumeMount(in.VolumeMounts[i], &out.VolumeMounts[i], c); err != nil {
				return err
			}
		}
	} else {
		out.VolumeMounts = nil
	}
	if in.LivenessProbe != nil {
		out.LivenessProbe = new(Probe)
		if err := deepCopy_api_Probe(*in.LivenessProbe, out.LivenessProbe, c); err != nil {
			return err
		}
	} else {
		out.LivenessProbe = nil
	}
	if in.ReadinessProbe != nil {
		out.ReadinessProbe = new(Probe)
		if err := deepCopy_api_Probe(*in.ReadinessProbe, out.ReadinessProbe, c); err != nil {
			return err
		}
	} else {
		out.ReadinessProbe = nil
	}
	if in.Lifecycle != nil {
		out.Lifecycle = new(Lifecycle)
		if err := deepCopy_api_Lifecycle(*in.Lifecycle, out.Lifecycle, c); err != nil {
			return err
		}
	} else {
		out.Lifecycle = nil
	}
	out.TerminationMessagePath = in.TerminationMessagePath
	out.ImagePullPolicy = in.ImagePullPolicy
	if in.SecurityContext != nil {
		out.SecurityContext = new(SecurityContext)
		if err := deepCopy_api_SecurityContext(*in.SecurityContext, out.SecurityContext, c); err != nil {
			return err
		}
	} else {
		out.SecurityContext = nil
	}
	return nil
}

func deepCopy_api_ContainerPort(in ContainerPort, out *ContainerPort, c *conversion.Cloner) error {
	out.Name = in.Name
	out.HostPort = in.HostPort
	out.ContainerPort = in.ContainerPort
	out.Protocol = in.Protocol
	out.HostIP = in.HostIP
	return nil
}

func deepCopy_api_ContainerState(in ContainerState, out *ContainerState, c *conversion.Cloner) error {
	if in.Waiting != nil {
		out.Waiting = new(ContainerStateWaiting)
		if err := deepCopy_api_ContainerStateWaiting(*in.Waiting, out.Waiting, c); err != nil {
			return err
		}
	} else {
		out.Waiting = nil
	}
	if in.Running != nil {
		out.Running = new(ContainerStateRunning)
		if err := deepCopy_api_ContainerStateRunning(*in.Running, out.Running, c); err != nil {
			return err
		}
	} else {
		out.Running = nil
	}
	if in.Terminated != nil {
		out.Terminated = new(ContainerStateTerminated)
		if err := deepCopy_api_ContainerStateTerminated(*in.Terminated, out.Terminated, c); err != nil {
			return err
		}
	} else {
		out.Terminated = nil
	}
	return nil
}

func deepCopy_api_ContainerStateRunning(in ContainerStateRunning, out *ContainerStateRunning, c *conversion.Cloner) error {
	if err := deepCopy_util_Time(in.StartedAt, &out.StartedAt, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_ContainerStateTerminated(in ContainerStateTerminated, out *ContainerStateTerminated, c *conversion.Cloner) error {
	out.ExitCode = in.ExitCode
	out.Signal = in.Signal
	out.Reason = in.Reason
	out.Message = in.Message
	if err := deepCopy_util_Time(in.StartedAt, &out.StartedAt, c); err != nil {
		return err
	}
	if err := deepCopy_util_Time(in.FinishedAt, &out.FinishedAt, c); err != nil {
		return err
	}
	out.ContainerID = in.ContainerID
	return nil
}

func deepCopy_api_ContainerStateWaiting(in ContainerStateWaiting, out *ContainerStateWaiting, c *conversion.Cloner) error {
	out.Reason = in.Reason
	return nil
}

func deepCopy_api_ContainerStatus(in ContainerStatus, out *ContainerStatus, c *conversion.Cloner) error {
	out.Name = in.Name
	if err := deepCopy_api_ContainerState(in.State, &out.State, c); err != nil {
		return err
	}
	if err := deepCopy_api_ContainerState(in.LastTerminationState, &out.LastTerminationState, c); err != nil {
		return err
	}
	out.Ready = in.Ready
	out.RestartCount = in.RestartCount
	out.Image = in.Image
	out.ImageID = in.ImageID
	out.ContainerID = in.ContainerID
	return nil
}

func deepCopy_api_DeleteOptions(in DeleteOptions, out *DeleteOptions, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if in.GracePeriodSeconds != nil {
		out.GracePeriodSeconds = new(int64)
		*out.GracePeriodSeconds = *in.GracePeriodSeconds
	} else {
		out.GracePeriodSeconds = nil
	}
	return nil
}

func deepCopy_api_EmptyDirVolumeSource(in EmptyDirVolumeSource, out *EmptyDirVolumeSource, c *conversion.Cloner) error {
	out.Medium = in.Medium
	return nil
}

func deepCopy_api_EndpointAddress(in EndpointAddress, out *EndpointAddress, c *conversion.Cloner) error {
	out.IP = in.IP
	if in.TargetRef != nil {
		out.TargetRef = new(ObjectReference)
		if err := deepCopy_api_ObjectReference(*in.TargetRef, out.TargetRef, c); err != nil {
			return err
		}
	} else {
		out.TargetRef = nil
	}
	return nil
}

func deepCopy_api_EndpointPort(in EndpointPort, out *EndpointPort, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Port = in.Port
	out.Protocol = in.Protocol
	return nil
}

func deepCopy_api_EndpointSubset(in EndpointSubset, out *EndpointSubset, c *conversion.Cloner) error {
	if in.Addresses != nil {
		out.Addresses = make([]EndpointAddress, len(in.Addresses))
		for i := range in.Addresses {
			if err := deepCopy_api_EndpointAddress(in.Addresses[i], &out.Addresses[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Addresses = nil
	}
	if in.Ports != nil {
		out.Ports = make([]EndpointPort, len(in.Ports))
		for i := range in.Ports {
			if err := deepCopy_api_EndpointPort(in.Ports[i], &out.Ports[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Ports = nil
	}
	return nil
}

func deepCopy_api_Endpoints(in Endpoints, out *Endpoints, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if in.Subsets != nil {
		out.Subsets = make([]EndpointSubset, len(in.Subsets))
		for i := range in.Subsets {
			if err := deepCopy_api_EndpointSubset(in.Subsets[i], &out.Subsets[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Subsets = nil
	}
	return nil
}

func deepCopy_api_EndpointsList(in EndpointsList, out *EndpointsList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]Endpoints, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_Endpoints(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_EnvVar(in EnvVar, out *EnvVar, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Value = in.Value
	if in.ValueFrom != nil {
		out.ValueFrom = new(EnvVarSource)
		if err := deepCopy_api_EnvVarSource(*in.ValueFrom, out.ValueFrom, c); err != nil {
			return err
		}
	} else {
		out.ValueFrom = nil
	}
	return nil
}

func deepCopy_api_EnvVarSource(in EnvVarSource, out *EnvVarSource, c *conversion.Cloner) error {
	if in.FieldRef != nil {
		out.FieldRef = new(ObjectFieldSelector)
		if err := deepCopy_api_ObjectFieldSelector(*in.FieldRef, out.FieldRef, c); err != nil {
			return err
		}
	} else {
		out.FieldRef = nil
	}
	return nil
}

func deepCopy_api_Event(in Event, out *Event, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectReference(in.InvolvedObject, &out.InvolvedObject, c); err != nil {
		return err
	}
	out.Reason = in.Reason
	out.Message = in.Message
	if err := deepCopy_api_EventSource(in.Source, &out.Source, c); err != nil {
		return err
	}
	if err := deepCopy_util_Time(in.FirstTimestamp, &out.FirstTimestamp, c); err != nil {
		return err
	}
	if err := deepCopy_util_Time(in.LastTimestamp, &out.LastTimestamp, c); err != nil {
		return err
	}
	out.Count = in.Count
	return nil
}

func deepCopy_api_EventList(in EventList, out *EventList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]Event, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_Event(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_EventSource(in EventSource, out *EventSource, c *conversion.Cloner) error {
	out.Component = in.Component
	out.Host = in.Host
	return nil
}

func deepCopy_api_ExecAction(in ExecAction, out *ExecAction, c *conversion.Cloner) error {
	if in.Command != nil {
		out.Command = make([]string, len(in.Command))
		for i := range in.Command {
			out.Command[i] = in.Command[i]
		}
	} else {
		out.Command = nil
	}
	return nil
}

func deepCopy_api_GCEPersistentDiskVolumeSource(in GCEPersistentDiskVolumeSource, out *GCEPersistentDiskVolumeSource, c *conversion.Cloner) error {
	out.PDName = in.PDName
	out.FSType = in.FSType
	out.Partition = in.Partition
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_api_GitRepoVolumeSource(in GitRepoVolumeSource, out *GitRepoVolumeSource, c *conversion.Cloner) error {
	out.Repository = in.Repository
	out.Revision = in.Revision
	return nil
}

func deepCopy_api_GlusterfsVolumeSource(in GlusterfsVolumeSource, out *GlusterfsVolumeSource, c *conversion.Cloner) error {
	out.EndpointsName = in.EndpointsName
	out.Path = in.Path
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_api_HTTPGetAction(in HTTPGetAction, out *HTTPGetAction, c *conversion.Cloner) error {
	out.Path = in.Path
	if err := deepCopy_util_IntOrString(in.Port, &out.Port, c); err != nil {
		return err
	}
	out.Host = in.Host
	out.Scheme = in.Scheme
	return nil
}

func deepCopy_api_Handler(in Handler, out *Handler, c *conversion.Cloner) error {
	if in.Exec != nil {
		out.Exec = new(ExecAction)
		if err := deepCopy_api_ExecAction(*in.Exec, out.Exec, c); err != nil {
			return err
		}
	} else {
		out.Exec = nil
	}
	if in.HTTPGet != nil {
		out.HTTPGet = new(HTTPGetAction)
		if err := deepCopy_api_HTTPGetAction(*in.HTTPGet, out.HTTPGet, c); err != nil {
			return err
		}
	} else {
		out.HTTPGet = nil
	}
	if in.TCPSocket != nil {
		out.TCPSocket = new(TCPSocketAction)
		if err := deepCopy_api_TCPSocketAction(*in.TCPSocket, out.TCPSocket, c); err != nil {
			return err
		}
	} else {
		out.TCPSocket = nil
	}
	return nil
}

func deepCopy_api_HostPathVolumeSource(in HostPathVolumeSource, out *HostPathVolumeSource, c *conversion.Cloner) error {
	out.Path = in.Path
	return nil
}

func deepCopy_api_ISCSIVolumeSource(in ISCSIVolumeSource, out *ISCSIVolumeSource, c *conversion.Cloner) error {
	out.TargetPortal = in.TargetPortal
	out.IQN = in.IQN
	out.Lun = in.Lun
	out.FSType = in.FSType
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_api_Lifecycle(in Lifecycle, out *Lifecycle, c *conversion.Cloner) error {
	if in.PostStart != nil {
		out.PostStart = new(Handler)
		if err := deepCopy_api_Handler(*in.PostStart, out.PostStart, c); err != nil {
			return err
		}
	} else {
		out.PostStart = nil
	}
	if in.PreStop != nil {
		out.PreStop = new(Handler)
		if err := deepCopy_api_Handler(*in.PreStop, out.PreStop, c); err != nil {
			return err
		}
	} else {
		out.PreStop = nil
	}
	return nil
}

func deepCopy_api_LimitRange(in LimitRange, out *LimitRange, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_LimitRangeSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_LimitRangeItem(in LimitRangeItem, out *LimitRangeItem, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.Max != nil {
		out.Max = make(ResourceList)
		for key, val := range in.Max {
			newVal := new(resource.Quantity)
			if err := deepCopy_resource_Quantity(val, newVal, c); err != nil {
				return err
			}
			out.Max[key] = *newVal
		}
	} else {
		out.Max = nil
	}
	if in.Min != nil {
		out.Min = make(ResourceList)
		for key, val := range in.Min {
			newVal := new(resource.Quantity)
			if err := deepCopy_resource_Quantity(val, newVal, c); err != nil {
				return err
			}
			out.Min[key] = *newVal
		}
	} else {
		out.Min = nil
	}
	if in.Default != nil {
		out.Default = make(ResourceList)
		for key, val := range in.Default {
			newVal := new(resource.Quantity)
			if err := deepCopy_resource_Quantity(val, newVal, c); err != nil {
				return err
			}
			out.Default[key] = *newVal
		}
	} else {
		out.Default = nil
	}
	return nil
}

func deepCopy_api_LimitRangeList(in LimitRangeList, out *LimitRangeList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]LimitRange, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_LimitRange(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_LimitRangeSpec(in LimitRangeSpec, out *LimitRangeSpec, c *conversion.Cloner) error {
	if in.Limits != nil {
		out.Limits = make([]LimitRangeItem, len(in.Limits))
		for i := range in.Limits {
			if err := deepCopy_api_LimitRangeItem(in.Limits[i], &out.Limits[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Limits = nil
	}
	return nil
}

func deepCopy_api_List(in List, out *List, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]runtime.Object, len(in.Items))
		for i := range in.Items {
			if newVal, err := c.DeepCopy(in.Items[i]); err != nil {
				return err
			} else {
				out.Items[i] = newVal.(runtime.Object)
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_ListMeta(in ListMeta, out *ListMeta, c *conversion.Cloner) error {
	out.SelfLink = in.SelfLink
	out.ResourceVersion = in.ResourceVersion
	return nil
}

func deepCopy_api_ListOptions(in ListOptions, out *ListOptions, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if newVal, err := c.DeepCopy(in.LabelSelector); err != nil {
		return err
	} else {
		out.LabelSelector = newVal.(labels.Selector)
	}
	if newVal, err := c.DeepCopy(in.FieldSelector); err != nil {
		return err
	} else {
		out.FieldSelector = newVal.(fields.Selector)
	}
	out.Watch = in.Watch
	out.ResourceVersion = in.ResourceVersion
	return nil
}

func deepCopy_api_LoadBalancerIngress(in LoadBalancerIngress, out *LoadBalancerIngress, c *conversion.Cloner) error {
	out.IP = in.IP
	out.Hostname = in.Hostname
	return nil
}

func deepCopy_api_LoadBalancerStatus(in LoadBalancerStatus, out *LoadBalancerStatus, c *conversion.Cloner) error {
	if in.Ingress != nil {
		out.Ingress = make([]LoadBalancerIngress, len(in.Ingress))
		for i := range in.Ingress {
			if err := deepCopy_api_LoadBalancerIngress(in.Ingress[i], &out.Ingress[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Ingress = nil
	}
	return nil
}

func deepCopy_api_LocalObjectReference(in LocalObjectReference, out *LocalObjectReference, c *conversion.Cloner) error {
	out.Name = in.Name
	return nil
}

func deepCopy_api_NFSVolumeSource(in NFSVolumeSource, out *NFSVolumeSource, c *conversion.Cloner) error {
	out.Server = in.Server
	out.Path = in.Path
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_api_Namespace(in Namespace, out *Namespace, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_NamespaceSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_api_NamespaceStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_NamespaceList(in NamespaceList, out *NamespaceList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]Namespace, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_Namespace(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_NamespaceSpec(in NamespaceSpec, out *NamespaceSpec, c *conversion.Cloner) error {
	if in.Finalizers != nil {
		out.Finalizers = make([]FinalizerName, len(in.Finalizers))
		for i := range in.Finalizers {
			out.Finalizers[i] = in.Finalizers[i]
		}
	} else {
		out.Finalizers = nil
	}
	return nil
}

func deepCopy_api_NamespaceStatus(in NamespaceStatus, out *NamespaceStatus, c *conversion.Cloner) error {
	out.Phase = in.Phase
	return nil
}

func deepCopy_api_Node(in Node, out *Node, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_NodeSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_api_NodeStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_NodeAddress(in NodeAddress, out *NodeAddress, c *conversion.Cloner) error {
	out.Type = in.Type
	out.Address = in.Address
	return nil
}

func deepCopy_api_NodeCondition(in NodeCondition, out *NodeCondition, c *conversion.Cloner) error {
	out.Type = in.Type
	out.Status = in.Status
	if err := deepCopy_util_Time(in.LastHeartbeatTime, &out.LastHeartbeatTime, c); err != nil {
		return err
	}
	if err := deepCopy_util_Time(in.LastTransitionTime, &out.LastTransitionTime, c); err != nil {
		return err
	}
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func deepCopy_api_NodeList(in NodeList, out *NodeList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]Node, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_Node(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_NodeSpec(in NodeSpec, out *NodeSpec, c *conversion.Cloner) error {
	out.PodCIDR = in.PodCIDR
	out.ExternalID = in.ExternalID
	out.ProviderID = in.ProviderID
	out.Unschedulable = in.Unschedulable
	return nil
}

func deepCopy_api_NodeStatus(in NodeStatus, out *NodeStatus, c *conversion.Cloner) error {
	if in.Capacity != nil {
		out.Capacity = make(ResourceList)
		for key, val := range in.Capacity {
			newVal := new(resource.Quantity)
			if err := deepCopy_resource_Quantity(val, newVal, c); err != nil {
				return err
			}
			out.Capacity[key] = *newVal
		}
	} else {
		out.Capacity = nil
	}
	out.Phase = in.Phase
	if in.Conditions != nil {
		out.Conditions = make([]NodeCondition, len(in.Conditions))
		for i := range in.Conditions {
			if err := deepCopy_api_NodeCondition(in.Conditions[i], &out.Conditions[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	if in.Addresses != nil {
		out.Addresses = make([]NodeAddress, len(in.Addresses))
		for i := range in.Addresses {
			if err := deepCopy_api_NodeAddress(in.Addresses[i], &out.Addresses[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Addresses = nil
	}
	if err := deepCopy_api_NodeSystemInfo(in.NodeInfo, &out.NodeInfo, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_NodeSystemInfo(in NodeSystemInfo, out *NodeSystemInfo, c *conversion.Cloner) error {
	out.MachineID = in.MachineID
	out.SystemUUID = in.SystemUUID
	out.BootID = in.BootID
	out.KernelVersion = in.KernelVersion
	out.OsImage = in.OsImage
	out.ContainerRuntimeVersion = in.ContainerRuntimeVersion
	out.KubeletVersion = in.KubeletVersion
	out.KubeProxyVersion = in.KubeProxyVersion
	return nil
}

func deepCopy_api_ObjectFieldSelector(in ObjectFieldSelector, out *ObjectFieldSelector, c *conversion.Cloner) error {
	out.APIVersion = in.APIVersion
	out.FieldPath = in.FieldPath
	return nil
}

func deepCopy_api_ObjectMeta(in ObjectMeta, out *ObjectMeta, c *conversion.Cloner) error {
	out.Name = in.Name
	out.GenerateName = in.GenerateName
	out.Namespace = in.Namespace
	out.SelfLink = in.SelfLink
	out.UID = in.UID
	out.ResourceVersion = in.ResourceVersion
	out.Generation = in.Generation
	if err := deepCopy_util_Time(in.CreationTimestamp, &out.CreationTimestamp, c); err != nil {
		return err
	}
	if in.DeletionTimestamp != nil {
		out.DeletionTimestamp = new(util.Time)
		if err := deepCopy_util_Time(*in.DeletionTimestamp, out.DeletionTimestamp, c); err != nil {
			return err
		}
	} else {
		out.DeletionTimestamp = nil
	}
	if in.Labels != nil {
		out.Labels = make(map[string]string)
		for key, val := range in.Labels {
			out.Labels[key] = val
		}
	} else {
		out.Labels = nil
	}
	if in.Annotations != nil {
		out.Annotations = make(map[string]string)
		for key, val := range in.Annotations {
			out.Annotations[key] = val
		}
	} else {
		out.Annotations = nil
	}
	return nil
}

func deepCopy_api_ObjectReference(in ObjectReference, out *ObjectReference, c *conversion.Cloner) error {
	out.Kind = in.Kind
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.UID = in.UID
	out.APIVersion = in.APIVersion
	out.ResourceVersion = in.ResourceVersion
	out.FieldPath = in.FieldPath
	return nil
}

func deepCopy_api_PersistentVolume(in PersistentVolume, out *PersistentVolume, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_PersistentVolumeSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_api_PersistentVolumeStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_PersistentVolumeClaim(in PersistentVolumeClaim, out *PersistentVolumeClaim, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_PersistentVolumeClaimSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_api_PersistentVolumeClaimStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_PersistentVolumeClaimList(in PersistentVolumeClaimList, out *PersistentVolumeClaimList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]PersistentVolumeClaim, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_PersistentVolumeClaim(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_PersistentVolumeClaimSpec(in PersistentVolumeClaimSpec, out *PersistentVolumeClaimSpec, c *conversion.Cloner) error {
	if in.AccessModes != nil {
		out.AccessModes = make([]PersistentVolumeAccessMode, len(in.AccessModes))
		for i := range in.AccessModes {
			out.AccessModes[i] = in.AccessModes[i]
		}
	} else {
		out.AccessModes = nil
	}
	if err := deepCopy_api_ResourceRequirements(in.Resources, &out.Resources, c); err != nil {
		return err
	}
	out.VolumeName = in.VolumeName
	return nil
}

func deepCopy_api_PersistentVolumeClaimStatus(in PersistentVolumeClaimStatus, out *PersistentVolumeClaimStatus, c *conversion.Cloner) error {
	out.Phase = in.Phase
	if in.AccessModes != nil {
		out.AccessModes = make([]PersistentVolumeAccessMode, len(in.AccessModes))
		for i := range in.AccessModes {
			out.AccessModes[i] = in.AccessModes[i]
		}
	} else {
		out.AccessModes = nil
	}
	if in.Capacity != nil {
		out.Capacity = make(ResourceList)
		for key, val := range in.Capacity {
			newVal := new(resource.Quantity)
			if err := deepCopy_resource_Quantity(val, newVal, c); err != nil {
				return err
			}
			out.Capacity[key] = *newVal
		}
	} else {
		out.Capacity = nil
	}
	return nil
}

func deepCopy_api_PersistentVolumeClaimVolumeSource(in PersistentVolumeClaimVolumeSource, out *PersistentVolumeClaimVolumeSource, c *conversion.Cloner) error {
	out.ClaimName = in.ClaimName
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_api_PersistentVolumeList(in PersistentVolumeList, out *PersistentVolumeList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]PersistentVolume, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_PersistentVolume(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_PersistentVolumeSource(in PersistentVolumeSource, out *PersistentVolumeSource, c *conversion.Cloner) error {
	if in.GCEPersistentDisk != nil {
		out.GCEPersistentDisk = new(GCEPersistentDiskVolumeSource)
		if err := deepCopy_api_GCEPersistentDiskVolumeSource(*in.GCEPersistentDisk, out.GCEPersistentDisk, c); err != nil {
			return err
		}
	} else {
		out.GCEPersistentDisk = nil
	}
	if in.AWSElasticBlockStore != nil {
		out.AWSElasticBlockStore = new(AWSElasticBlockStoreVolumeSource)
		if err := deepCopy_api_AWSElasticBlockStoreVolumeSource(*in.AWSElasticBlockStore, out.AWSElasticBlockStore, c); err != nil {
			return err
		}
	} else {
		out.AWSElasticBlockStore = nil
	}
	if in.HostPath != nil {
		out.HostPath = new(HostPathVolumeSource)
		if err := deepCopy_api_HostPathVolumeSource(*in.HostPath, out.HostPath, c); err != nil {
			return err
		}
	} else {
		out.HostPath = nil
	}
	if in.Glusterfs != nil {
		out.Glusterfs = new(GlusterfsVolumeSource)
		if err := deepCopy_api_GlusterfsVolumeSource(*in.Glusterfs, out.Glusterfs, c); err != nil {
			return err
		}
	} else {
		out.Glusterfs = nil
	}
	if in.NFS != nil {
		out.NFS = new(NFSVolumeSource)
		if err := deepCopy_api_NFSVolumeSource(*in.NFS, out.NFS, c); err != nil {
			return err
		}
	} else {
		out.NFS = nil
	}
	if in.RBD != nil {
		out.RBD = new(RBDVolumeSource)
		if err := deepCopy_api_RBDVolumeSource(*in.RBD, out.RBD, c); err != nil {
			return err
		}
	} else {
		out.RBD = nil
	}
	if in.ISCSI != nil {
		out.ISCSI = new(ISCSIVolumeSource)
		if err := deepCopy_api_ISCSIVolumeSource(*in.ISCSI, out.ISCSI, c); err != nil {
			return err
		}
	} else {
		out.ISCSI = nil
	}
	return nil
}

func deepCopy_api_PersistentVolumeSpec(in PersistentVolumeSpec, out *PersistentVolumeSpec, c *conversion.Cloner) error {
	if in.Capacity != nil {
		out.Capacity = make(ResourceList)
		for key, val := range in.Capacity {
			newVal := new(resource.Quantity)
			if err := deepCopy_resource_Quantity(val, newVal, c); err != nil {
				return err
			}
			out.Capacity[key] = *newVal
		}
	} else {
		out.Capacity = nil
	}
	if err := deepCopy_api_PersistentVolumeSource(in.PersistentVolumeSource, &out.PersistentVolumeSource, c); err != nil {
		return err
	}
	if in.AccessModes != nil {
		out.AccessModes = make([]PersistentVolumeAccessMode, len(in.AccessModes))
		for i := range in.AccessModes {
			out.AccessModes[i] = in.AccessModes[i]
		}
	} else {
		out.AccessModes = nil
	}
	if in.ClaimRef != nil {
		out.ClaimRef = new(ObjectReference)
		if err := deepCopy_api_ObjectReference(*in.ClaimRef, out.ClaimRef, c); err != nil {
			return err
		}
	} else {
		out.ClaimRef = nil
	}
	out.PersistentVolumeReclaimPolicy = in.PersistentVolumeReclaimPolicy
	return nil
}

func deepCopy_api_PersistentVolumeStatus(in PersistentVolumeStatus, out *PersistentVolumeStatus, c *conversion.Cloner) error {
	out.Phase = in.Phase
	out.Message = in.Message
	out.Reason = in.Reason
	return nil
}

func deepCopy_api_Pod(in Pod, out *Pod, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_PodSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_api_PodStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_PodAttachOptions(in PodAttachOptions, out *PodAttachOptions, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.Stdin = in.Stdin
	out.Stdout = in.Stdout
	out.Stderr = in.Stderr
	out.TTY = in.TTY
	out.Container = in.Container
	return nil
}

func deepCopy_api_PodCondition(in PodCondition, out *PodCondition, c *conversion.Cloner) error {
	out.Type = in.Type
	out.Status = in.Status
	return nil
}

func deepCopy_api_PodExecOptions(in PodExecOptions, out *PodExecOptions, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.Stdin = in.Stdin
	out.Stdout = in.Stdout
	out.Stderr = in.Stderr
	out.TTY = in.TTY
	out.Container = in.Container
	if in.Command != nil {
		out.Command = make([]string, len(in.Command))
		for i := range in.Command {
			out.Command[i] = in.Command[i]
		}
	} else {
		out.Command = nil
	}
	return nil
}

func deepCopy_api_PodList(in PodList, out *PodList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]Pod, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_Pod(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_PodLogOptions(in PodLogOptions, out *PodLogOptions, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.Container = in.Container
	out.Follow = in.Follow
	out.Previous = in.Previous
	return nil
}

func deepCopy_api_PodProxyOptions(in PodProxyOptions, out *PodProxyOptions, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.Path = in.Path
	return nil
}

func deepCopy_api_PodSpec(in PodSpec, out *PodSpec, c *conversion.Cloner) error {
	if in.Volumes != nil {
		out.Volumes = make([]Volume, len(in.Volumes))
		for i := range in.Volumes {
			if err := deepCopy_api_Volume(in.Volumes[i], &out.Volumes[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Volumes = nil
	}
	if in.Containers != nil {
		out.Containers = make([]Container, len(in.Containers))
		for i := range in.Containers {
			if err := deepCopy_api_Container(in.Containers[i], &out.Containers[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Containers = nil
	}
	out.RestartPolicy = in.RestartPolicy
	if in.TerminationGracePeriodSeconds != nil {
		out.TerminationGracePeriodSeconds = new(int64)
		*out.TerminationGracePeriodSeconds = *in.TerminationGracePeriodSeconds
	} else {
		out.TerminationGracePeriodSeconds = nil
	}
	if in.ActiveDeadlineSeconds != nil {
		out.ActiveDeadlineSeconds = new(int64)
		*out.ActiveDeadlineSeconds = *in.ActiveDeadlineSeconds
	} else {
		out.ActiveDeadlineSeconds = nil
	}
	out.DNSPolicy = in.DNSPolicy
	if in.NodeSelector != nil {
		out.NodeSelector = make(map[string]string)
		for key, val := range in.NodeSelector {
			out.NodeSelector[key] = val
		}
	} else {
		out.NodeSelector = nil
	}
	out.ServiceAccountName = in.ServiceAccountName
	out.NodeName = in.NodeName
	out.HostNetwork = in.HostNetwork
	if in.ImagePullSecrets != nil {
		out.ImagePullSecrets = make([]LocalObjectReference, len(in.ImagePullSecrets))
		for i := range in.ImagePullSecrets {
			if err := deepCopy_api_LocalObjectReference(in.ImagePullSecrets[i], &out.ImagePullSecrets[i], c); err != nil {
				return err
			}
		}
	} else {
		out.ImagePullSecrets = nil
	}
	return nil
}

func deepCopy_api_PodStatus(in PodStatus, out *PodStatus, c *conversion.Cloner) error {
	out.Phase = in.Phase
	if in.Conditions != nil {
		out.Conditions = make([]PodCondition, len(in.Conditions))
		for i := range in.Conditions {
			if err := deepCopy_api_PodCondition(in.Conditions[i], &out.Conditions[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	out.Message = in.Message
	out.Reason = in.Reason
	out.HostIP = in.HostIP
	out.PodIP = in.PodIP
	if in.StartTime != nil {
		out.StartTime = new(util.Time)
		if err := deepCopy_util_Time(*in.StartTime, out.StartTime, c); err != nil {
			return err
		}
	} else {
		out.StartTime = nil
	}
	if in.ContainerStatuses != nil {
		out.ContainerStatuses = make([]ContainerStatus, len(in.ContainerStatuses))
		for i := range in.ContainerStatuses {
			if err := deepCopy_api_ContainerStatus(in.ContainerStatuses[i], &out.ContainerStatuses[i], c); err != nil {
				return err
			}
		}
	} else {
		out.ContainerStatuses = nil
	}
	return nil
}

func deepCopy_api_PodStatusResult(in PodStatusResult, out *PodStatusResult, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_PodStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_PodTemplate(in PodTemplate, out *PodTemplate, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_PodTemplateSpec(in.Template, &out.Template, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_PodTemplateList(in PodTemplateList, out *PodTemplateList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]PodTemplate, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_PodTemplate(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_PodTemplateSpec(in PodTemplateSpec, out *PodTemplateSpec, c *conversion.Cloner) error {
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_PodSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_Probe(in Probe, out *Probe, c *conversion.Cloner) error {
	if err := deepCopy_api_Handler(in.Handler, &out.Handler, c); err != nil {
		return err
	}
	out.InitialDelaySeconds = in.InitialDelaySeconds
	out.TimeoutSeconds = in.TimeoutSeconds
	return nil
}

func deepCopy_api_RBDVolumeSource(in RBDVolumeSource, out *RBDVolumeSource, c *conversion.Cloner) error {
	if in.CephMonitors != nil {
		out.CephMonitors = make([]string, len(in.CephMonitors))
		for i := range in.CephMonitors {
			out.CephMonitors[i] = in.CephMonitors[i]
		}
	} else {
		out.CephMonitors = nil
	}
	out.RBDImage = in.RBDImage
	out.FSType = in.FSType
	out.RBDPool = in.RBDPool
	out.RadosUser = in.RadosUser
	out.Keyring = in.Keyring
	if in.SecretRef != nil {
		out.SecretRef = new(LocalObjectReference)
		if err := deepCopy_api_LocalObjectReference(*in.SecretRef, out.SecretRef, c); err != nil {
			return err
		}
	} else {
		out.SecretRef = nil
	}
	out.ReadOnly = in.ReadOnly
	return nil
}

func deepCopy_api_RangeAllocation(in RangeAllocation, out *RangeAllocation, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	out.Range = in.Range
	if in.Data != nil {
		out.Data = make([]uint8, len(in.Data))
		for i := range in.Data {
			out.Data[i] = in.Data[i]
		}
	} else {
		out.Data = nil
	}
	return nil
}

func deepCopy_api_ReplicationController(in ReplicationController, out *ReplicationController, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ReplicationControllerSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_api_ReplicationControllerStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_ReplicationControllerList(in ReplicationControllerList, out *ReplicationControllerList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]ReplicationController, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_ReplicationController(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_ReplicationControllerSpec(in ReplicationControllerSpec, out *ReplicationControllerSpec, c *conversion.Cloner) error {
	out.Replicas = in.Replicas
	if in.Selector != nil {
		out.Selector = make(map[string]string)
		for key, val := range in.Selector {
			out.Selector[key] = val
		}
	} else {
		out.Selector = nil
	}
	if in.Template != nil {
		out.Template = new(PodTemplateSpec)
		if err := deepCopy_api_PodTemplateSpec(*in.Template, out.Template, c); err != nil {
			return err
		}
	} else {
		out.Template = nil
	}
	return nil
}

func deepCopy_api_ReplicationControllerStatus(in ReplicationControllerStatus, out *ReplicationControllerStatus, c *conversion.Cloner) error {
	out.Replicas = in.Replicas
	out.ObservedGeneration = in.ObservedGeneration
	return nil
}

func deepCopy_api_ResourceQuota(in ResourceQuota, out *ResourceQuota, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ResourceQuotaSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_api_ResourceQuotaStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_ResourceQuotaList(in ResourceQuotaList, out *ResourceQuotaList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]ResourceQuota, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_ResourceQuota(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_ResourceQuotaSpec(in ResourceQuotaSpec, out *ResourceQuotaSpec, c *conversion.Cloner) error {
	if in.Hard != nil {
		out.Hard = make(ResourceList)
		for key, val := range in.Hard {
			newVal := new(resource.Quantity)
			if err := deepCopy_resource_Quantity(val, newVal, c); err != nil {
				return err
			}
			out.Hard[key] = *newVal
		}
	} else {
		out.Hard = nil
	}
	return nil
}

func deepCopy_api_ResourceQuotaStatus(in ResourceQuotaStatus, out *ResourceQuotaStatus, c *conversion.Cloner) error {
	if in.Hard != nil {
		out.Hard = make(ResourceList)
		for key, val := range in.Hard {
			newVal := new(resource.Quantity)
			if err := deepCopy_resource_Quantity(val, newVal, c); err != nil {
				return err
			}
			out.Hard[key] = *newVal
		}
	} else {
		out.Hard = nil
	}
	if in.Used != nil {
		out.Used = make(ResourceList)
		for key, val := range in.Used {
			newVal := new(resource.Quantity)
			if err := deepCopy_resource_Quantity(val, newVal, c); err != nil {
				return err
			}
			out.Used[key] = *newVal
		}
	} else {
		out.Used = nil
	}
	return nil
}

func deepCopy_api_ResourceRequirements(in ResourceRequirements, out *ResourceRequirements, c *conversion.Cloner) error {
	if in.Limits != nil {
		out.Limits = make(ResourceList)
		for key, val := range in.Limits {
			newVal := new(resource.Quantity)
			if err := deepCopy_resource_Quantity(val, newVal, c); err != nil {
				return err
			}
			out.Limits[key] = *newVal
		}
	} else {
		out.Limits = nil
	}
	if in.Requests != nil {
		out.Requests = make(ResourceList)
		for key, val := range in.Requests {
			newVal := new(resource.Quantity)
			if err := deepCopy_resource_Quantity(val, newVal, c); err != nil {
				return err
			}
			out.Requests[key] = *newVal
		}
	} else {
		out.Requests = nil
	}
	return nil
}

func deepCopy_api_SELinuxOptions(in SELinuxOptions, out *SELinuxOptions, c *conversion.Cloner) error {
	out.User = in.User
	out.Role = in.Role
	out.Type = in.Type
	out.Level = in.Level
	return nil
}

func deepCopy_api_Secret(in Secret, out *Secret, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if in.Data != nil {
		out.Data = make(map[string][]uint8)
		for key, val := range in.Data {
			if newVal, err := c.DeepCopy(val); err != nil {
				return err
			} else {
				out.Data[key] = newVal.([]uint8)
			}
		}
	} else {
		out.Data = nil
	}
	out.Type = in.Type
	return nil
}

func deepCopy_api_SecretList(in SecretList, out *SecretList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]Secret, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_Secret(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_SecretVolumeSource(in SecretVolumeSource, out *SecretVolumeSource, c *conversion.Cloner) error {
	out.SecretName = in.SecretName
	return nil
}

func deepCopy_api_SecurityContext(in SecurityContext, out *SecurityContext, c *conversion.Cloner) error {
	if in.Capabilities != nil {
		out.Capabilities = new(Capabilities)
		if err := deepCopy_api_Capabilities(*in.Capabilities, out.Capabilities, c); err != nil {
			return err
		}
	} else {
		out.Capabilities = nil
	}
	if in.Privileged != nil {
		out.Privileged = new(bool)
		*out.Privileged = *in.Privileged
	} else {
		out.Privileged = nil
	}
	if in.SELinuxOptions != nil {
		out.SELinuxOptions = new(SELinuxOptions)
		if err := deepCopy_api_SELinuxOptions(*in.SELinuxOptions, out.SELinuxOptions, c); err != nil {
			return err
		}
	} else {
		out.SELinuxOptions = nil
	}
	if in.RunAsUser != nil {
		out.RunAsUser = new(int64)
		*out.RunAsUser = *in.RunAsUser
	} else {
		out.RunAsUser = nil
	}
	return nil
}

func deepCopy_api_SerializedReference(in SerializedReference, out *SerializedReference, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectReference(in.Reference, &out.Reference, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_Service(in Service, out *Service, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ServiceSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_api_ServiceStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_ServiceAccount(in ServiceAccount, out *ServiceAccount, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if in.Secrets != nil {
		out.Secrets = make([]ObjectReference, len(in.Secrets))
		for i := range in.Secrets {
			if err := deepCopy_api_ObjectReference(in.Secrets[i], &out.Secrets[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Secrets = nil
	}
	if in.ImagePullSecrets != nil {
		out.ImagePullSecrets = make([]LocalObjectReference, len(in.ImagePullSecrets))
		for i := range in.ImagePullSecrets {
			if err := deepCopy_api_LocalObjectReference(in.ImagePullSecrets[i], &out.ImagePullSecrets[i], c); err != nil {
				return err
			}
		}
	} else {
		out.ImagePullSecrets = nil
	}
	return nil
}

func deepCopy_api_ServiceAccountList(in ServiceAccountList, out *ServiceAccountList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]ServiceAccount, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_ServiceAccount(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_ServiceList(in ServiceList, out *ServiceList, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]Service, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_api_Service(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_api_ServicePort(in ServicePort, out *ServicePort, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Protocol = in.Protocol
	out.Port = in.Port
	if err := deepCopy_util_IntOrString(in.TargetPort, &out.TargetPort, c); err != nil {
		return err
	}
	out.NodePort = in.NodePort
	return nil
}

func deepCopy_api_ServiceSpec(in ServiceSpec, out *ServiceSpec, c *conversion.Cloner) error {
	if in.Ports != nil {
		out.Ports = make([]ServicePort, len(in.Ports))
		for i := range in.Ports {
			if err := deepCopy_api_ServicePort(in.Ports[i], &out.Ports[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Ports = nil
	}
	if in.Selector != nil {
		out.Selector = make(map[string]string)
		for key, val := range in.Selector {
			out.Selector[key] = val
		}
	} else {
		out.Selector = nil
	}
	out.ClusterIP = in.ClusterIP
	out.Type = in.Type
	if in.DeprecatedPublicIPs != nil {
		out.DeprecatedPublicIPs = make([]string, len(in.DeprecatedPublicIPs))
		for i := range in.DeprecatedPublicIPs {
			out.DeprecatedPublicIPs[i] = in.DeprecatedPublicIPs[i]
		}
	} else {
		out.DeprecatedPublicIPs = nil
	}
	out.SessionAffinity = in.SessionAffinity
	return nil
}

func deepCopy_api_ServiceStatus(in ServiceStatus, out *ServiceStatus, c *conversion.Cloner) error {
	if err := deepCopy_api_LoadBalancerStatus(in.LoadBalancer, &out.LoadBalancer, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_Status(in Status, out *Status, c *conversion.Cloner) error {
	if err := deepCopy_api_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_api_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	out.Status = in.Status
	out.Message = in.Message
	out.Reason = in.Reason
	if in.Details != nil {
		out.Details = new(StatusDetails)
		if err := deepCopy_api_StatusDetails(*in.Details, out.Details, c); err != nil {
			return err
		}
	} else {
		out.Details = nil
	}
	out.Code = in.Code
	return nil
}

func deepCopy_api_StatusCause(in StatusCause, out *StatusCause, c *conversion.Cloner) error {
	out.Type = in.Type
	out.Message = in.Message
	out.Field = in.Field
	return nil
}

func deepCopy_api_StatusDetails(in StatusDetails, out *StatusDetails, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Kind = in.Kind
	if in.Causes != nil {
		out.Causes = make([]StatusCause, len(in.Causes))
		for i := range in.Causes {
			if err := deepCopy_api_StatusCause(in.Causes[i], &out.Causes[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Causes = nil
	}
	out.RetryAfterSeconds = in.RetryAfterSeconds
	return nil
}

func deepCopy_api_TCPSocketAction(in TCPSocketAction, out *TCPSocketAction, c *conversion.Cloner) error {
	if err := deepCopy_util_IntOrString(in.Port, &out.Port, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_TypeMeta(in TypeMeta, out *TypeMeta, c *conversion.Cloner) error {
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	return nil
}

func deepCopy_api_Volume(in Volume, out *Volume, c *conversion.Cloner) error {
	out.Name = in.Name
	if err := deepCopy_api_VolumeSource(in.VolumeSource, &out.VolumeSource, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_api_VolumeMount(in VolumeMount, out *VolumeMount, c *conversion.Cloner) error {
	out.Name = in.Name
	out.ReadOnly = in.ReadOnly
	out.MountPath = in.MountPath
	return nil
}

func deepCopy_api_VolumeSource(in VolumeSource, out *VolumeSource, c *conversion.Cloner) error {
	if in.HostPath != nil {
		out.HostPath = new(HostPathVolumeSource)
		if err := deepCopy_api_HostPathVolumeSource(*in.HostPath, out.HostPath, c); err != nil {
			return err
		}
	} else {
		out.HostPath = nil
	}
	if in.EmptyDir != nil {
		out.EmptyDir = new(EmptyDirVolumeSource)
		if err := deepCopy_api_EmptyDirVolumeSource(*in.EmptyDir, out.EmptyDir, c); err != nil {
			return err
		}
	} else {
		out.EmptyDir = nil
	}
	if in.GCEPersistentDisk != nil {
		out.GCEPersistentDisk = new(GCEPersistentDiskVolumeSource)
		if err := deepCopy_api_GCEPersistentDiskVolumeSource(*in.GCEPersistentDisk, out.GCEPersistentDisk, c); err != nil {
			return err
		}
	} else {
		out.GCEPersistentDisk = nil
	}
	if in.AWSElasticBlockStore != nil {
		out.AWSElasticBlockStore = new(AWSElasticBlockStoreVolumeSource)
		if err := deepCopy_api_AWSElasticBlockStoreVolumeSource(*in.AWSElasticBlockStore, out.AWSElasticBlockStore, c); err != nil {
			return err
		}
	} else {
		out.AWSElasticBlockStore = nil
	}
	if in.GitRepo != nil {
		out.GitRepo = new(GitRepoVolumeSource)
		if err := deepCopy_api_GitRepoVolumeSource(*in.GitRepo, out.GitRepo, c); err != nil {
			return err
		}
	} else {
		out.GitRepo = nil
	}
	if in.Secret != nil {
		out.Secret = new(SecretVolumeSource)
		if err := deepCopy_api_SecretVolumeSource(*in.Secret, out.Secret, c); err != nil {
			return err
		}
	} else {
		out.Secret = nil
	}
	if in.NFS != nil {
		out.NFS = new(NFSVolumeSource)
		if err := deepCopy_api_NFSVolumeSource(*in.NFS, out.NFS, c); err != nil {
			return err
		}
	} else {
		out.NFS = nil
	}
	if in.ISCSI != nil {
		out.ISCSI = new(ISCSIVolumeSource)
		if err := deepCopy_api_ISCSIVolumeSource(*in.ISCSI, out.ISCSI, c); err != nil {
			return err
		}
	} else {
		out.ISCSI = nil
	}
	if in.Glusterfs != nil {
		out.Glusterfs = new(GlusterfsVolumeSource)
		if err := deepCopy_api_GlusterfsVolumeSource(*in.Glusterfs, out.Glusterfs, c); err != nil {
			return err
		}
	} else {
		out.Glusterfs = nil
	}
	if in.PersistentVolumeClaim != nil {
		out.PersistentVolumeClaim = new(PersistentVolumeClaimVolumeSource)
		if err := deepCopy_api_PersistentVolumeClaimVolumeSource(*in.PersistentVolumeClaim, out.PersistentVolumeClaim, c); err != nil {
			return err
		}
	} else {
		out.PersistentVolumeClaim = nil
	}
	if in.RBD != nil {
		out.RBD = new(RBDVolumeSource)
		if err := deepCopy_api_RBDVolumeSource(*in.RBD, out.RBD, c); err != nil {
			return err
		}
	} else {
		out.RBD = nil
	}
	return nil
}

func deepCopy_resource_Quantity(in resource.Quantity, out *resource.Quantity, c *conversion.Cloner) error {
	if in.Amount != nil {
		if newVal, err := c.DeepCopy(in.Amount); err != nil {
			return err
		} else {
			out.Amount = newVal.(*inf.Dec)
		}
	} else {
		out.Amount = nil
	}
	out.Format = in.Format
	return nil
}

func deepCopy_util_IntOrString(in util.IntOrString, out *util.IntOrString, c *conversion.Cloner) error {
	out.Kind = in.Kind
	out.IntVal = in.IntVal
	out.StrVal = in.StrVal
	return nil
}

func deepCopy_util_Time(in util.Time, out *util.Time, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.Time); err != nil {
		return err
	} else {
		out.Time = newVal.(time.Time)
	}
	return nil
}

func init() {
	err := Scheme.AddGeneratedDeepCopyFuncs(
		deepCopy_api_AWSElasticBlockStoreVolumeSource,
		deepCopy_api_AutoScaleIntentionThresholdConfig,
		deepCopy_api_AutoScaleThreshold,
		deepCopy_api_AutoScaler,
		deepCopy_api_AutoScalerList,
		deepCopy_api_AutoScalerSpec,
		deepCopy_api_AutoScalerStatus,
		deepCopy_api_Binding,
		deepCopy_api_Capabilities,
		deepCopy_api_ComponentCondition,
		deepCopy_api_ComponentStatus,
		deepCopy_api_ComponentStatusList,
		deepCopy_api_Container,
		deepCopy_api_ContainerPort,
		deepCopy_api_ContainerState,
		deepCopy_api_ContainerStateRunning,
		deepCopy_api_ContainerStateTerminated,
		deepCopy_api_ContainerStateWaiting,
		deepCopy_api_ContainerStatus,
		deepCopy_api_DeleteOptions,
		deepCopy_api_EmptyDirVolumeSource,
		deepCopy_api_EndpointAddress,
		deepCopy_api_EndpointPort,
		deepCopy_api_EndpointSubset,
		deepCopy_api_Endpoints,
		deepCopy_api_EndpointsList,
		deepCopy_api_EnvVar,
		deepCopy_api_EnvVarSource,
		deepCopy_api_Event,
		deepCopy_api_EventList,
		deepCopy_api_EventSource,
		deepCopy_api_ExecAction,
		deepCopy_api_GCEPersistentDiskVolumeSource,
		deepCopy_api_GitRepoVolumeSource,
		deepCopy_api_GlusterfsVolumeSource,
		deepCopy_api_HTTPGetAction,
		deepCopy_api_Handler,
		deepCopy_api_HostPathVolumeSource,
		deepCopy_api_ISCSIVolumeSource,
		deepCopy_api_Lifecycle,
		deepCopy_api_LimitRange,
		deepCopy_api_LimitRangeItem,
		deepCopy_api_LimitRangeList,
		deepCopy_api_LimitRangeSpec,
		deepCopy_api_List,
		deepCopy_api_ListMeta,
		deepCopy_api_ListOptions,
		deepCopy_api_LoadBalancerIngress,
		deepCopy_api_LoadBalancerStatus,
		deepCopy_api_LocalObjectReference,
		deepCopy_api_NFSVolumeSource,
		deepCopy_api_Namespace,
		deepCopy_api_NamespaceList,
		deepCopy_api_NamespaceSpec,
		deepCopy_api_NamespaceStatus,
		deepCopy_api_Node,
		deepCopy_api_NodeAddress,
		deepCopy_api_NodeCondition,
		deepCopy_api_NodeList,
		deepCopy_api_NodeSpec,
		deepCopy_api_NodeStatus,
		deepCopy_api_NodeSystemInfo,
		deepCopy_api_ObjectFieldSelector,
		deepCopy_api_ObjectMeta,
		deepCopy_api_ObjectReference,
		deepCopy_api_PersistentVolume,
		deepCopy_api_PersistentVolumeClaim,
		deepCopy_api_PersistentVolumeClaimList,
		deepCopy_api_PersistentVolumeClaimSpec,
		deepCopy_api_PersistentVolumeClaimStatus,
		deepCopy_api_PersistentVolumeClaimVolumeSource,
		deepCopy_api_PersistentVolumeList,
		deepCopy_api_PersistentVolumeSource,
		deepCopy_api_PersistentVolumeSpec,
		deepCopy_api_PersistentVolumeStatus,
		deepCopy_api_Pod,
		deepCopy_api_PodAttachOptions,
		deepCopy_api_PodCondition,
		deepCopy_api_PodExecOptions,
		deepCopy_api_PodList,
		deepCopy_api_PodLogOptions,
		deepCopy_api_PodProxyOptions,
		deepCopy_api_PodSpec,
		deepCopy_api_PodStatus,
		deepCopy_api_PodStatusResult,
		deepCopy_api_PodTemplate,
		deepCopy_api_PodTemplateList,
		deepCopy_api_PodTemplateSpec,
		deepCopy_api_Probe,
		deepCopy_api_RBDVolumeSource,
		deepCopy_api_RangeAllocation,
		deepCopy_api_ReplicationController,
		deepCopy_api_ReplicationControllerList,
		deepCopy_api_ReplicationControllerSpec,
		deepCopy_api_ReplicationControllerStatus,
		deepCopy_api_ResourceQuota,
		deepCopy_api_ResourceQuotaList,
		deepCopy_api_ResourceQuotaSpec,
		deepCopy_api_ResourceQuotaStatus,
		deepCopy_api_ResourceRequirements,
		deepCopy_api_SELinuxOptions,
		deepCopy_api_Secret,
		deepCopy_api_SecretList,
		deepCopy_api_SecretVolumeSource,
		deepCopy_api_SecurityContext,
		deepCopy_api_SerializedReference,
		deepCopy_api_Service,
		deepCopy_api_ServiceAccount,
		deepCopy_api_ServiceAccountList,
		deepCopy_api_ServiceList,
		deepCopy_api_ServicePort,
		deepCopy_api_ServiceSpec,
		deepCopy_api_ServiceStatus,
		deepCopy_api_Status,
		deepCopy_api_StatusCause,
		deepCopy_api_StatusDetails,
		deepCopy_api_TCPSocketAction,
		deepCopy_api_TypeMeta,
		deepCopy_api_Volume,
		deepCopy_api_VolumeMount,
		deepCopy_api_VolumeSource,
		deepCopy_resource_Quantity,
		deepCopy_util_IntOrString,
		deepCopy_util_Time,
	)
	if err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

// AUTO-GENERATED FUNCTIONS END HERE
