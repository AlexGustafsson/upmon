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

	core.LogDebug("Initializing module '%s'", Module.Name)
}

func checkStatus(host *core.Host) (int, error) {
	core.LogDebug("Checking status of host '%s' by pinging '%s'", host.Name, host.IP)

	_, err := exec.Command("ping", host.IP, "-c 1", "-t 1").Output()
	if err != nil {
		core.LogDebug("Status of host '%s' was down", host.Name)
		return core.StatusDown, nil
	}

	core.LogDebug("Status of host '%s' was up", host.Name)
	return core.StatusUp, nil
}
