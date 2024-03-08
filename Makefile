# The project's root import path
PKG := github.com/shipengqi/commitizen
# set version package
VERSION_PKG=github.com/shipengqi/component-base/version

ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

# set git commit and tree state
GIT_COMMIT = $(shell git rev-parse HEAD)
ifneq ($(shell git status --porcelain 2> /dev/null),)
	GIT_TREE_STATE ?= dirty
else
	GIT_TREE_STATE ?= clean
endif

# set ldflags
GO_LDFLAGS += -X $(VERSION_PKG).Version=$(VERSION) \
	-X $(VERSION_PKG).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PKG).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PKG).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
	
.PHONY: go.build
go.build:
	@echo "===========> Building: $(OUTPUT_DIR)/$(BIN)"
	@CGO_ENABLED=0 go build -ldflags "$(GO_LDFLAGS)" -o $(OUTPUT_DIR)/$(BIN) ${PKG}