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

package v1

// This file contains a collection of methods that can be used from go-resful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_APIVersion = map[string]string{
	"":         "An APIVersion represents a single concrete version of an object model.",
	"name":     "Name of this version (e.g. 'v1').",
	"apiGroup": "The API group to add this object into, default 'experimental'.",
}

func (APIVersion) SwaggerDoc() map[string]string {
	return map_APIVersion
}

var map_DaemonSet = map[string]string{
	"":         "DaemonSet represents the configuration of a daemon set.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec":     "Spec defines the desired behavior of this daemon set. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
	"status":   "Status is the current status of this daemon set. This data may be out of date by some window of time. Populated by the system. Read-only. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
}

func (DaemonSet) SwaggerDoc() map[string]string {
	return map_DaemonSet
}

var map_DaemonSetList = map[string]string{
	"":         "DaemonSetList is a collection of daemon sets.",
	"metadata": "Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is a list of daemon sets.",
}

func (DaemonSetList) SwaggerDoc() map[string]string {
	return map_DaemonSetList
}

var map_DaemonSetSpec = map[string]string{
	"":         "DaemonSetSpec is the specification of a daemon set.",
	"selector": "Selector is a label query over pods that are managed by the daemon set. Must match in order to be controlled. If empty, defaulted to labels on Pod template. More info: http://releases.k8s.io/HEAD/docs/user-guide/labels.md#label-selectors",
	"template": "Template is the object that describes the pod that will be created. The DaemonSet will create exactly one copy of this pod on every node that matches the template's node selector (or on every node if no node selector is specified). More info: http://releases.k8s.io/HEAD/docs/user-guide/replication-controller.md#pod-template",
}

func (DaemonSetSpec) SwaggerDoc() map[string]string {
	return map_DaemonSetSpec
}

var map_DaemonSetStatus = map[string]string{
	"": "DaemonSetStatus represents the current status of a daemon set.",
	"currentNumberScheduled": "CurrentNumberScheduled is the number of nodes that are running exactly 1 daemon pod and are supposed to run the daemon pod.",
	"numberMisscheduled":     "NumberMisscheduled is the number of nodes that are running the daemon pod, but are not supposed to run the daemon pod.",
	"desiredNumberScheduled": "DesiredNumberScheduled is the total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod).",
}

func (DaemonSetStatus) SwaggerDoc() map[string]string {
	return map_DaemonSetStatus
}

var map_Deployment = map[string]string{
	"":         "Deployment enables declarative updates for Pods and ReplicationControllers.",
	"metadata": "Standard object metadata.",
	"spec":     "Specification of the desired behavior of the Deployment.",
	"status":   "Most recently observed status of the Deployment.",
}

func (Deployment) SwaggerDoc() map[string]string {
	return map_Deployment
}

var map_DeploymentList = map[string]string{
	"":         "DeploymentList is a list of Deployments.",
	"metadata": "Standard list metadata.",
	"items":    "Items is the list of Deployments.",
}

func (DeploymentList) SwaggerDoc() map[string]string {
	return map_DeploymentList
}

var map_DeploymentSpec = map[string]string{
	"":               "DeploymentSpec is the specification of the desired behavior of the Deployment.",
	"replicas":       "Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.",
	"selector":       "Label selector for pods. Existing ReplicationControllers whose pods are selected by this will be scaled down.",
	"template":       "Template describes the pods that will be created.",
	"strategy":       "The deployment strategy to use to replace existing pods with new ones.",
	"uniqueLabelKey": "Key of the selector that is added to existing RCs (and label key that is added to its pods) to prevent the existing RCs to select new pods (and old pods being selected by new RC). Users can set this to an empty string to indicate that the system should not add any selector and label. If unspecified, system uses \"deployment.kubernetes.io/podTemplateHash\". Value of this key is hash of DeploymentSpec.PodTemplateSpec. No label is added if this is set to empty string.",
}

func (DeploymentSpec) SwaggerDoc() map[string]string {
	return map_DeploymentSpec
}

var map_DeploymentStatus = map[string]string{
	"":                "DeploymentStatus is the most recently observed status of the Deployment.",
	"replicas":        "Total number of ready pods targeted by this deployment (this includes both the old and new pods).",
	"updatedReplicas": "Total number of new ready pods with the desired template spec.",
}

func (DeploymentStatus) SwaggerDoc() map[string]string {
	return map_DeploymentStatus
}

