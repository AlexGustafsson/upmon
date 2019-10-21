# Disable echoing of commands
MAKEFLAGS += --silent

modules := $(wildcard modules/*)
source := $(shell find ./ -type f -name '*.go')
coreSource := $(shell find core -type f -name '*.go')
mainSource := $(shell find ./ -depth 1 -type f -name '*.go')

.PHONY: build clean format lint $(modules)

build: build/upmon $(modules)

build/upmon: $(mainSource) $(coreSource)
	go build -o $@ $(mainSource)

$(modules):
	$(MAKE) -C $@

format: $(source)
	gofmt -l -s -w .

lint: $(source)
	golint $<

clean:
	rm -rf build
