# Disable echoing of commands
MAKEFLAGS += --silent

modules := $(wildcard modules/*)
source := $(shell find ./ -type f -name '*.go')

.PHONY: build clean format lint $(modules)

build: build/upmon $(modules)

build/upmon: main.go core/module.go
	go build -o $@ main.go

$(modules):
	$(MAKE) -C $@

format: $(source)
	gofmt -l -s -w .

lint: $(source)
	golint $<

clean:
	rm -rf build
