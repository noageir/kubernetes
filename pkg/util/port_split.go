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

package util

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Takes a string of the form "name:port" or "name".
//  * If id is of the form "name" or "name:", then return (name, "", true)
//  * If id is of the form "name:port", then return (name, port, true)
//  * Otherwise, return ("", "", false)
// Additionally, name must be non-empty or valid will be returned false.
//
// Port is returned as a string, and it is not required to be numeric (could be
// used for a named port, for example).
func SplitPort(id string) (name, port string, valid bool) {
	parts := strings.Split(id, ":")
	if len(parts) > 2 {
		return "", "", false
	}
	if len(parts) == 2 {
		return parts[0], parts[1], len(parts[0]) > 0
	}
	return id, "", len(id) > 0
}

// ParsePort parses a raw host string of the form "host:port".
// Will error if port is absent, non-numeric, or out of allowed port range.
func ParsePort(rawhost string) (host string, port int, err error) {
	var portString string
	if host, portString, err = net.SplitHostPort(rawhost); err != nil {
		return "", 0, fmt.Errorf("failed to parse port from host (%s): %v", rawhost, err)
	}
	if port, err = strconv.Atoi(portString); err != nil {
		return "", 0, fmt.Errorf("failed to parse port as int from host (%s): %v", rawhost, err)
	}
	if port < 0 || port > 0xFFFF {
		return "", 0, fmt.Errorf("failed to parse port from host (%s): port out of range (%d)", rawhost, port)
	}
	return host, port, nil
}
