.DEFAULT_GOAL 	= help

SHELL         	= bash
project       	= reward
GIT_AUTHOR    	= janosmiko

help: ## Outputs this help screen
	@grep -E '(^[\/a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

# If the first argument is "gen"...
ifeq (gen,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  GEN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(GEN_ARGS):;@:)
endif

## —— Commands —————————————————————————————————————————————————————————
.PHONY: remove-jsonld
remove-jsonld: ## Remove JSON-LD files
	python3 hack/main.py --input api.yaml --output reward-api.yaml

.PHONY: generate
generate: ## Generate Go code from OpenAPI spec
	docker run --rm -v "$${PWD}:/local" openapitools/openapi-generator-cli:v6.5.0 generate -i /local/reward-api.yaml -c /local/config.yaml -g go --global-property apiTests=false,apiDocs=false,modelTests=false,modelDocs=false -o /local/rewardcloud
	rm -f rewardcloud/go.mod && rm -f rewardcloud/go.sum

.PHONY: convert
convert: ## Convert json to yaml
	yq api.json -oy > api.yaml

.PHONY: all
all: ## Generate Go code from OpenAPI spec
	make convert
	make remove-jsonld
	make generate