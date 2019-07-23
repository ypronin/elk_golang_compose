SHELL := /bin/bash

GO ?= go
BIN ?= bin
BINARY ?= app
BIN_DIR ?= $(join $(dir $(lastword $(MAKEFILE_LIST))), $(BIN))


GIT_REVISION := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git name-rev --name-only HEAD)
COMMIT_ID   ?= $(shell git describe --tags --always --dirty=-dev)
COMMIT_UNIX ?= $(shell git show -s --format=%ct HEAD)
BUILD_UNIX  ?= $(shell date +%s)

BUILD_NUMBER ?=
BUILD = $(BUILD_NUMBER) $(shell git show -s --format='format:%h %aD') $(shell $(GO) version)

VERSION := 1.0.0

LDFLAGS =  -s -w \
	-X 'github.com/elk_golang_compose/pkg/logging.gitRevision=$(GIT_REVISION)' \
	-X 'github.com/elk_golang_compose/pkg/logging.gitBranch=$(GIT_BRANCH)' \
	-X 'github.com/elk_golang_compose/pkg/logging.commitID=$(COMMIT_ID)' \
	-X 'github.com/elk_golang_compose/pkg/logging.commitUnix=$(COMMIT_UNIX)' \
	-X 'github.com/elk_golang_compose/pkg/logging.buildNumber=$(BUILD_NUMBER)' \
	-X 'github.com/elk_golang_compose/pkg/logging.buildUnix=$(BUILD_UNIX)'
LDFLAGS := -ldflags "$(LDFLAGS)"
#LDFLAGS=-ldflags "-X=main.Version=$(VERSION)"

# Build application binary
build:
	@echo "==> building $(BINARY) binary"
	@$(GO) build $(LDFLAGS) -o $(BIN_DIR)/$(BINARY)


validate:
	swagger validate ./swagger.yml
.PHONY: validate

generate: validate
	swagger generate server \
		--target=./swagger \
		--spec=./swagger.yml \
		--exclude-main \
		--name=relax_test


.PHONY: build generate validate