var map_DeploymentStrategy = map[string]string{
	"":              "DeploymentStrategy describes how to replace existing pods with new ones.",
	"type":          "Type of deployment. Can be \"Recreate\" or \"RollingUpdate\". Default is RollingUpdate.",
	"rollingUpdate": "Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.",
}

func (DeploymentStrategy) SwaggerDoc() map[string]string {
	return map_DeploymentStrategy
}

var map_HorizontalPodAutoscaler = map[string]string{
	"":         "HorizontalPodAutoscaler represents the configuration of a horizontal pod autoscaler.",
	"metadata": "Standard object metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec":     "Spec defines the behaviour of autoscaler. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status.",
	"status":   "Status represents the current information about the autoscaler.",
}

func (HorizontalPodAutoscaler) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscaler
}

var map_HorizontalPodAutoscalerList = map[string]string{
	"":         "HorizontalPodAutoscalerList is a list of HorizontalPodAutoscalers.",
	"metadata": "Standard list metadata.",
	"items":    "Items is the list of HorizontalPodAutoscalers.",
}

func (HorizontalPodAutoscalerList) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerList
}

var map_HorizontalPodAutoscalerSpec = map[string]string{
	"":         "HorizontalPodAutoscalerSpec is the specification of a horizontal pod autoscaler.",
	"scaleRef": "ScaleRef is a reference to Scale subresource. HorizontalPodAutoscaler will learn the current resource consumption from its status, and will set the desired number of pods by modyfying its spec.",
	"minCount": "MinCount is the lower limit for the number of pods that can be set by the autoscaler.",
	"maxCount": "MaxCount is the upper limit for the number of pods that can be set by the autoscaler. It cannot be smaller than MinCount.",
	"target":   "Target is the target average consumption of the given resource that the autoscaler will try to maintain by adjusting the desired number of pods. Currently two types of resources are supported: \"cpu\" and \"memory\".",
}

func (HorizontalPodAutoscalerSpec) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerSpec
}

var map_HorizontalPodAutoscalerStatus = map[string]string{
	"":                   "HorizontalPodAutoscalerStatus contains the current status of a horizontal pod autoscaler",
	"currentReplicas":    "CurrentReplicas is the number of replicas of pods managed by this autoscaler.",
	"desiredReplicas":    "DesiredReplicas is the desired number of replicas of pods managed by this autoscaler.",
	"currentConsumption": "CurrentConsumption is the current average consumption of the given resource that the autoscaler will try to maintain by adjusting the desired number of pods. Two types of resources are supported: \"cpu\" and \"memory\".",
	"lastScaleTimestamp": "LastScaleTimestamp is the last time the HorizontalPodAutoscaler scaled the number of pods. This is used by the autoscaler to controll how often the number of pods is changed.",
}

func (HorizontalPodAutoscalerStatus) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerStatus
}

var map_Ingress = map[string]string{
	"":         "Ingress encapsulates the inputs needed to config the load-balancer. It represents a rule collection of inbound connections from the external network that would be satisfied by a load-balancer.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec":     "Spec is the desired state of the Ingress. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
	"status":   "Status is the current state of the Ingress. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
}

func (Ingress) SwaggerDoc() map[string]string {
	return map_Ingress
}

var map_IngressBackend = map[string]string{
	"":                "IngressBackend describes an unique service endpoint.",
	"serviceName":     "Specify the name of the referenced service.",
	"serviceProtocol": "Specify the protocol of the referenced service backend.",
	"servicePort":     "Specify the port of the referenced service backend.",
}

func (IngressBackend) SwaggerDoc() map[string]string {
	return map_IngressBackend
}

var map_IngressList = map[string]string{
	"":         "IngressList is a collection of Ingress.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is the list of ingress.",
}

func (IngressList) SwaggerDoc() map[string]string {
	return map_IngressList
}

var map_IngressMap = map[string]string{
	"":      "IngressMap represents the rules mapping the paths under a specified host to the related backend services. It will contains a single host and multiple paths.",
	"host":  "Host that points to the services, can be host or host:port. Examples: \"www.example.com\", \"www.example.com:8080\".",
	"paths": "Paths describe a list of load-balancer rules under the specified host.",
}

func (IngressMap) SwaggerDoc() map[string]string {
	return map_IngressMap
}

