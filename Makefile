# Disable echoing of commands
MAKEFLAGS += --silent

# Add build-time variables
PREFIX := $(shell go list ./internal/version)
VERSION := v0.1.0
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
GO_VERSION := $(shell go version)
COMPILE_TIME := $(shell LC_ALL=en_US date)

BUILD_VARIABLES := -X "$(PREFIX).Version=$(VERSION)" -X "$(PREFIX).Commit=$(COMMIT)" -X "$(PREFIX).GoVersion=$(GO_VERSION)" -X "$(PREFIX).CompileTime=$(COMPILE_TIME)"
BUILD_FLAGS := -ldflags '$(BUILD_VARIABLES)'

server_source := $(shell find . -type f -name '*.go')

# Force macOS to use clang
# https://gcc.gnu.org/bugzilla/show_bug.cgi?id=93082
# https://bugs.llvm.org/show_bug.cgi?id=44406
# https://openradar.appspot.com/radar?id=4952611266494464
ifeq ($(shell uname),Darwin)
	CC=clang
endif

.PHONY: help build generate format lint test clean install-tools

# Produce a short description of available make commands
help:
	pcregrep -Mo '^(#.*\n)+^[^# ]+:' Makefile | sed "s/^\([^# ]\+\):/> \1/g" | sed "s/^#\s\+\(.\+\)/\1/g" | GREP_COLORS='ms=1;34' grep -E --color=always '^>.*|$$' | GREP_COLORS='ms=1;37' grep -E --color=always '^[^>].*|$$'

# Build for the native platform
build: build/upmon

# Generate clients etc.
# Requires openapi-generator
# brew install openapi-generator
generate: api.yml
	openapi-generator generate --package-name api --generator-name go --input-spec api.yml --output api
# The code is not automatically formatted after generation
	gofmt -l -s -w api

# Format Go code
format: $(server_source) Makefile
	gofmt -l -s -w .

# Lint Go code
lint: $(server_source) Makefile
	golint .

# Check the Go code for issues
check: vet gosec test

# Vet Go code
vet: $(server_source) Makefile
	go vet ./...

# Check the code for security issues
# Requires gosec
# https://github.com/securego/gosec
gosec: $(server_source) Makefile
	gosec -exclude-dir=api ./...

# Test Go code
test: $(server_source) Makefile
	go test -v ./...

# Build for the native platform
build/upmon: $(server_source) Makefile
	go generate ./...
	go build $(BUILD_FLAGS) -o $@ cmd/upmon.go

install-tools:
	cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

# Clean all dynamically created files
clean:
	rm -rf ./build &> /dev/null || true
