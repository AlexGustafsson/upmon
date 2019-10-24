# Disable echoing of commands
MAKEFLAGS += --silent

../../build/modules/irc.so: irc.go
	go build --buildmode=plugin -o $@ $<
