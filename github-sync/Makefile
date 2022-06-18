.DEFAULT_GOAL:=help
SHELL:=/usr/bin/env bash

COLOR:=\\033[36m
NOCOLOR:=\\033[0m

build: ## Build mattermost-gitops
	go build

## --------------------------------------
## Sort YAML
## --------------------------------------

.PHONY: yaml-sort sort-repo sort-teams
yaml_files = $(shell find github-data -type f -name "*.yaml" ! -name '*repositories*' ! -name '*teams*' | sort -u)
yaml-sort: sort-repo sort-teams ## Sort all YAML files alphabetically by the username key
	@for f in $(yaml_files); do (yq eval -o=json "$$f" | jq '.[] |= sort_by(.username)' | yq eval -P - > "$$f-new") && mv "$$f-new" "$$f"  || exit 1; done; \

sort-repo:
	yq eval -o=json github-data/repositories.yaml | jq '.repositories |= sort_by(.name) | .repositories[].collaborators |= sort_by(.username)' | yq eval -P - > github-data/repositories-new.yaml || exit 1 ; \
	mv github-data/repositories-new.yaml github-data/repositories.yaml

sort-teams:
	yq eval -o=json github-data/teams.yaml | jq '.teams |= sort_by(.name)' | yq eval -P - > github-data/teams-new.yaml || exit 1 ; \
	mv github-data/teams-new.yaml github-data/teams.yaml

##@ Helpers

.PHONY: help

help:  ## Display this help
	@awk \
		-v "col=${COLOR}" -v "nocol=${NOCOLOR}" \
		' \
			BEGIN { \
				FS = ":.*##" ; \
				printf "\nUsage:\n  make %s<target>%s\n", col, nocol \
			} \
			/^[a-zA-Z_-]+:.*?##/ { \
				printf "  %s%-15s%s %s\n", col, $$1, nocol, $$2 \
			} \
			/^##@/ { \
				printf "\n%s%s%s\n", col, substr($$0, 5), nocol \
			} \
		' $(MAKEFILE_LIST)
