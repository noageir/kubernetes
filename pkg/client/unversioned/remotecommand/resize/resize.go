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

package resize

import (
	"encoding/json"
	"io"

	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/util/term"
)

// GetResizeFunc will return function that handles terminal resize
func GetResizeFunc(resizeQueue term.TerminalSizeQueue) func(io.Writer) {
	return func(stream io.Writer) {
		defer runtime.HandleCrash()

		encoder := json.NewEncoder(stream)
		for {
			size := resizeQueue.Next()
			if size == nil {
				return
			}
			if err := encoder.Encode(&size); err != nil {
				runtime.HandleError(err)
			}
		}
	}
}
