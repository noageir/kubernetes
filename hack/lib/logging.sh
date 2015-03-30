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

# Handler for when we exit automatically on an error.
# Borrowed from https://gist.github.com/ahendrix/7030300
lmktfy::log::errexit() {
  local err="${PIPESTATUS[@]}"

  # If the shell we are in doesn't have errexit set (common in subshells) then
  # don't dump stacks.
  set +o | grep -qe "-o errexit" || return

  set +o xtrace
  local code="${1:-1}"
  lmktfy::log::error_exit "'${BASH_COMMAND}' exited with status $err" "${1:-1}" 1
}

lmktfy::log::install_errexit() {
  # trap ERR to provide an error handler whenever a command exits nonzero  this
  # is a more verbose version of set -o errexit
  trap 'lmktfy::log::errexit' ERR

  # setting errtrace allows our ERR trap handler to be propagated to functions,
  # expansions and subshells
  set -o errtrace
}

# Print out the stack trace
#
# Args:
#   $1 The number of stack frames to skip when printing.
lmktfy::log::stack() {
  local stack_skip=${1:-0}
  stack_skip=$((stack_skip + 1))
  if [[ ${#FUNCNAME[@]} -gt $stack_skip ]]; then
    echo "Call stack:" >&2
    local i
    for ((i=1 ; i <= ${#FUNCNAME[@]} - $stack_skip ; i++))
    do
      local frame_no=$((i - 1 + stack_skip))
      local source_file=${BASH_SOURCE[$frame_no]}
      local source_lineno=${BASH_LINENO[$((frame_no - 1))]}
      local funcname=${FUNCNAME[$frame_no]}
      echo "  $i: ${source_file}:${source_lineno} ${funcname}(...)" >&2
    done
  fi
}

# Log an error and exit.
# Args:
#   $1 Message to log with the error
#   $2 The error code to return
#   $3 The number of stack frames to skip when printing.
lmktfy::log::error_exit() {
  local message="${1:-}"
  local code="${2:-1}"
  local stack_skip="${3:-0}"
  stack_skip=$((stack_skip + 1))

  local source_file=${BASH_SOURCE[$stack_skip]}
  local source_line=${BASH_LINENO[$((stack_skip - 1))]}
  echo "!!! Error in ${source_file}:${source_line}" >&2
  [[ -z ${1-} ]] || {
    echo "  ${1}" >&2
  }

  lmktfy::log::stack $stack_skip

  echo "Exiting with status ${code}" >&2
  exit "${code}"
}

# Log an error but keep going.  Don't dump the stack or exit.
lmktfy::log::error() {
  timestamp=$(date +"[%m%d %H:%M:%S]")
  echo "!!! $timestamp ${1-}" >&2
  shift
  for message; do
    echo "    $message" >&2
  done
}

# Print an usage message to stderr.  The arguments are printed directly.
lmktfy::log::usage() {
  echo >&2
  local message
  for message; do
    echo "$message" >&2
  done
  echo >&2
}

lmktfy::log::usage_from_stdin() {
  local messages=()
  while read -r line; do
    messages+=$line
  done

  lmktfy::log::usage "${messages[@]}"
}

# Print out some info that isn't a top level status line
lmktfy::log::info() {
  for message; do
    echo "$message"
  done
}

lmktfy::log::info_from_stdin() {
  local messages=()
  while read -r line; do
    messages+=$line
  done

  lmktfy::log::info "${messages[@]}"
}

# Print a status line.  Formatted to show up in a stream of output.
lmktfy::log::status() {
  timestamp=$(date +"[%m%d %H:%M:%S]")
  echo "+++ $timestamp $1"
  shift
  for message; do
    echo "    $message"
  done
}
