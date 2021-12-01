#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

[[ -n ${DEBUG:-} ]] && set -o xtrace

_version="4.15.1"

# Change directories to the parent directory of the one in which this
# script is located.
cd "$(dirname "${BASH_SOURCE[0]}")/.."

source scripts/utils.sh

if command -v jq >/dev/null 2>&1; then exit 0; fi

mkdir -p .local/bin && pushd .local/bin

if [[ ${HOSTOS} == "linux" ]]; then
  _binfile="yq_linux_amd64"
elif [[ ${HOSTOS} == "darwin" ]]; then
  _binfile="yq_darwin_amd64"
fi
_bin_url="https://github.com/mikefarah/yq/releases/download/v${_version}/${_binfile}"
curl -SsL "${_bin_url}" -o yq
chmod 0755 yq
echo "'yq' has been installed to $(pwd), make sure this directory is in your \$PATH"

export PATH=$(pwd):$PATH
yq --version

popd
