BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')
BUILD_DIR ?= $(CURDIR)/build
DOCKER := $(shell which docker)
LEDGER_ENABLED ?= true
DB_BACKEND ?= goleveldb
CGO_ENABLED=1

# don't override user values
ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match 2>/dev/null)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
TM_VERSION := $(shell go list -m github.com/cometbft/cometbft | sed 's:.* ::')

###############################################################################
###                            Build Tags/Flags                             ###
###############################################################################
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

ifeq ($(DB_BACKEND), goleveldb)
  build_tags += goleveldb
endif

ifeq ($(DB_BACKEND), cleveldb)
  build_tags += gcc
  build_tags += cleveldb
endif

ifeq ($(DB_BACKEND), boltdb)
  build_tags += boltdb
endif

ifeq ($(DB_BACKEND), rocksdb)
  build_tags += rocksdb
endif

ifeq ($(DB_BACKEND), badgerdb)
  build_tags += badgerdb
endif

build_tags += $(TAG)
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

# process linker flags
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=re \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=red \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ifeq ($(DB_BACKEND), goleveldb)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=goleveldb
endif

ifeq ($(DB_BACKEND), cleveldb)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif

ifeq ($(DB_BACKEND), boltdb)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb
endif

ifeq ($(DB_BACKEND), rocksdb)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif

ifeq ($(DB_BACKEND), badgerdb)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb
endif

ldflags += -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)

ifeq ($(NO_STRIP),false)
  ldflags += -w -s
endif

ldflags += $(LD_FLAGS)
ldflags := $(strip $(ldflags))

# set build flags

BUILD_FLAGS := -tags '$(build_tags)' -ldflags '$(ldflags)'

ifeq ($(NO_STRIP),false)
  BUILD_FLAGS += -trimpath
endif

###############################################################################
###                               Go Version                                ###
###############################################################################

GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
MIN_GO_MAJOR_VERSION = 1
MIN_GO_MINOR_VERSION = 22
GO_VERSION_ERROR = Golang version $(GO_MAJOR_VERSION).$(GO_MINOR_VERSION) is not supported, \
please update to at least $(MIN_GO_MAJOR_VERSION).$(MIN_GO_MINOR_VERSION)

go-version:
	@echo "Verifying go version..."
	@if [ $(GO_MAJOR_VERSION) -gt $(MIN_GO_MAJOR_VERSION) ]; then \
		exit 0; \
	elif [ $(GO_MAJOR_VERSION) -lt $(MIN_GO_MAJOR_VERSION) ]; then \
		echo $(GO_VERSION_ERROR); \
		exit 1; \
	elif [ $(GO_MINOR_VERSION) -lt $(MIN_GO_MINOR_VERSION) ]; then \
		echo $(GO_VERSION_ERROR); \
		exit 1; \
	fi

.PHONY: go-version

###############################################################################
###                             Build / Install                             ###
###############################################################################

all: install

build: go.sum go-version
	@mkdir -p $(BUILD_DIR)
	@if [ -z "$(GOARCH)" ]; then \
		go build -o $(BUILD_DIR)/red $(CURDIR)/cmd/red; \
	else \
		go build -o $(BUILD_DIR)/red-$(GOARCH) $(CURDIR)/cmd/red; \
	fi

install:
	@echo "--> ensure dependencies have not been modified"
	@go mod verify
	@echo "--> installing red"
	@go install $(BUILD_FLAGS) -mod=readonly $(CURDIR)/cmd/red

.PHONY: build install

###############################################################################
###                          Tools & Dependencies                           ###
###############################################################################
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

###############################################################################
###                                 Tests                                   ###
###############################################################################

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

###############################################################################
###                                Linting                                  ###
###############################################################################

golangci_version=v1.56.0

lint:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@golangci-lint run ./... --timeout 15m

lint-fix:
	@echo "--> Running linter and fixing issues"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@golangci-lint run ./... --fix --timeout 15m

.PHONY: lint lint-fix

###############################################################################
###                                Protobuf                                 ###
###############################################################################
protoVer=0.14.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating protobuf files..."
	@$(protoImage) sh ./scripts/protocgen.sh
	@go mod tidy

proto-format:
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

proto-lint:
	@$(protoImage) buf lint proto/ --error-format=json

.PHONY: proto-all proto-gen proto-format proto-lint

###############################################################################
###                                LOCALNET                                 ###
###############################################################################
localnet-single-node:
	./scripts/localnet_single_node.sh