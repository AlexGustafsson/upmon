# Disable echoing of commands
MAKEFLAGS += --silent

../../build/modules/ping.so: ping.go
	go build --buildmode=plugin -o $@ $<
