#!/bin/bash

# Copyright 2014 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Clean out the output directory on the docker host.
set -o errexit
set -o nounset
set -o pipefail

LMKTFY_ROOT=$(dirname "${BASH_SOURCE}")/..
source "$LMKTFY_ROOT/build/common.sh"

lmktfy::build::verify_prereqs clean
lmktfy::build::clean_output
lmktfy::build::clean_images
