#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

[[ -n ${DEBUG:-} ]] && set -o xtrace

GH_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
cd "${GH_ROOT}/github-sync" || exit 1

if ! command -v jq >/dev/null 2>&1; then echo "jq command is not available"; exit 1; fi
if ! command -v yq >/dev/null 2>&1; then echo "jq command is not available"; exit 1; fi

sort_repositories() {
    yq eval -o=json github-data/repositories.yaml | jq '.repositories |= sort_by(.name) | .repositories[].collaborators |= sort_by(.username)' | yq eval -P - > github-data/repositories-sorted.yaml || exit 1 ; \
    mv github-data/repositories-sorted.yaml github-data/repositories.yaml
}

sort_teams() {
    yq eval -o=json github-data/teams.yaml | jq '.teams |= sort_by(.name)' | yq eval -P - > github-data/teams-sorted.yaml || exit 1 ; \
    mv github-data/teams-sorted.yaml github-data/teams.yaml
}

sort_rest() {
    # NOTE: all that is not repositories.yaml or teams.yaml is considered sortable by .username
    yaml_files=$(find github-data -type f -name "*.yaml" ! -name '*repositories*' ! -name '*teams*' | sort -u)
    for f in ${yaml_files}
    do
        (yq eval -o=json "$$f" | jq '.[] |= sort_by(.username)' | yq eval -P - > "$$f-sorted") && mv "$$f-sorted" "$$f"  || exit 1;
    done
}

sort_repositories
sort_teams
sort_rest