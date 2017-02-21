/*
Copyright 2014 The Kubernetes Authors.

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

package rest

import (
	"strings"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	expapi "k8s.io/kubernetes/pkg/apis/extensions"
)

type FakeAPIInterface struct {
	removed   []string
	installed []*expapi.ThirdPartyResource
	apis      []string
	t         *testing.T
}

func (f *FakeAPIInterface) RemoveThirdPartyResource(path string) error {
	f.removed = append(f.removed, path)
	return nil
}

func (f *FakeAPIInterface) InstallThirdPartyResource(rsrc *expapi.ThirdPartyResource) error {
	f.installed = append(f.installed, rsrc)
	f.apis = append(f.apis, MakeThirdPartyPath(rsrc.Spec.Group))
	return nil
}

func (f *FakeAPIInterface) HasThirdPartyResource(rsrc *expapi.ThirdPartyResource) (bool, error) {
	if f.apis == nil {
		return false, nil
	}
	path := MakeThirdPartyPath(rsrc.Spec.Group)
	for _, api := range f.apis {
		if api == path {
			return true, nil
		}
	}
	return false, nil
}

func (f *FakeAPIInterface) ListThirdPartyResources() []string {
	return f.apis
}

func TestSyncAPIs(t *testing.T) {
	resourcesNamed := func(names ...string) []expapi.ThirdPartyResource {
		result := []expapi.ThirdPartyResource{}
		for _, name := range names {
			parts := strings.SplitN(name, ".", 2)
			result = append(result, expapi.ThirdPartyResource{
				ObjectMeta: metav1.ObjectMeta{Name: name},
				Spec: expapi.ThirdPartyResourceSpec{
					Group:    parts[1],
					Version:  "v1",
					Resource: parts[0],
				},
			})
		}
		return result
	}

	tests := []struct {
		list              *expapi.ThirdPartyResourceList
		apis              []string
		expectedInstalled []string
		expectedRemoved   []string
		name              string
	}{
		{
			list: &expapi.ThirdPartyResourceList{
				Items: resourcesNamed("foo.example.com"),
			},
			expectedInstalled: []string{"foo.example.com"},
			name:              "simple add",
		},
		{
			list: &expapi.ThirdPartyResourceList{
				Items: resourcesNamed("foo.example.com"),
			},
			apis: []string{
				"/apis/example.com",
				"/apis/example.com/v1",
			},
			name: "does nothing",
		},
		{
			list: &expapi.ThirdPartyResourceList{
				Items: resourcesNamed("foo.example.com"),
			},
			apis: []string{
				"/apis/example.com",
				"/apis/example.com/v1",
				"/apis/example.co",
				"/apis/example.co/v1",
			},
			name: "deletes substring API",
			expectedRemoved: []string{
				"/apis/example.co",
				"/apis/example.co/v1",
			},
		},
		{
			list: &expapi.ThirdPartyResourceList{
				Items: resourcesNamed("foo.example.com", "foo.company.com"),
			},
			apis: []string{
				"/apis/company.com",
				"/apis/company.com/v1",
			},
			expectedInstalled: []string{"foo.example.com"},
			name:              "adds with existing",
		},
		{
			list: &expapi.ThirdPartyResourceList{
				Items: resourcesNamed("foo.example.com"),
			},
			apis: []string{
				"/apis/company.com",
				"/apis/company.com/v1",
			},
			expectedInstalled: []string{"foo.example.com"},
			expectedRemoved:   []string{"/apis/company.com", "/apis/company.com/v1"},
			name:              "removes with existing",
		},
	}

	for _, test := range tests {
		fake := FakeAPIInterface{
			apis: test.apis,
			t:    t,
		}

		cntrl := ThirdPartyController{master: &fake}

		if err := cntrl.syncResourceList(test.list); err != nil {
			t.Errorf("[%s] unexpected error: %v", test.name, err)
		}
		if len(test.expectedInstalled) != len(fake.installed) {
			t.Errorf("[%s] unexpected installed APIs: %d, expected %d (%#v)", test.name, len(fake.installed), len(test.expectedInstalled), fake.installed[0])
			continue
		} else {
			names := sets.String{}
			for ix := range fake.installed {
				names.Insert(fake.installed[ix].Name)
			}
			for _, name := range test.expectedInstalled {
				if !names.Has(name) {
					t.Errorf("[%s] missing installed API: %s", test.name, name)
				}
			}
		}
		if len(test.expectedRemoved) != len(fake.removed) {
			t.Errorf("[%s] unexpected installed APIs: %d, expected %d", test.name, len(fake.removed), len(test.expectedRemoved))
			continue
		} else {
			names := sets.String{}
			names.Insert(fake.removed...)
			for _, name := range test.expectedRemoved {
				if !names.Has(name) {
					t.Errorf("[%s] missing removed API: %s (%s)", test.name, name, names)
				}
			}
		}
	}
}
