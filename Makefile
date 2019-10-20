# Disable echoing of commands
MAKEFLAGS += --silent

modules := $(wildcard modules/*)

.PHONY: build clean $(modules)

build: build/upmon $(modules)

build/upmon: main.go core/module.go
	go build -o $@ main.go

$(modules):
	$(MAKE) -C $@

clean:
	rm -rf build
