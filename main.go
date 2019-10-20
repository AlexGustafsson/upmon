package main

import (
	"fmt"
	"github.com/AlexGustafsson/upmon/core"
	"log"
	"plugin"
)

func main() {
	module, err := loadModule("./build/modules/ping.so")
	if err != nil {
		log.Fatal(err)
	}

	var host = new(core.Host)
	host.Name = "Test"
	host.Description = "Test host"
	host.IP = "localhost"

	status, err := module.CheckStatus(host)
	if err != nil {
		log.Fatal(err)
	}

	if status == core.StatusUp {
		fmt.Println("The host is up")
	} else {
		fmt.Println("The host is down")
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
