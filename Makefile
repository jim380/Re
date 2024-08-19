#!/usr/bin/make -f

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell git describe --tags 2>/dev/null | sed 's/^v//')?
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
TM_VERSION := $(shell go list -m github.com/cometbft/cometbft | sed 's:.* ::')
BINDIR ?= $(CURDIR)/build
DOCKER := $(shell which docker)

SIMAPP = ./app
ENABLED_PROPOSALS := SudoContract,UpdateAdmin,ClearAdmin,PinCodes,UnpinCodes
GO_VERSION=1.22
BUILDDIR ?= $(CURDIR)/build

export GO111MODULE = on

GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

build_tags_test_binary = $(build_tags)
build_tags_test_binary += skip_ccv_msg_filter

whitespace :=
empty = $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(empty),$(comma),$(build_tags))
build_tags_test_binary_comma_sep := $(subst $(empty),$(comma),$(build_tags_test_binary))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=re \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=red \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
		  -X "github.com/jim380/re/app.EnableSpecificProposals=$(ENABLED_PROPOSALS)"

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags_comma_sep)" -ldflags '$(ldflags)' -trimpath
BUILD_FLAGS_TEST_BINARY := -tags "$(build_tags_test_binary_comma_sep)" -ldflags '$(ldflags)' -trimpath

# The below include contains the tools and runsim targets.
include contrib/devtools/Makefile

check_go_version:
	@echo "Go version: $(GO_MAJOR_VERSION).$(GO_MINOR_VERSION)"
ifneq ($(GO_MINOR_VERSION),22)
	@echo "ERROR: Go version 1.22 is required for this version of Re"
	exit 1
endif

all: install lint test

build: check_go_version
ifeq ($(OS),Windows_NT)
	exit 1
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/red ./cmd/red
endif

build-static-linux-amd64: go.sum $(BUILDDIR)/
	$(DOCKER) buildx create --name rebuilder || true
	$(DOCKER) buildx use rebuilder
	$(DOCKER) buildx build \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--build-arg GIT_VERSION=$(VERSION) \
		--build-arg GIT_COMMIT=$(COMMIT) \
		--build-arg BUILD_TAGS=$(build_tags_comma_sep) \
		--build-arg ENABLED_PROPOSALS=$(ENABLED_PROPOSALS) \
		--platform linux/amd64 \
		-t re-amd64 \
		--load \
		-f Dockerfile.builder .
	$(DOCKER) rm -f rebinary || true
	$(DOCKER) create -ti --name rebinary re-amd64
	$(DOCKER) cp rebinary:/bin/red $(BUILDDIR)/red-linux-amd64
	$(DOCKER) rm -f rebinary

install: check_go_version
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/red

install-test-binary: check_go_version
	go install -mod=readonly $(BUILD_FLAGS_TEST_BINARY) ./cmd/red

########################################
### Tools & dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

draw-deps:
	@# requires brew install graphviz or apt-get install graphviz
	go get github.com/RobotsAndPencils/goviz
	@goviz -i ./cmd/red -d 2 | dot -Tpng -o dependency-graph.png

clean:
	rm -rf snapcraft-local.yaml build/

distclean: clean
	rm -rf vendor/

########################################
### Testing


test: test-unit
test-all: check test-race test-cover

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./...

test-race:
	@VERSION=$(VERSION) go test -mod=readonly -race -tags='ledger test_ledger_mock' ./...

test-cover:
	@go test -mod=readonly -timeout 30m -race -coverprofile=coverage.txt -covermode=atomic -tags='ledger test_ledger_mock' ./...

benchmark:
	@go test -mod=readonly -bench=. ./...

test-sim-import-export: runsim
	@echo "Running application import/export simulation. This may take several minutes..."
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 50 5 TestAppImportExport

test-sim-multi-seed-short: runsim
	@echo "Running short multi-seed application simulation. This may take awhile!"
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 50 10 TestFullAppSimulation

###############################################################################
###                                Linting                                  ###
###############################################################################

lint:
	golangci-lint run 
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*_test.go" | xargs gofmt -d -s

format:
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name "*.pb.go" -not -name "*.pb.gw.go" -not -name "*.pulsar.go" -not -path "./crypto/keys/secp256k1/*" | xargs gofumpt -w -l
	golangci-lint run --fix
.PHONY: format

###############################################################################
###                                Protobuf                                 ###
###############################################################################
protoVer=0.14.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)
PROTO_FORMATTER_IMAGE=tendermintdev/docker-build-proto@sha256:aabcfe2fc19c31c0f198d4cd26393f5e5ca9502d7ea3feafbfe972448fee7cae

proto-all: proto-format proto-lint proto-gen format

proto-gen:
	@$(protoImage) sh ./scripts/protocgen.sh
	@go mod tidy
	
proto-format:
	@echo "Formatting Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace \
	--workdir /workspace $(PROTO_FORMATTER_IMAGE) \
	find ./ -not -path "./third_party/*" -name *.proto -exec clang-format -i {} \;


.PHONY: all install install-debug \
	go-mod-cache draw-deps clean build format \
	test test-all test-build test-cover test-unit test-race \
	test-sim-import-export

mocks:
	@echo "Regenerate mocks..."
	@go generate ./...