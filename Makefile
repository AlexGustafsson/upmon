# Disable echoing of commands
MAKEFLAGS += --silent

CC=clang
CXX=clang

UPMON_VERSION=0.1.0

BUILD_VARIABLES :=-X "main.upmonVersion=$(UPMON_VERSION)" -X "main.goVersion=$(shell go version)" -X "main.compileTime=$(shell LC_ALL=en_US date)"
BUILD_FLAGS :=-ldflags '$(BUILD_VARIABLES)'

modules := $(wildcard modules/*)
source := $(shell find ./ -type f -name '*.go')
coreSource := $(shell find core -type f -name '*.go')
rpcSource := $(shell find rpc -type f -name '*.go')
rpcDefinitions := $(shell find rpc -type f -name '*.proto')
rpcGenerated :=$(rpcDefinitions:.proto=.pb.go)
mainSource := $(shell find ./ -depth 1 -type f -name '*.go')

.PHONY: build clean format lint $(modules)

build: build/upmon $(modules)

build/upmon: $(rpcGenerated) $(mainSource) $(coreSource) $(rpcSource)
	go build $(BUILD_FLAGS) -o $@ $(mainSource)

$(rpcGenerated): rpc/%.pb.go: rpc/%.proto
	protoc --go_out=plugins=grpc:. $<

$(modules): $(coreSource)
	$(MAKE) -C $@

format: $(source)
	gofmt -l -s -w .

lint: $(source)
	golint $<

clean:
	rm -rf build
	rm rpc/*.pb.go &> /dev/null || exit 0
