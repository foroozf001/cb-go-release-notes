MAKEFLAGS += --silent
SHELL:=/bin/bash -o pipefail -o errexit
.ONESHELL:
.SHELLFLAGS:=-ec

NAME="releasenotes"
BINARIES='go'

.PHONY: requirements
requirements:
	for item in $(BINARIES) ; do \
		command -v $$item ; \
	done

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: fmt
fmt: ## Format and lint the source code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: run
run: fmt vet build ## Run from your host.
	bin/logger-service

##@ Build

.PHONY: build
build: requirements fmt vet ## Build binary.
	go build -o bin/"${NAME}"-$(shell cat VERSION) cmd/releasenotes/*.go