var map_IngressPath = map[string]string{
	"":        "IngressPath describes a single rule mapping the url path to specified service backend.",
	"path":    "Path defines URL path. Examples: \"/images\", \"/images/*\".",
	"backend": "Define the referenced service endpoint which the traffic will be forwarded to.",
}

func (IngressPath) SwaggerDoc() map[string]string {
	return map_IngressPath
}

var map_IngressSpec = map[string]string{
	"":        "IngressSpec describes the ingress the user wishes to exist.",
	"pathMap": "PathMap defines a list of rules which mapping multiple hosts and paths to the backend services. The rules will be used to configure the corresponding load-balancer. Example: Host + Path -> Service Backend (Service Name + Service Protocol + Service Port)",
}

func (IngressSpec) SwaggerDoc() map[string]string {
	return map_IngressSpec
}

var map_IngressStatus = map[string]string{
	"":             "IngressStatus describe the current state of the ingress.",
	"loadBalancer": "LoadBalancer contains the current status of the load-balancer.",
}

func (IngressStatus) SwaggerDoc() map[string]string {
	return map_IngressStatus
}

var map_Job = map[string]string{
	"":         "Job represents the configuration of a single job.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec":     "Spec is a structure defining the expected behavior of a job. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
	"status":   "Status is a structure describing current status of a job. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
}

func (Job) SwaggerDoc() map[string]string {
	return map_Job
}

var map_JobCondition = map[string]string{
	"":                   "JobCondition describes current state of a job.",
	"type":               "Type of job condition, currently only Complete.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"lastProbeTime":      "Last time the condition was checked.",
	"lastTransitionTime": "Last time the condition transit from one status to another.",
	"reason":             "(brief) reason for the condition's last transition.",
	"message":            "Human readable message indicating details about last transition.",
}

func (JobCondition) SwaggerDoc() map[string]string {
	return map_JobCondition
}

var map_JobList = map[string]string{
	"":         "JobList is a collection of jobs.",
	"metadata": "Standard list metadata More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is the list of Job.",
}

func (JobList) SwaggerDoc() map[string]string {
	return map_JobList
}

var map_JobSpec = map[string]string{
	"":            "JobSpec describes how the job execution will look like.",
	"parallelism": "Parallelism specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism.",
	"completions": "Completions specifies the desired number of successfully finished pods the job should be run with. Defaults to 1.",
	"selector":    "Selector is a label query over pods that should match the pod count.",
	"template":    "Template is the object that describes the pod that will be created when executing a job.",
}

func (JobSpec) SwaggerDoc() map[string]string {
	return map_JobSpec
}

var map_JobStatus = map[string]string{
	"":               "JobStatus represents the current state of a Job.",
	"conditions":     "Conditions represent the latest available observations of an object's current state.",
	"startTime":      "StartTime represents time when the job was acknowledged by the Job Manager. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
	"completionTime": "CompletionTime represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
	"active":         "Active is the number of actively running pods.",
	"successful":     "Successful is the number of pods which reached Phase Succeeded.",
	"unsuccessful":   "Unsuccessful is the number of pods failures, this applies only to jobs created with RestartPolicyNever, otherwise this value will always be 0.",
}

func (JobStatus) SwaggerDoc() map[string]string {
	return map_JobStatus
}

var map_ReplicationControllerDummy = map[string]string{
	"": "Dummy definition",
}

func (ReplicationControllerDummy) SwaggerDoc() map[string]string {
	return map_ReplicationControllerDummy
}

var map_ResourceConsumption = map[string]string{
	"":         "ResourceConsumption is an object for specifying average resource consumption of a particular resource.",
	"resource": "Resource specifies either the name of the target resource when present in the spec, or the name of the observed resource when present in the status.",
	"quantity": "Quantity specifies either the target average consumption of the resource when present in the spec, or the observed average consumption when present in the status.",
}

func (ResourceConsumption) SwaggerDoc() map[string]string {
	return map_ResourceConsumption
}

