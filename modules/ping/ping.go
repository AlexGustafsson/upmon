package main

import (
	"github.com/AlexGustafsson/upmon/core"
	"os/exec"
)

// Module is the exported ping module
var Module *core.Module

// Initialize the module
func Initialize() {
	Module.Name = "ping"
	Module.Description = "Get the current status of a host by pinging them"
	Module.Version = core.ModuleVersion1
	Module.CheckStatus = checkStatus
}

func checkStatus(host *core.Host) (int, error) {
	_, err := exec.Command("ping", host.IP, "-c 1", "-t 1").Output()
	if err != nil {
		return core.StatusDown, nil
	}

	return core.StatusUp, nil
}
