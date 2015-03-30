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

# A library of helper functions and constant for the local config.

# Use the config file specified in $LMKTFY_CONFIG_FILE, or default to
# config-default.sh.

LMKTFY_ROOT=$(dirname "${BASH_SOURCE}")/../..
source "${LMKTFY_ROOT}/cluster/gke/${LMKTFY_CONFIG_FILE:-config-default.sh}"

# Perform preparations required to run e2e tests
#
# Assumed vars:
#   GCLOUD
function prepare-e2e() {
  echo "... in prepare-e2e()" >&2

  # Ensure GCLOUD is set to some gcloud binary.
  if [[ -z "${GCLOUD:-}" ]]; then
    echo "GCLOUD environment variable is not set. It should be your gcloud binary. " >&2
    echo "A sane default is probably \$ export GCLOUD=gcloud" >&2
    exit 1
  fi
}


# Use the gcloud defaults to find the project.  If it is already set in the
# environment then go with that.
#
# Assumed vars:
#   GCLOUD
# Vars set:
#   PROJECT
function detect-project() {
  echo "... in detect-project()" >&2
  if [[ -z "${PROJECT:-}" ]]; then
    export PROJECT=$("${GCLOUD}" config list project | tail -n 1 | cut -f 3 -d ' ')
  fi

  if [[ -z "${PROJECT:-}" ]]; then
    echo "Could not detect Google Cloud Platform project. Set the default project using " >&2
    echo "'gcloud config set project <PROJECT>'" >&2
    exit 1
  fi
  echo "Project: ${PROJECT}" >&2
}

# Execute prior to running tests to build a release if required for env.
function test-build-release() {
  echo "... in test-build-release()" >&2
  # We currently use the LMKTFY version that GKE supports (not testing
  # bleeding-edge builds).
}

# Verify needed binaries exist.
function verify-prereqs() {
  echo "... in verify-prereqs()" >&2

  ${GCLOUD} preview --help >/dev/null || {
    echo "Either the GCLOUD environment variable is wrong, or the 'preview' component"
    echo "is not installed. (Fix with 'gcloud components update preview')"
  }
}

# Instantiate a lmktfy cluster
#
# Assumed vars:
#   GCLOUD
#   CLUSTER_NAME
#   ZONE
#   CLUSTER_API_VERSION (optional)
#   NUM_MINIONS
function lmktfy-up() {
  echo "... in lmktfy-up()" >&2
  detect-project >&2

  # Make the specified network if we need to.
  if ! gcloud compute networks describe "${NETWORK}" &>/dev/null; then
    echo "Creating new network: ${NETWORK}" >&2
    gcloud compute networks create "${NETWORK}" --project="${PROJECT}" --range "${NETWORK_RANGE}"
  else
    echo "Using network: ${NETWORK}" >&2
  fi

  # Allow SSH on all nodes in the network. This doesn't actually check whether
  # such a rule exists, only whether we've created this exact rule.
  if ! gcloud compute firewall-rules describe "${FIREWALL_SSH}" &>/dev/null; then
    echo "Creating new firewall for SSH: ${FIREWALL_SSH}" >&2
    gcloud compute firewall-rules create "${FIREWALL_SSH}" \
      --allow="tcp:22" \
      --network="${NETWORK}" \
      --project="${PROJECT}" \
      --source-ranges="0.0.0.0/0"
  else
    echo "Using firewall-rule: ${FIREWALL_SSH}" >&2
  fi

  # Bring up the cluster.
  "${GCLOUD}" preview container clusters create "${CLUSTER_NAME}" \
    --zone="${ZONE}" \
    --project="${PROJECT}" \
    --cluster-api-version="${CLUSTER_API_VERSION:-}" \
    --num-nodes="${NUM_MINIONS}" \
    --network="${NETWORK}"
}

# Execute prior to running tests to initialize required structure. This is
# called from hack/e2e-go only when running -up (it is run after lmktfy-up, so
# the cluster already exists at this point).
#
# Assumed vars:
#   CLUSTER_NAME
#   GCLOUD
# Vars set:
#   MINION_TAG
function test-setup() {
  echo "... in test-setup()" >&2
  # Detect the project into $PROJECT if it isn't set
  detect-project >&2

  # At this point, CLUSTER_NAME should have been used, so its value is final.
  MINION_TAG="lmktfy-${CLUSTER_NAME}-node"

  # Open up port 80 & 8080 so common containers on minions can be reached.
  # TODO(mbforbes): Is adding ${USER} necessary, and sufficient, to avoid
  #                 collisions here?
  "${GCLOUD}" compute firewall-rules create \
    "${MINION_TAG}-${USER}-http-alt" \
    --allow tcp:80 tcp:8080 \
    --project "${PROJECT}" \
    --target-tags "${MINION_TAG}" \
    --network="${NETWORK}"
}

