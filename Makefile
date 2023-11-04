MAKEFLAGS += --silent
SHELL:=/bin/bash -o pipefail -o errexit
.ONESHELL:
.SHELLFLAGS:=-ec

APP="changelog"
BINARIES='go'
GOOS=$(shell uname -s | tr '[:upper:]' '[:lower:]')
GOARCH=$(shell uname -m)

.PHONY: requirements
requirements:
	for item in $(BINARIES) ; do \
		command -v $$item ; \
	done

.PHONY: help
help: ## Displays help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: fmt
fmt: ## Formats the source code
	go fmt ./...

.PHONY: vet
vet: ## Validates the source code
	go vet ./...

.PHONY: run
run: fmt vet ## Runs the source code
	go run cmd/"${APP}"/*.go "${NEW_VERSION}" "${OLD_VERSION}" "./"

.PHONY: build
build: requirements fmt vet ## Builds the binary
	GOOS="${GOOS}" GOARCH="${GOARCH}" go build -o bin/$(shell cat VERSION)/"${GOOS}_${GOARCH}"/"${APP}" cmd/"${APP}"/*.go

.PHONY: changelog
changelog: ## Generates changelogs
	if [[ -f bin/$(shell cat VERSION)/"${GOOS}_${GOARCH}"/"${APP}" ]]; then \
		bin/$(shell cat VERSION)/"${GOOS}_${GOARCH}"/"${APP}" "${NEW_VERSION}" "${OLD_VERSION}" . ; \
	else \
		echo 'missing changelog binary';\
	fi