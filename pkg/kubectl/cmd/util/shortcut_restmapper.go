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

package util

import (
	"errors"
	"strings"

	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/kubernetes/pkg/kubectl"
)

// shortcutExpander is a RESTMapper that can be used for Kubernetes resources.   It expands the resource first, then invokes the wrapped
type shortcutExpander struct {
	RESTMapper      meta.RESTMapper
	discoveryClient discovery.DiscoveryInterface
}

var _ meta.RESTMapper = &shortcutExpander{}

type namespacedResources struct{}

func (ns namespacedResources) Match(groupVersion string, r *metav1.APIResource) bool {
	return r.Namespaced
}

func NewShortcutExpander(delegate meta.RESTMapper, client discovery.DiscoveryInterface) (shortcutExpander, error) {
	if client == nil {
		return shortcutExpander{}, errors.New("Please provide discovery client to shortcut expander")
	}
	return shortcutExpander{RESTMapper: delegate, discoveryClient: client}, nil
}

func (e shortcutExpander) KindFor(resource schema.GroupVersionResource) (schema.GroupVersionKind, error) {
	return e.RESTMapper.KindFor(e.expandResourceShortcut(resource))
}

func (e shortcutExpander) KindsFor(resource schema.GroupVersionResource) ([]schema.GroupVersionKind, error) {
	return e.RESTMapper.KindsFor(e.expandResourceShortcut(resource))
}

func (e shortcutExpander) ResourcesFor(resource schema.GroupVersionResource) ([]schema.GroupVersionResource, error) {
	return e.RESTMapper.ResourcesFor(e.expandResourceShortcut(resource))
}

func (e shortcutExpander) ResourceFor(resource schema.GroupVersionResource) (schema.GroupVersionResource, error) {
	return e.RESTMapper.ResourceFor(e.expandResourceShortcut(resource))
}

func (e shortcutExpander) ResourceSingularizer(resource string) (string, error) {
	return e.RESTMapper.ResourceSingularizer(e.expandResourceShortcut(schema.GroupVersionResource{Resource: resource}).Resource)
}

func (e shortcutExpander) RESTMapping(gk schema.GroupKind, versions ...string) (*meta.RESTMapping, error) {
	return e.RESTMapper.RESTMapping(gk, versions...)
}

func (e shortcutExpander) RESTMappings(gk schema.GroupKind, versions ...string) ([]*meta.RESTMapping, error) {
	return e.RESTMapper.RESTMappings(gk, versions...)
}

// UserResources are the resource names that apply to the primary, user facing resources used by
// client tools. They are in deletion-first order - dependent resources should be last.
// Should remain exported in order to expose a current list of resources to downstream
// composition that wants to build on the concept of 'all' for their CLIs.
var UserResources = []schema.GroupResource{
	{Group: "", Resource: "pods"},
	{Group: "", Resource: "replicationcontrollers"},
	{Group: "", Resource: "services"},
	{Group: "apps", Resource: "statefulsets"},
	{Group: "autoscaling", Resource: "horizontalpodautoscalers"},
	{Group: "batch", Resource: "jobs"},
	{Group: "extensions", Resource: "deployments"},
	{Group: "extensions", Resource: "replicasets"},
}

// AliasesForResource returns the aliases for a resource, and whether it has an alias or not
func (e shortcutExpander) AliasesForResource(resource string) ([]string, bool) {
	if strings.ToLower(resource) == "all" {
		// first try to check the user resources wanted by the client that are also known by the server
		userResources := e.filterKnownUserResources(UserResources)

		// if none of the wanted resources is known by the server, try to discover all the ones available there
		if len(userResources) == 0 {
			userResources = e.userResourcesOnServer()
		}

		// as a last resort, just try to use the list of resources initially wanted by the client
		if len(userResources) == 0 {
			userResources = UserResources
		}

		aliases := []string{}
		for _, r := range userResources {
			aliases = append(aliases, r.Resource)
		}
		return aliases, true
	}
	expanded := e.expandResourceShortcut(schema.GroupVersionResource{Resource: resource}).Resource
	return []string{expanded}, (expanded != resource)
}

