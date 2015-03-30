#!/bin/bash

# Copyright 2015 Google Inc. All rights reserved.
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

set -o errexit
set -o nounset
set -o pipefail

LMKTFY_ROOT=$(dirname "${BASH_SOURCE}")/../..

: ${LMKTFY_VERSION_ROOT:=${LMKTFY_ROOT}}
: ${LMKTFYCTL:="${LMKTFY_VERSION_ROOT}/cluster/lmktfyctl.sh"}
: ${LMKTFY_CONFIG_FILE:="config-test.sh"}

export LMKTFYCTL LMKTFY_CONFIG_FILE

source "${LMKTFY_ROOT}/cluster/lmktfy-env.sh"
source "${LMKTFY_VERSION_ROOT}/cluster/${LMKTFYRNETES_PROVIDER}/util.sh"

prepare-e2e

test-teardown
