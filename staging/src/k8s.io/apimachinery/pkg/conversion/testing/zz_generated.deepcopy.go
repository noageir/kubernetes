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

package testing

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	reflect "reflect"
)

// Deprecated: GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TestStruct).DeepCopyInto(out.(*TestStruct))
			return nil
		}, InType: reflect.TypeOf(&TestStruct{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TestSubStruct).DeepCopyInto(out.(*TestSubStruct))
			return nil
		}, InType: reflect.TypeOf(&TestSubStruct{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TestSubSubStruct).DeepCopyInto(out.(*TestSubSubStruct))
			return nil
		}, InType: reflect.TypeOf(&TestSubSubStruct{})},
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestStruct) DeepCopyInto(out *TestStruct) {
	*out = *in
	if in.Map != nil {
		in, out := &in.Map, &out.Map
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Pointer != nil {
		in, out := &in.Pointer, &out.Pointer
		if *in == nil {
			*out = nil
		} else {
			*out = new(int)
			**out = **in
		}
	}
	in.Struct.DeepCopyInto(&out.Struct)
	if in.StructPointer != nil {
		in, out := &in.StructPointer, &out.StructPointer
		if *in == nil {
			*out = nil
		} else {
			*out = new(TestSubStruct)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.StructSlice != nil {
		in, out := &in.StructSlice, &out.StructSlice
		*out = make([]*TestSubStruct, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(TestSubStruct)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.StructMap != nil {
		in, out := &in.StructMap, &out.StructMap
		*out = make(map[string]*TestSubStruct, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(TestSubStruct)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new TestStruct.
func (x *TestStruct) DeepCopy() *TestStruct {
	if x == nil {
		return nil
	}
	out := new(TestStruct)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyTestObject is an autogenerated deepcopy function, copying the receiver, creating a new TestObject.
func (x *TestStruct) DeepCopyTestObject() TestObject {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSubStruct) DeepCopyInto(out *TestSubStruct) {
	*out = *in
	in.A.DeepCopyInto(&out.A)
	in.B.DeepCopyInto(&out.B)
	in.C.DeepCopyInto(&out.C)
	if in.X != nil {
		in, out := &in.X, &out.X
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.Y != nil {
		in, out := &in.Y, &out.Y
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new TestSubStruct.
func (x *TestSubStruct) DeepCopy() *TestSubStruct {
	if x == nil {
		return nil
	}
	out := new(TestSubStruct)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSubSubStruct) DeepCopyInto(out *TestSubSubStruct) {
	*out = *in
	if in.D != nil {
		in, out := &in.D, &out.D
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new TestSubSubStruct.
func (x *TestSubSubStruct) DeepCopy() *TestSubSubStruct {
	if x == nil {
		return nil
	}
	out := new(TestSubSubStruct)
	x.DeepCopyInto(out)
	return out
}