// filterKnownUserResources takes a list of resource types and return the ones also known by the server.
// This doesn't error out, if we can't access the server the original list is returned.
func (e shortcutExpander) filterKnownUserResources(wantedUserResources []schema.GroupResource) []schema.GroupResource {
	serverResources, err := e.discoveryClient.ServerResources()
	if err != nil {
		return wantedUserResources
	}

	groupedServerResources, err := discovery.GroupVersionResources(serverResources)
	if err != nil {
		return wantedUserResources
	}

	knownUserResources := []schema.GroupResource{}

	// only take the user resources known by the server
	for _, wantedUserResource := range wantedUserResources {
		for serverResource := range groupedServerResources {
			if wantedUserResource.Group == serverResource.Group &&
				wantedUserResource.Resource == serverResource.Resource {
				knownUserResources = append(knownUserResources, wantedUserResource)
				break
			}
		}
	}

	return knownUserResources
}

// userResourcesOnServer tries to discover all user resource types available on the server.
// This doesn't error out, if we can't access the server an empty list is returned.
func (e shortcutExpander) userResourcesOnServer() []schema.GroupResource {
	serverResources, err := e.discoveryClient.ServerResources()
	if err != nil {
		return []schema.GroupResource{}
	}

	// only take the namespaced resources
	namespacedServerResources := discovery.FilteredBy(namespacedResources{}, serverResources)
	namespacedGroupedServerResources, err := discovery.GroupVersionResources(namespacedServerResources)
	if err != nil {
		return []schema.GroupResource{}
	}

	// filters out additional api endpoints for the given resources
	discoveredUserResources := []schema.GroupResource{}
	for resource := range namespacedGroupedServerResources {
		r := resource.GroupResource()
		rSplit := strings.Split(r.Resource, "/")
		if len(rSplit) > 1 {
			continue
		}
		discoveredUserResources = append(discoveredUserResources, r)
	}

	return discoveredUserResources
}

// getShortcutMappings returns a set of tuples which holds short names for resources.
// First the list of potential resources will be taken from the API server.
// Next we will append the hardcoded list of resources - to be backward compatible with old servers.
// NOTE that the list is ordered by group priority.
func (e shortcutExpander) getShortcutMappings() ([]kubectl.ResourceShortcuts, error) {
	res := []kubectl.ResourceShortcuts{}
	// get server resources
	apiResList, err := e.discoveryClient.ServerResources()
	if err == nil {
		for _, apiResources := range apiResList {
			for _, apiRes := range apiResources.APIResources {
				for _, shortName := range apiRes.ShortNames {
					gv, err := schema.ParseGroupVersion(apiResources.GroupVersion)
					if err != nil {
						glog.V(1).Infof("Unable to parse groupversion = %s due to = %s", apiResources.GroupVersion, err.Error())
						continue
					}
					rs := kubectl.ResourceShortcuts{
						ShortForm: schema.GroupResource{Group: gv.Group, Resource: shortName},
						LongForm:  schema.GroupResource{Group: gv.Group, Resource: apiRes.Name},
					}
					res = append(res, rs)
				}
			}
		}
	}

	// append hardcoded short forms at the end of the list
	res = append(res, kubectl.ResourcesShortcutStatic...)
	return res, nil
}

// expandResourceShortcut will return the expanded version of resource
// (something that a pkg/api/meta.RESTMapper can understand), if it is
// indeed a shortcut. If no match has been found, we will match on group prefixing.
// Lastly we will return resource unmodified.
func (e shortcutExpander) expandResourceShortcut(resource schema.GroupVersionResource) schema.GroupVersionResource {
	// get the shortcut mappings and return on first match.
	if resources, err := e.getShortcutMappings(); err == nil {
		for _, item := range resources {
			if len(resource.Group) != 0 && resource.Group != item.ShortForm.Group {
				continue
			}
			if resource.Resource == item.ShortForm.Resource {
				resource.Resource = item.LongForm.Resource
				return resource
			}
		}

		// we didn't find exact match so match on group prefixing. This allows autoscal to match autoscaling
		if len(resource.Group) == 0 {
			return resource
		}
		for _, item := range resources {
			if !strings.HasPrefix(item.ShortForm.Group, resource.Group) {
				continue
			}
			if resource.Resource == item.ShortForm.Resource {
				resource.Resource = item.LongForm.Resource
				return resource
			}
		}
	}

	return resource
}
