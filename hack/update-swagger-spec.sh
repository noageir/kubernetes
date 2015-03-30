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

# Script to fetch latest swagger spec.
# Puts the updated spec at swagger-spec/

set -o errexit
set -o nounset
set -o pipefail

LMKTFY_ROOT=$(dirname "${BASH_SOURCE}")/..
SWAGGER_ROOT_DIR="${LMKTFY_ROOT}/api/swagger-spec"
source "${LMKTFY_ROOT}/hack/lib/init.sh"

function cleanup()
{
    [[ -n ${APISERVER_PID-} ]] && kill ${APISERVER_PID} 1>&2 2>/dev/null

    lmktfy::etcd::cleanup

    lmktfy::log::status "Clean up complete"
}

trap cleanup EXIT SIGINT

lmktfy::etcd::start

ETCD_HOST=${ETCD_HOST:-127.0.0.1}
ETCD_PORT=${ETCD_PORT:-4001}
API_PORT=${API_PORT:-8050}
API_HOST=${API_HOST:-127.0.0.1}
LMKTFYLET_PORT=${LMKTFYLET_PORT:-10250}

# Start lmktfy-apiserver
lmktfy::log::status "Starting lmktfy-apiserver"
"${LMKTFY_OUTPUT_HOSTBIN}/lmktfy-apiserver" \
  --address="127.0.0.1" \
  --public_address_override="127.0.0.1" \
  --port="${API_PORT}" \
  --etcd_servers="http://${ETCD_HOST}:${ETCD_PORT}" \
  --public_address_override="127.0.0.1" \
  --lmktfylet_port=${LMKTFYLET_PORT} \
  --runtime_config=api/v1beta3 \
  --portal_net="10.0.0.0/24" 1>&2 &
APISERVER_PID=$!

lmktfy::util::wait_for_url "http://127.0.0.1:${API_PORT}/healthz" "apiserver: "

SWAGGER_API_PATH="http://127.0.0.1:${API_PORT}/swaggerapi/"
lmktfy::log::status "Updating " ${SWAGGER_ROOT_DIR}
curl ${SWAGGER_API_PATH} > ${SWAGGER_ROOT_DIR}/resourceListing.json
curl ${SWAGGER_API_PATH}version > ${SWAGGER_ROOT_DIR}/version.json
curl ${SWAGGER_API_PATH}api > ${SWAGGER_ROOT_DIR}/api.json
curl ${SWAGGER_API_PATH}api/v1beta1 > ${SWAGGER_ROOT_DIR}/v1beta1.json
curl ${SWAGGER_API_PATH}api/v1beta2 > ${SWAGGER_ROOT_DIR}/v1beta2.json
curl ${SWAGGER_API_PATH}api/v1beta3 > ${SWAGGER_ROOT_DIR}/v1beta3.json

lmktfy::log::status "SUCCESS"
