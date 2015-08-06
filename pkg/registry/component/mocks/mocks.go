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

// This file was generated by counterfeiter
package mocks

import (
	"sync"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/component"
	"k8s.io/kubernetes/pkg/watch"
)

type MockRegistry struct {
	ListComponentsStub        func(api.Context, labels.Selector, fields.Selector) (*api.ComponentList, error)
	listComponentsMutex       sync.RWMutex
	listComponentsArgsForCall []struct {
		arg1 api.Context
		arg2 labels.Selector
		arg3 fields.Selector
	}
	listComponentsReturns struct {
		result1 *api.ComponentList
		result2 error
	}
	GetComponentStub        func(ctx api.Context, name string) (*api.Component, error)
	getComponentMutex       sync.RWMutex
	getComponentArgsForCall []struct {
		ctx  api.Context
		name string
	}
	getComponentReturns struct {
		result1 *api.Component
		result2 error
	}
	CreateComponentStub        func(api.Context, *api.Component) (*api.Component, error)
	createComponentMutex       sync.RWMutex
	createComponentArgsForCall []struct {
		arg1 api.Context
		arg2 *api.Component
	}
	createComponentReturns struct {
		result1 *api.Component
		result2 error
	}
	UpdateComponentStub        func(api.Context, *api.Component) (*api.Component, error)
	updateComponentMutex       sync.RWMutex
	updateComponentArgsForCall []struct {
		arg1 api.Context
		arg2 *api.Component
	}
	updateComponentReturns struct {
		result1 *api.Component
		result2 error
	}
	DeleteComponentStub        func(ctx api.Context, name string, options *api.DeleteOptions) error
	deleteComponentMutex       sync.RWMutex
	deleteComponentArgsForCall []struct {
		ctx     api.Context
		name    string
		options *api.DeleteOptions
	}
	deleteComponentReturns struct {
		result1 error
	}
	WatchComponentsStub        func(ctx api.Context, lSelector labels.Selector, fSelector fields.Selector, resourceVersion string) (watch.Interface, error)
	watchComponentsMutex       sync.RWMutex
	watchComponentsArgsForCall []struct {
		ctx             api.Context
		lSelector       labels.Selector
		fSelector       fields.Selector
		resourceVersion string
	}
	watchComponentsReturns struct {
		result1 watch.Interface
		result2 error
	}
}

func (fake *MockRegistry) ListComponents(arg1 api.Context, arg2 labels.Selector, arg3 fields.Selector) (*api.ComponentList, error) {
	fake.listComponentsMutex.Lock()
	fake.listComponentsArgsForCall = append(fake.listComponentsArgsForCall, struct {
		arg1 api.Context
		arg2 labels.Selector
		arg3 fields.Selector
	}{arg1, arg2, arg3})
	fake.listComponentsMutex.Unlock()
	if fake.ListComponentsStub != nil {
		return fake.ListComponentsStub(arg1, arg2, arg3)
	} else {
		return fake.listComponentsReturns.result1, fake.listComponentsReturns.result2
	}
}

func (fake *MockRegistry) ListComponentsCallCount() int {
	fake.listComponentsMutex.RLock()
	defer fake.listComponentsMutex.RUnlock()
	return len(fake.listComponentsArgsForCall)
}

func (fake *MockRegistry) ListComponentsArgsForCall(i int) (api.Context, labels.Selector, fields.Selector) {
	fake.listComponentsMutex.RLock()
	defer fake.listComponentsMutex.RUnlock()
	return fake.listComponentsArgsForCall[i].arg1, fake.listComponentsArgsForCall[i].arg2, fake.listComponentsArgsForCall[i].arg3
}

func (fake *MockRegistry) ListComponentsReturns(result1 *api.ComponentList, result2 error) {
	fake.ListComponentsStub = nil
	fake.listComponentsReturns = struct {
		result1 *api.ComponentList
		result2 error
	}{result1, result2}
}

func (fake *MockRegistry) GetComponent(ctx api.Context, name string) (*api.Component, error) {
	fake.getComponentMutex.Lock()
	fake.getComponentArgsForCall = append(fake.getComponentArgsForCall, struct {
		ctx  api.Context
		name string
	}{ctx, name})
	fake.getComponentMutex.Unlock()
	if fake.GetComponentStub != nil {
		return fake.GetComponentStub(ctx, name)
	} else {
		return fake.getComponentReturns.result1, fake.getComponentReturns.result2
	}
}

func (fake *MockRegistry) GetComponentCallCount() int {
	fake.getComponentMutex.RLock()
	defer fake.getComponentMutex.RUnlock()
	return len(fake.getComponentArgsForCall)
}

func (fake *MockRegistry) GetComponentArgsForCall(i int) (api.Context, string) {
	fake.getComponentMutex.RLock()
	defer fake.getComponentMutex.RUnlock()
	return fake.getComponentArgsForCall[i].ctx, fake.getComponentArgsForCall[i].name
}

func (fake *MockRegistry) GetComponentReturns(result1 *api.Component, result2 error) {
	fake.GetComponentStub = nil
	fake.getComponentReturns = struct {
		result1 *api.Component
		result2 error
	}{result1, result2}
}

func (fake *MockRegistry) CreateComponent(arg1 api.Context, arg2 *api.Component) (*api.Component, error) {
	fake.createComponentMutex.Lock()
	fake.createComponentArgsForCall = append(fake.createComponentArgsForCall, struct {
		arg1 api.Context
		arg2 *api.Component
	}{arg1, arg2})
	fake.createComponentMutex.Unlock()
	if fake.CreateComponentStub != nil {
		return fake.CreateComponentStub(arg1, arg2)
	} else {
		return fake.createComponentReturns.result1, fake.createComponentReturns.result2
	}
}

