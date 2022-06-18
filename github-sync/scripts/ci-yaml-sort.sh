#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

[[ -n ${DEBUG:-} ]] && set -o xtrace

GH_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
cd "${GH_ROOT}" || exit 1

# shellcheck source=ensure-yq.sh
source "./scripts/ensure-yq.sh"
# shellcheck source=ensure-jq.sh
source "./scripts/ensure-jq.sh"

export PATH=.local/bin:$PATH

cleanup() {
    returnCode="$?"
    rm -f "*-new"
    exit "${returnCode}"
}

trap cleanup EXIT

find github-data -type f -name "*.yaml" ! -name '*repositories*' ! -name '*teams*' | sort -u

yaml_files=$(find github-data -type f -name "*.yaml" ! -name '*repositories*' ! -name '*teams*' | sort -u)
for f in ${yaml_files}
do
  yq eval -o=json "${f}" | jq '.[] |= sort_by(.username)' | yq eval -P - > "${f}-new"
  if ! diff "${f}-new" "${f}" >> /dev/null; then
    echo "yaml files are not sorted!! Please sort them with \"make yaml-sort\" before commit"
    echo "Unsorted file: ${f}"
    exit 1
  fi
done

yq eval -o=json github-data/repositories.yaml | jq '.repositories |= sort_by(.name) | .repositories[].collaborators |= sort_by(.username)' | yq eval -P - > github-data/repositories.yaml-new
if ! diff github-data/repositories.yaml-new github-data/repositories.yaml >> /dev/null; then
  echo "yaml files are not sorted!! Please sort them with \"make yaml-sort\" before commit"
  echo "Unsorted file: github-data/repositories.yaml"
  exit 1
fi

yq eval -o=json github-data/teams.yaml | jq '.teams |= sort_by(.name)' | yq eval -P - > github-data/teams.yaml-new
if ! diff github-data/teams.yaml-new github-data/teams.yaml >> /dev/null; then
  echo "yaml files are not sorted!! Please sort them with \"make yaml-sort\" before commit"
  echo "Unsorted file: github-data/teams.yaml"
  exit 1
fi