var map_RollingUpdateDeployment = map[string]string{
	"":                "Spec to control the desired behavior of rolling update.",
	"maxUnavailable":  "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of total pods at the start of update (ex: 10%). Absolute number is calculated from percentage by rounding up. This can not be 0 if MaxSurge is 0. By default, a fixed value of 1 is used. Example: when this is set to 30%, the old RC can be scaled down by 30% immediately when the rolling update starts. Once new pods are ready, old RC can be scaled down further, followed by scaling up the new RC, ensuring that at least 70% of original number of pods are available at all times during the update.",
	"maxSurge":        "The maximum number of pods that can be scheduled above the original number of pods. Value can be an absolute number (ex: 5) or a percentage of total pods at the start of the update (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. By default, a value of 1 is used. Example: when this is set to 30%, the new RC can be scaled up by 30% immediately when the rolling update starts. Once old pods have been killed, new RC can be scaled up further, ensuring that total number of pods running at any time during the update is atmost 130% of original pods.",
	"minReadySeconds": "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)",
}

func (RollingUpdateDeployment) SwaggerDoc() map[string]string {
	return map_RollingUpdateDeployment
}

var map_Scale = map[string]string{
	"":         "Scale subresource, applicable to ReplicationControllers and (in future) Deployment.",
	"metadata": "Standard object metadata; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata.",
	"spec":     "Spec defines the behavior of the scale. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status.",
	"status":   "Status represents the current status of the scale. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status. Read-only.",
}

func (Scale) SwaggerDoc() map[string]string {
	return map_Scale
}

var map_ScaleSpec = map[string]string{
	"":         "ScaleSpec describes the attributes a Scale subresource",
	"replicas": "Replicas is the number of desired replicas. More info: http://releases.k8s.io/HEAD/docs/user-guide/replication-controller.md#what-is-a-replication-controller\"",
}

func (ScaleSpec) SwaggerDoc() map[string]string {
	return map_ScaleSpec
}

var map_ScaleStatus = map[string]string{
	"":         "ScaleStatus represents the current status of a Scale subresource.",
	"replicas": "Replicas is the number of actual replicas. More info: http://releases.k8s.io/HEAD/docs/user-guide/replication-controller.md#what-is-a-replication-controller",
	"selector": "Selector is a label query over pods that should match the replicas count. If it is empty, it is defaulted to labels on Pod template; More info: http://releases.k8s.io/HEAD/docs/user-guide/labels.md#label-selectors",
}

func (ScaleStatus) SwaggerDoc() map[string]string {
	return map_ScaleStatus
}

var map_SubresourceReference = map[string]string{
	"":            "SubresourceReference contains enough information to let you inspect or modify the referred subresource.",
	"kind":        "Kind of the referent; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds\"",
	"namespace":   "Namespace of the referent; More info: http://releases.k8s.io/HEAD/docs/user-guide/namespaces.md",
	"name":        "Name of the referent; More info: http://releases.k8s.io/HEAD/docs/user-guide/identifiers.md#names",
	"apiVersion":  "API version of the referent",
	"subresource": "Subresource name of the referent",
}

func (SubresourceReference) SwaggerDoc() map[string]string {
	return map_SubresourceReference
}

var map_ThirdPartyResource = map[string]string{
	"":            "A ThirdPartyResource is a generic representation of a resource, it is used by add-ons and plugins to add new resource types to the API.  It consists of one or more Versions of the api.",
	"metadata":    "Standard object metadata",
	"description": "Description is the description of this object.",
	"versions":    "Versions are versions for this third party object",
}

func (ThirdPartyResource) SwaggerDoc() map[string]string {
	return map_ThirdPartyResource
}

var map_ThirdPartyResourceData = map[string]string{
	"":         "An internal object, used for versioned storage in etcd.  Not exposed to the end user.",
	"metadata": "Standard object metadata.",
	"name":     "Data is the raw JSON data for this data.",
}

func (ThirdPartyResourceData) SwaggerDoc() map[string]string {
	return map_ThirdPartyResourceData
}

var map_ThirdPartyResourceDataList = map[string]string{
	"":         "ThirdPartyResrouceDataList is a list of ThirdPartyResourceData.",
	"metadata": "Standard list metadata More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is the list of ThirdpartyResourceData.",
}

func (ThirdPartyResourceDataList) SwaggerDoc() map[string]string {
	return map_ThirdPartyResourceDataList
}

var map_ThirdPartyResourceList = map[string]string{
	"":         "ThirdPartyResourceList is a list of ThirdPartyResources.",
	"metadata": "Standard list metadata.",
	"items":    "Items is the list of ThirdPartyResources.",
}

func (ThirdPartyResourceList) SwaggerDoc() map[string]string {
	return map_ThirdPartyResourceList
}

// AUTO-GENERATED FUNCTIONS END HERE
