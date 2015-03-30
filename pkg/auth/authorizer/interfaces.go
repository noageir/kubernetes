/*
Copyright 2014 Google Inc. All rights reserved.

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

package authorizer

import (
	"github.com/GoogleCloudPlatform/lmktfy/pkg/auth/user"
)

// Attributes is an interface used by an Authorizer to get information about a request
// that is used to make an authorization decision.
type Attributes interface {
	// The user string which the request was authenticated as, or empty if
	// no authentication occured and the request was allowed to proceed.
	GetUserName() string

	// The list of group names the authenticated user is a member of. Can be
	// empty if the authenticated user is not in any groups, or if no
	// authentication occurred.
	GetGroups() []string

	// When IsReadOnly() == true, the request has no side effects, other than
	// caching, logging, and other incidentals.
	IsReadOnly() bool

	// The namespace of the object, if a request is for a REST object.
	GetNamespace() string

	// The kind of object, if a request is for a REST object.
	GetResource() string
}

// Authorizer makes an authorization decision based on information gained by making
// zero or more calls to methods of the Attributes interface.  It returns nil when an action is
// authorized, otherwise it returns an error.
type Authorizer interface {
	Authorize(a Attributes) (err error)
}

// AttributesRecord implements Attributes interface.
type AttributesRecord struct {
	User      user.Info
	ReadOnly  bool
	Namespace string
	Resource  string
}

func (a AttributesRecord) GetUserName() string {
	return a.User.GetName()
}

func (a AttributesRecord) GetGroups() []string {
	return a.User.GetGroups()
}

func (a AttributesRecord) IsReadOnly() bool {
	return a.ReadOnly
}

func (a AttributesRecord) GetNamespace() string {
	return a.Namespace
}

func (a AttributesRecord) GetResource() string {
	return a.Resource
}