# Ensure that we have a password created for validating to the master.
#
# Assumed vars:
#  ZONE
#  CLUSTER_NAME
# Vars set:
#   LMKTFY_USER
#   LMKTFY_PASSWORD
function get-password() {
  echo "... in get-password()" >&2
  detect-project >&2
  LMKTFY_USER=$("${GCLOUD}" preview container clusters describe \
    --project="${PROJECT}" --zone="${ZONE}" "${CLUSTER_NAME}" \
    | grep user | cut -f 4 -d ' ')
  LMKTFY_PASSWORD=$("${GCLOUD}" preview container clusters describe \
    --project="${PROJECT}" --zone="${ZONE}" "${CLUSTER_NAME}" \
    | grep password | cut -f 4 -d ' ')
}

# Detect the instance name and IP for the master
#
# Assumed vars:
#   ZONE
#   CLUSTER_NAME
# Vars set:
#   LMKTFY_MASTER
#   LMKTFY_MASTER_IP
function detect-master() {
  echo "... in detect-master()" >&2
  detect-project >&2
  LMKTFY_MASTER="lmktfy-${CLUSTER_NAME}-master"
  LMKTFY_MASTER_IP=$("${GCLOUD}" preview container clusters describe \
    --project="${PROJECT}" --zone="${ZONE}" "${CLUSTER_NAME}" \
    | grep endpoint | cut -f 2 -d ' ')
}

# Assumed vars:
#   NUM_MINIONS
#   CLUSTER_NAME
# Vars set:
#   (none)
function detect-minions() {
  echo "... in detect-minions()" >&2
}

# Detect minions created in the minion group
#
# Assumed vars:
#   none
# Vars set:
#   MINION_NAMES
function detect-minion-names {
  detect-project
  export MINION_NAMES=""
  count=$("${GCLOUD}" preview container clusters describe --project="${PROJECT}" --zone="${ZONE}" "${CLUSTER_NAME}" | grep numNodes | cut -f 2 -d ' ')
  for x in $(seq 1 $count); do
    export MINION_NAMES="${MINION_NAMES} lmktfy-${CLUSTER_NAME}-node-${x} ";
  done
  MINION_NAMES=(${MINION_NAMES})
  echo "MINION_NAMES=${MINION_NAMES[*]}"
}

# SSH to a node by name ($1) and run a command ($2).
#
# Assumed vars:
#   GCLOUD
#   ZONE
function ssh-to-node() {
  echo "... in ssh-to-node()" >&2
  detect-project >&2

  local node="$1"
  local cmd="$2"
  "${GCLOUD}" compute ssh --ssh-flag="-o LogLevel=quiet" --project "${PROJECT}" \
    --zone="${ZONE}" "${node}" --command "${cmd}"
}

# Restart the lmktfy-proxy on a node ($1)
function restart-lmktfy-proxy() {
  echo "... in restart-lmktfy-proxy()"  >&2
  ssh-to-node "$1" "sudo /etc/init.d/lmktfy-proxy restart"
}

# Restart the lmktfy-proxy on master ($1)
function restart-apiserver() {
  echo "... in restart-lmktfy-apiserver()"  >&2
  ssh-to-node "$1" "sudo /etc/init.d/lmktfy-apiserver restart"
}

# Execute after running tests to perform any required clean-up.  This is called
# from hack/e2e-test.sh. This calls lmktfy-down, so the cluster still exists when
# this is called.
#
# Assumed vars:
#   CLUSTER_NAME
#   GCLOUD
#   LMKTFY_ROOT
function test-teardown() {
  echo "... in test-teardown()" >&2

  detect-project >&2
  # At this point, CLUSTER_NAME should have been used, so its value is final.
  MINION_TAG="lmktfy-${CLUSTER_NAME}-node"

  # First, remove anything we did with test-setup (currently, the firewall).
  # NOTE: Keep in sync with name above in test-setup.
  "${GCLOUD}" compute firewall-rules delete "${MINION_TAG}-${USER}-http-alt" \
    --project="${PROJECT}" || true

  # Then actually turn down the cluster.
  "${LMKTFY_ROOT}/cluster/lmktfy-down.sh"
}

# Actually take down the cluster. This is called from test-teardown.
#
# Assumed vars:
#  GCLOUD
#  ZONE
#  CLUSTER_NAME
function lmktfy-down() {
  echo "... in lmktfy-down()" >&2
  detect-project >&2
  "${GCLOUD}" preview container clusters delete --project="${PROJECT}" \
    --zone="${ZONE}" "${CLUSTER_NAME}"
}
