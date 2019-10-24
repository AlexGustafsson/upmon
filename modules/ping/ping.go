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
	Module.CheckService = checkService

	core.LogDebug("Initializing module '%s'", Module.Name)
}

func checkService(service *core.ServiceConfig) (*core.ServiceInfo, error) {
	core.LogDebug("Checking status of service '%s' by pinging '%s'", service.Name, service.Hostname)

	serviceInfo := &core.ServiceInfo{}

	_, err := exec.Command("ping", service.Hostname, "-c 1", "-t 1").Output()
	if err != nil {
		core.LogDebug("Status of service '%s' was down", service.Name)
		serviceInfo.Status = core.StatusDown
		return serviceInfo, nil
	}

	core.LogDebug("Status of service '%s' was up", service.Name)
	serviceInfo.Status = core.StatusUp
	return serviceInfo, nil
}