func (fake *MockRegistry) CreateComponentCallCount() int {
	fake.createComponentMutex.RLock()
	defer fake.createComponentMutex.RUnlock()
	return len(fake.createComponentArgsForCall)
}

func (fake *MockRegistry) CreateComponentArgsForCall(i int) (api.Context, *api.Component) {
	fake.createComponentMutex.RLock()
	defer fake.createComponentMutex.RUnlock()
	return fake.createComponentArgsForCall[i].arg1, fake.createComponentArgsForCall[i].arg2
}

func (fake *MockRegistry) CreateComponentReturns(result1 *api.Component, result2 error) {
	fake.CreateComponentStub = nil
	fake.createComponentReturns = struct {
		result1 *api.Component
		result2 error
	}{result1, result2}
}

func (fake *MockRegistry) UpdateComponent(arg1 api.Context, arg2 *api.Component) (*api.Component, error) {
	fake.updateComponentMutex.Lock()
	fake.updateComponentArgsForCall = append(fake.updateComponentArgsForCall, struct {
		arg1 api.Context
		arg2 *api.Component
	}{arg1, arg2})
	fake.updateComponentMutex.Unlock()
	if fake.UpdateComponentStub != nil {
		return fake.UpdateComponentStub(arg1, arg2)
	} else {
		return fake.updateComponentReturns.result1, fake.updateComponentReturns.result2
	}
}

func (fake *MockRegistry) UpdateComponentCallCount() int {
	fake.updateComponentMutex.RLock()
	defer fake.updateComponentMutex.RUnlock()
	return len(fake.updateComponentArgsForCall)
}

func (fake *MockRegistry) UpdateComponentArgsForCall(i int) (api.Context, *api.Component) {
	fake.updateComponentMutex.RLock()
	defer fake.updateComponentMutex.RUnlock()
	return fake.updateComponentArgsForCall[i].arg1, fake.updateComponentArgsForCall[i].arg2
}

func (fake *MockRegistry) UpdateComponentReturns(result1 *api.Component, result2 error) {
	fake.UpdateComponentStub = nil
	fake.updateComponentReturns = struct {
		result1 *api.Component
		result2 error
	}{result1, result2}
}

func (fake *MockRegistry) DeleteComponent(ctx api.Context, name string, options *api.DeleteOptions) error {
	fake.deleteComponentMutex.Lock()
	fake.deleteComponentArgsForCall = append(fake.deleteComponentArgsForCall, struct {
		ctx     api.Context
		name    string
		options *api.DeleteOptions
	}{ctx, name, options})
	fake.deleteComponentMutex.Unlock()
	if fake.DeleteComponentStub != nil {
		return fake.DeleteComponentStub(ctx, name, options)
	} else {
		return fake.deleteComponentReturns.result1
	}
}

func (fake *MockRegistry) DeleteComponentCallCount() int {
	fake.deleteComponentMutex.RLock()
	defer fake.deleteComponentMutex.RUnlock()
	return len(fake.deleteComponentArgsForCall)
}

func (fake *MockRegistry) DeleteComponentArgsForCall(i int) (api.Context, string, *api.DeleteOptions) {
	fake.deleteComponentMutex.RLock()
	defer fake.deleteComponentMutex.RUnlock()
	return fake.deleteComponentArgsForCall[i].ctx, fake.deleteComponentArgsForCall[i].name, fake.deleteComponentArgsForCall[i].options
}

func (fake *MockRegistry) DeleteComponentReturns(result1 error) {
	fake.DeleteComponentStub = nil
	fake.deleteComponentReturns = struct {
		result1 error
	}{result1}
}

func (fake *MockRegistry) WatchComponents(ctx api.Context, lSelector labels.Selector, fSelector fields.Selector, resourceVersion string) (watch.Interface, error) {
	fake.watchComponentsMutex.Lock()
	fake.watchComponentsArgsForCall = append(fake.watchComponentsArgsForCall, struct {
		ctx             api.Context
		lSelector       labels.Selector
		fSelector       fields.Selector
		resourceVersion string
	}{ctx, lSelector, fSelector, resourceVersion})
	fake.watchComponentsMutex.Unlock()
	if fake.WatchComponentsStub != nil {
		return fake.WatchComponentsStub(ctx, lSelector, fSelector, resourceVersion)
	} else {
		return fake.watchComponentsReturns.result1, fake.watchComponentsReturns.result2
	}
}

func (fake *MockRegistry) WatchComponentsCallCount() int {
	fake.watchComponentsMutex.RLock()
	defer fake.watchComponentsMutex.RUnlock()
	return len(fake.watchComponentsArgsForCall)
}

func (fake *MockRegistry) WatchComponentsArgsForCall(i int) (api.Context, labels.Selector, fields.Selector, string) {
	fake.watchComponentsMutex.RLock()
	defer fake.watchComponentsMutex.RUnlock()
	return fake.watchComponentsArgsForCall[i].ctx, fake.watchComponentsArgsForCall[i].lSelector, fake.watchComponentsArgsForCall[i].fSelector, fake.watchComponentsArgsForCall[i].resourceVersion
}

func (fake *MockRegistry) WatchComponentsReturns(result1 watch.Interface, result2 error) {
	fake.WatchComponentsStub = nil
	fake.watchComponentsReturns = struct {
		result1 watch.Interface
		result2 error
	}{result1, result2}
}

var _ component.Registry = new(MockRegistry)
