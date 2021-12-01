#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

[[ -n ${DEBUG:-} ]] && set -o xtrace

_version="1.6" # earlier versions don't follow the same OS/ARCH patterns

# Change directories to the parent directory of the one in which this
# script is located.
cd "$(dirname "${BASH_SOURCE[0]}")/.."

source scripts/utils.sh

if command -v jq >/dev/null 2>&1; then exit 0; fi

mkdir -p .local/bin && pushd .local/bin

if [[ ${HOSTOS} == "linux" ]]; then
  _binfile="jq-linux64"
elif [[ ${HOSTOS} == "darwin" ]]; then
  _binfile="jq-osx-amd64"
fi
_bin_url="https://github.com/stedolan/jq/releases/download/jq-${_version}/${_binfile}"
curl -SsL "${_bin_url}" -o jq
chmod 0755 jq
echo "'jq' has been installed to $(pwd), make sure this directory is in your \$PATH"

export PATH=$(pwd):$PATH
jq --version

popd
