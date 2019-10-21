package main

import (
	"fmt"
	"github.com/AlexGustafsson/upmon/core"
	"plugin"
	"os"
	"log"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Expected a command to be given")
	}

	command := os.Args[1]

	if command == "generate-keys" {
		generateKeys()
	} else if command == "start" {
		start()
	} else {
		log.Fatal("Unknown command: ", command)
	}
}

func generateKeys() {
	fmt.Println("Private Key: xxxx")
	fmt.Println("Public Key: xxxx")
}

func start() {
	var config = new(core.Config)
	var logLevel string
	var err error = nil
	for i := 1; i < len(os.Args) && err == nil; i++ {
		argument := os.Args[i]
		if argument == "-c" || argument == "--config" {
			i++
			path := os.Args[i]
			config, err = loadConfig(path)
		} else if argument == "-d" || argument == "--debug" {
			i++
			logLevel = os.Args[i]
		}
	}

	if err != nil {
		log.Fatal("Unable to parse arguments, got error:", err)
	}

	if logLevel == "" {
		config.LogLevel = logLevel
	}
}

func loadModule(path string) (*core.Module, error) {
	handle, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	module, err := handle.Lookup("Module")
	if err != nil {
		return nil, err
	}
	*module.(**core.Module) = new(core.Module)

	initialize, err := handle.Lookup("Initialize")
	if err != nil {
		return nil, err
	}
	initialize.(func())()

	return *module.(**core.Module), nil
}
