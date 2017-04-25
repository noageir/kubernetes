/*
Copyright 2015 The Kubernetes Authors.

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

package helper

import (
	"reflect"
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/kubernetes/pkg/api/v1"
)

func TestAddToNodeAddresses(t *testing.T) {
	testCases := []struct {
		existing []v1.NodeAddress
		toAdd    []v1.NodeAddress
		expected []v1.NodeAddress
	}{
		{
			existing: []v1.NodeAddress{},
			toAdd:    []v1.NodeAddress{},
			expected: []v1.NodeAddress{},
		},
		{
			existing: []v1.NodeAddress{},
			toAdd: []v1.NodeAddress{
				{Type: v1.NodeExternalIP, Address: "1.1.1.1"},
				{Type: v1.NodeHostName, Address: "localhost"},
			},
			expected: []v1.NodeAddress{
				{Type: v1.NodeExternalIP, Address: "1.1.1.1"},
				{Type: v1.NodeHostName, Address: "localhost"},
			},
		},
		{
			existing: []v1.NodeAddress{},
			toAdd: []v1.NodeAddress{
				{Type: v1.NodeExternalIP, Address: "1.1.1.1"},
				{Type: v1.NodeExternalIP, Address: "1.1.1.1"},
			},
			expected: []v1.NodeAddress{
				{Type: v1.NodeExternalIP, Address: "1.1.1.1"},
			},
		},
		{
			existing: []v1.NodeAddress{
				{Type: v1.NodeExternalIP, Address: "1.1.1.1"},
				{Type: v1.NodeInternalIP, Address: "10.1.1.1"},
			},
			toAdd: []v1.NodeAddress{
				{Type: v1.NodeExternalIP, Address: "1.1.1.1"},
				{Type: v1.NodeHostName, Address: "localhost"},
			},
			expected: []v1.NodeAddress{
				{Type: v1.NodeExternalIP, Address: "1.1.1.1"},
				{Type: v1.NodeInternalIP, Address: "10.1.1.1"},
				{Type: v1.NodeHostName, Address: "localhost"},
			},
		},
	}

	for i, tc := range testCases {
		AddToNodeAddresses(&tc.existing, tc.toAdd...)
		if !apiequality.Semantic.DeepEqual(tc.expected, tc.existing) {
			t.Errorf("case[%d], expected: %v, got: %v", i, tc.expected, tc.existing)
		}
	}
}

func TestGetAccessModesFromString(t *testing.T) {
	modes := GetAccessModesFromString("ROX")
	if !containsAccessMode(modes, v1.ReadOnlyMany) {
		t.Errorf("Expected mode %s, but got %+v", v1.ReadOnlyMany, modes)
	}

	modes = GetAccessModesFromString("ROX,RWX")
	if !containsAccessMode(modes, v1.ReadOnlyMany) {
		t.Errorf("Expected mode %s, but got %+v", v1.ReadOnlyMany, modes)
	}
	if !containsAccessMode(modes, v1.ReadWriteMany) {
		t.Errorf("Expected mode %s, but got %+v", v1.ReadWriteMany, modes)
	}

	modes = GetAccessModesFromString("RWO,ROX,RWX")
	if !containsAccessMode(modes, v1.ReadOnlyMany) {
		t.Errorf("Expected mode %s, but got %+v", v1.ReadOnlyMany, modes)
	}
	if !containsAccessMode(modes, v1.ReadWriteMany) {
		t.Errorf("Expected mode %s, but got %+v", v1.ReadWriteMany, modes)
	}
}

func TestRemoveDuplicateAccessModes(t *testing.T) {
	modes := []v1.PersistentVolumeAccessMode{
		v1.ReadWriteOnce, v1.ReadOnlyMany, v1.ReadOnlyMany, v1.ReadOnlyMany,
	}
	modes = removeDuplicateAccessModes(modes)
	if len(modes) != 2 {
		t.Errorf("Expected 2 distinct modes in set but found %v", len(modes))
	}
}

func TestNodeSelectorRequirementsAsSelector(t *testing.T) {
	matchExpressions := []v1.NodeSelectorRequirement{{
		Key:      "foo",
		Operator: v1.NodeSelectorOpIn,
		Values:   []string{"bar", "baz"},
	}}
	mustParse := func(s string) labels.Selector {
		out, e := labels.Parse(s)
		if e != nil {
			panic(e)
		}
		return out
	}
	tc := []struct {
		in        []v1.NodeSelectorRequirement
		out       labels.Selector
		expectErr bool
	}{
		{in: nil, out: labels.Nothing()},
		{in: []v1.NodeSelectorRequirement{}, out: labels.Nothing()},
		{
			in:  matchExpressions,
			out: mustParse("foo in (baz,bar)"),
		},
		{
			in: []v1.NodeSelectorRequirement{{
				Key:      "foo",
				Operator: v1.NodeSelectorOpExists,
				Values:   []string{"bar", "baz"},
			}},
			expectErr: true,
		},
		{
			in: []v1.NodeSelectorRequirement{{
				Key:      "foo",
				Operator: v1.NodeSelectorOpGt,
				Values:   []string{"1"},
			}},
			out: mustParse("foo>1"),
		},
		{
			in: []v1.NodeSelectorRequirement{{
				Key:      "bar",
				Operator: v1.NodeSelectorOpLt,
				Values:   []string{"7"},
			}},
			out: mustParse("bar<7"),
		},
	}

	for i, tc := range tc {
		out, err := NodeSelectorRequirementsAsSelector(tc.in)
		if err == nil && tc.expectErr {
			t.Errorf("[%v]expected error but got none.", i)
		}
		if err != nil && !tc.expectErr {
			t.Errorf("[%v]did not expect error but got: %v", i, err)
		}
		if !reflect.DeepEqual(out, tc.out) {
			t.Errorf("[%v]expected:\n\t%+v\nbut got:\n\t%+v", i, tc.out, out)
		}
	}
}

func TestTolerationsTolerateTaintsWithFilter(t *testing.T) {
	testCases := []struct {
		description     string
		tolerations     []v1.Toleration
		taints          []v1.Taint
		applyFilter     taintsFilterFunc
		expectTolerated bool
	}{
		{
			description:     "empty tolerations tolerate empty taints",
			tolerations:     []v1.Toleration{},
			taints:          []v1.Taint{},
			applyFilter:     func(t *v1.Taint) bool { return true },
			expectTolerated: true,
		},
		{
			description: "non-empty tolerations tolerate empty taints",
			tolerations: []v1.Toleration{
				{
					Key:      "foo",
					Operator: "Exists",
					Effect:   v1.TaintEffectNoSchedule,
				},
			},
			taints:          []v1.Taint{},
			applyFilter:     func(t *v1.Taint) bool { return true },
			expectTolerated: true,
		},
		{
			description: "tolerations match all taints, expect tolerated",
			tolerations: []v1.Toleration{
				{
					Key:      "foo",
					Operator: "Exists",
					Effect:   v1.TaintEffectNoSchedule,
				},
			},
			taints: []v1.Taint{
				{
					Key:    "foo",
					Effect: v1.TaintEffectNoSchedule,
				},
			},
			applyFilter:     func(t *v1.Taint) bool { return true },
			expectTolerated: true,
		},
		{
			description: "tolerations don't match taints, but no taints apply to the filter, expect tolerated",
			tolerations: []v1.Toleration{
				{
					Key:      "foo",
					Operator: "Exists",
					Effect:   v1.TaintEffectNoSchedule,
				},
			},
			taints: []v1.Taint{
				{
					Key:    "bar",
					Effect: v1.TaintEffectNoSchedule,
				},
			},
			applyFilter:     func(t *v1.Taint) bool { return false },
			expectTolerated: true,
		},
		{
			description: "no filterFunc indicated, means all taints apply to the filter, tolerations don't match taints, expect untolerated",
			tolerations: []v1.Toleration{
				{
					Key:      "foo",
					Operator: "Exists",
					Effect:   v1.TaintEffectNoSchedule,
				},
			},
			taints: []v1.Taint{
				{
					Key:    "bar",
					Effect: v1.TaintEffectNoSchedule,
				},
			},
			applyFilter:     nil,
			expectTolerated: false,
		},
		{
			description: "tolerations match taints, expect tolerated",
			tolerations: []v1.Toleration{
				{
					Key:      "foo",
					Operator: "Exists",
					Effect:   v1.TaintEffectNoExecute,
				},
			},
			taints: []v1.Taint{
				{
					Key:    "foo",
					Effect: v1.TaintEffectNoExecute,
				},
				{
					Key:    "bar",
					Effect: v1.TaintEffectNoSchedule,
				},
			},
			applyFilter:     func(t *v1.Taint) bool { return t.Effect == v1.TaintEffectNoExecute },
			expectTolerated: true,
		},
	}

	for _, tc := range testCases {
		if tc.expectTolerated != TolerationsTolerateTaintsWithFilter(tc.tolerations, tc.taints, tc.applyFilter) {
			filteredTaints := []v1.Taint{}
			for _, taint := range tc.taints {
				if tc.applyFilter != nil && !tc.applyFilter(&taint) {
					continue
				}
				filteredTaints = append(filteredTaints, taint)
			}
			t.Errorf("[%s] expect tolerations %+v tolerate filtered taints %+v in taints %+v", tc.description, tc.tolerations, filteredTaints, tc.taints)
		}
	}
}

func TestGetAvoidPodsFromNode(t *testing.T) {
	controllerFlag := true
	testCases := []struct {
		node        *v1.Node
		expectValue v1.AvoidPods
		expectErr   bool
	}{
		{
			node:        &v1.Node{},
			expectValue: v1.AvoidPods{},
			expectErr:   false,
		},
		{
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						v1.PreferAvoidPodsAnnotationKey: `
							{
							    "preferAvoidPods": [
							        {
							            "podSignature": {
							                "podController": {
						                            "apiVersion": "v1",
						                            "kind": "ReplicationController",
						                            "name": "foo",
						                            "uid": "abcdef123456",
						                            "controller": true
							                }
							            },
							            "reason": "some reason",
							            "message": "some message"
							        }
							    ]
							}`,
					},
				},
			},
			expectValue: v1.AvoidPods{
				PreferAvoidPods: []v1.PreferAvoidPodsEntry{
					{
						PodSignature: v1.PodSignature{
							PodController: &metav1.OwnerReference{
								APIVersion: "v1",
								Kind:       "ReplicationController",
								Name:       "foo",
								UID:        "abcdef123456",
								Controller: &controllerFlag,
							},
						},
						Reason:  "some reason",
						Message: "some message",
					},
				},
			},
			expectErr: false,
		},
		{
			node: &v1.Node{
				// Missing end symbol of "podController" and "podSignature"
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						v1.PreferAvoidPodsAnnotationKey: `
							{
							    "preferAvoidPods": [
							        {
							            "podSignature": {
							                "podController": {
							                    "kind": "ReplicationController",
							                    "apiVersion": "v1"
							            "reason": "some reason",
							            "message": "some message"
							        }
							    ]
							}`,
					},
				},
			},
			expectValue: v1.AvoidPods{},
			expectErr:   true,
		},
	}

	for i, tc := range testCases {
		v, err := GetAvoidPodsFromNodeAnnotations(tc.node.Annotations)
		if err == nil && tc.expectErr {
			t.Errorf("[%v]expected error but got none.", i)
		}
		if err != nil && !tc.expectErr {
			t.Errorf("[%v]did not expect error but got: %v", i, err)
		}
		if !reflect.DeepEqual(tc.expectValue, v) {
			t.Errorf("[%v]expect value %v but got %v with %v", i, tc.expectValue, v, v.PreferAvoidPods[0].PodSignature.PodController.Controller)
		}
	}
}

func TestSysctlsFromPodAnnotation(t *testing.T) {
	type Test struct {
		annotation  string
		expectValue []v1.Sysctl
		expectErr   bool
	}
	for i, test := range []Test{
		{
			annotation:  "",
			expectValue: nil,
		},
		{
			annotation: "foo.bar",
			expectErr:  true,
		},
		{
			annotation: "=123",
			expectErr:  true,
		},
		{
			annotation:  "foo.bar=",
			expectValue: []v1.Sysctl{{Name: "foo.bar", Value: ""}},
		},
		{
			annotation:  "foo.bar=42",
			expectValue: []v1.Sysctl{{Name: "foo.bar", Value: "42"}},
		},
		{
			annotation: "foo.bar=42,",
			expectErr:  true,
		},
		{
			annotation:  "foo.bar=42,abc.def=1",
			expectValue: []v1.Sysctl{{Name: "foo.bar", Value: "42"}, {Name: "abc.def", Value: "1"}},
		},
	} {
		sysctls, err := SysctlsFromPodAnnotation(test.annotation)
		if test.expectErr && err == nil {
			t.Errorf("[%v]expected error but got none", i)
		} else if !test.expectErr && err != nil {
			t.Errorf("[%v]did not expect error but got: %v", i, err)
		} else if !reflect.DeepEqual(sysctls, test.expectValue) {
			t.Errorf("[%v]expect value %v but got %v", i, test.expectValue, sysctls)
		}
	}
}

// TODO: remove when alpha support for affinity is removed
func TestGetAffinityFromPodAnnotations(t *testing.T) {
	testCases := []struct {
		pod       *v1.Pod
		expectErr bool
	}{
		{
			pod:       &v1.Pod{},
			expectErr: false,
		},
		{
			pod: &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						v1.AffinityAnnotationKey: `
						{"nodeAffinity": { "requiredDuringSchedulingIgnoredDuringExecution": {
							"nodeSelectorTerms": [{
								"matchExpressions": [{
									"key": "foo",
									"operator": "In",
									"values": ["value1", "value2"]
								}]
							}]
						}}}`,
					},
				},
			},
			expectErr: false,
		},
		{
			pod: &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						v1.AffinityAnnotationKey: `
						{"nodeAffinity": { "requiredDuringSchedulingIgnoredDuringExecution": {
							"nodeSelectorTerms": [{
								"matchExpressions": [{
									"key": "foo",
						`,
					},
				},
			},
			expectErr: true,
		},
	}

	for i, tc := range testCases {
		_, err := GetAffinityFromPodAnnotations(tc.pod.Annotations)
		if err == nil && tc.expectErr {
			t.Errorf("[%v]expected error but got none.", i)
		}
		if err != nil && !tc.expectErr {
			t.Errorf("[%v]did not expect error but got: %v", i, err)
		}
	}
}

func TestGetPodQOS(t *testing.T) {
	testCases := []struct {
		pod      *v1.Pod
		expected v1.PodQOSClass
	}{
		{
			pod: newPod("guaranteed", []v1.Container{
				newContainer("guaranteed", getResourceList("100m", "100Mi"), getResourceList("100m", "100Mi")),
			}),
			expected: v1.PodQOSGuaranteed,
		},
		{
			pod: newPod("guaranteed-with-gpu", []v1.Container{
				newContainer("guaranteed", getResourceList("100m", "100Mi"), addResource("nvidia-gpu", "2", getResourceList("100m", "100Mi"))),
			}),
			expected: v1.PodQOSGuaranteed,
		},
		{
			pod: newPod("guaranteed-guaranteed", []v1.Container{
				newContainer("guaranteed", getResourceList("100m", "100Mi"), getResourceList("100m", "100Mi")),
				newContainer("guaranteed", getResourceList("100m", "100Mi"), getResourceList("100m", "100Mi")),
			}),
			expected: v1.PodQOSGuaranteed,
		},
		{
			pod: newPod("guaranteed-guaranteed-with-gpu", []v1.Container{
				newContainer("guaranteed", getResourceList("100m", "100Mi"), addResource("nvidia-gpu", "2", getResourceList("100m", "100Mi"))),
				newContainer("guaranteed", getResourceList("100m", "100Mi"), getResourceList("100m", "100Mi")),
			}),
			expected: v1.PodQOSGuaranteed,
		},
		{
			pod: newPod("best-effort-best-effort", []v1.Container{
				newContainer("best-effort", getResourceList("", ""), getResourceList("", "")),
				newContainer("best-effort", getResourceList("", ""), getResourceList("", "")),
			}),
			expected: v1.PodQOSBestEffort,
		},
		{
			pod: newPod("best-effort-best-effort-with-gpu", []v1.Container{
				newContainer("best-effort", getResourceList("", ""), addResource("nvidia-gpu", "2", getResourceList("", ""))),
				newContainer("best-effort", getResourceList("", ""), getResourceList("", "")),
			}),
			expected: v1.PodQOSBestEffort,
		},
		{
			pod: newPod("best-effort-with-gpu", []v1.Container{
				newContainer("best-effort", getResourceList("", ""), addResource("nvidia-gpu", "2", getResourceList("", ""))),
			}),
			expected: v1.PodQOSBestEffort,
		},
		{
			pod: newPod("best-effort-burstable", []v1.Container{
				newContainer("best-effort", getResourceList("", ""), addResource("nvidia-gpu", "2", getResourceList("", ""))),
				newContainer("burstable", getResourceList("1", ""), getResourceList("2", "")),
			}),
			expected: v1.PodQOSBurstable,
		},
		{
			pod: newPod("best-effort-guaranteed", []v1.Container{
				newContainer("best-effort", getResourceList("", ""), addResource("nvidia-gpu", "2", getResourceList("", ""))),
				newContainer("guaranteed", getResourceList("10m", "100Mi"), getResourceList("10m", "100Mi")),
			}),
			expected: v1.PodQOSBurstable,
		},
		{
			pod: newPod("burstable-cpu-guaranteed-memory", []v1.Container{
				newContainer("burstable", getResourceList("", "100Mi"), getResourceList("", "100Mi")),
			}),
			expected: v1.PodQOSBurstable,
		},
		{
			pod: newPod("burstable-no-limits", []v1.Container{
				newContainer("burstable", getResourceList("100m", "100Mi"), getResourceList("", "")),
			}),
			expected: v1.PodQOSBurstable,
		},
		{
			pod: newPod("burstable-guaranteed", []v1.Container{
				newContainer("burstable", getResourceList("1", "100Mi"), getResourceList("2", "100Mi")),
				newContainer("guaranteed", getResourceList("100m", "100Mi"), getResourceList("100m", "100Mi")),
			}),
			expected: v1.PodQOSBurstable,
		},
		{
			pod: newPod("burstable-unbounded-but-requests-match-limits", []v1.Container{
				newContainer("burstable", getResourceList("100m", "100Mi"), getResourceList("200m", "200Mi")),
				newContainer("burstable-unbounded", getResourceList("100m", "100Mi"), getResourceList("", "")),
			}),
			expected: v1.PodQOSBurstable,
		},
		{
			pod: newPod("burstable-1", []v1.Container{
				newContainer("burstable", getResourceList("10m", "100Mi"), getResourceList("100m", "200Mi")),
			}),
			expected: v1.PodQOSBurstable,
		},
		{
			pod: newPod("burstable-2", []v1.Container{
				newContainer("burstable", getResourceList("0", "0"), addResource("nvidia-gpu", "2", getResourceList("100m", "200Mi"))),
			}),
			expected: v1.PodQOSBurstable,
		},
	}
	for id, testCase := range testCases {
		if actual := GetPodQOS(testCase.pod); testCase.expected != actual {
			t.Errorf("[%d]: invalid qos pod %s, expected: %s, actual: %s", id, testCase.pod.Name, testCase.expected, actual)
		}
	}
}

func getResourceList(cpu, memory string) v1.ResourceList {
	res := v1.ResourceList{}
	if cpu != "" {
		res[v1.ResourceCPU] = resource.MustParse(cpu)
	}
	if memory != "" {
		res[v1.ResourceMemory] = resource.MustParse(memory)
	}
	return res
}

func addResource(rName, value string, rl v1.ResourceList) v1.ResourceList {
	rl[v1.ResourceName(rName)] = resource.MustParse(value)
	return rl
}

func getResourceRequirements(requests, limits v1.ResourceList) v1.ResourceRequirements {
	res := v1.ResourceRequirements{}
	res.Requests = requests
	res.Limits = limits
	return res
}

func newContainer(name string, requests v1.ResourceList, limits v1.ResourceList) v1.Container {
	return v1.Container{
		Name:      name,
		Resources: getResourceRequirements(requests, limits),
	}
}

func newPod(name string, containers []v1.Container) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1.PodSpec{
			Containers: containers,
		},
	}
}
