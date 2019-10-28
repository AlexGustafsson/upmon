package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"github.com/AlexGustafsson/upmon/core"
	"net"
	"strings"
	"time"
)

// Module is the exported ping module
var Module *core.Module

// Initialize the module
func Initialize() {
	Module.Name = "irc"
	Module.Description = "Get the current status of an IRC server"
	Module.Version = core.ModuleVersion1
	Module.CheckService = checkService

	core.LogDebug("Initializing module '%s'", Module.Name)
}

func checkService(service *core.ServiceConfig) (*core.ServiceInfo, error) {
	core.LogDebug("Checking status of service '%s' by connecting via IRC '%s'", service.Name, service.Hostname)

	serviceInfo := &core.ServiceInfo{}

	connection, err := net.Dial("tcp", fmt.Sprintf("%v:%v", service.Hostname, service.Port))
	if err != nil {
		core.LogError("Unable to connect to IRC server (service '%v'), got error: ", service.Name, err)
		return serviceInfo, err
	}
	defer connection.Close()

	tlsConfig := tls.Config{
		InsecureSkipVerify: true,
		ServerName:         service.Hostname,
	}

	core.LogDebug("Connected to IRC server, trying to upgrade to TLS")
	tlsConnection := tls.Client(connection, &tlsConfig)
	err = tlsConnection.Handshake()
	if err != nil {
		core.LogDebug("Unable to perform TLS handshake with IRC server '%s', assuming HTTP", service.Name)
	} else {
		core.LogDebug("Upgraded connection to TLS")
		// Overwrite the connection variable to make TLS transparent
		connection = net.Conn(tlsConnection)
	}

	err = sendPing(connection, service)
	if err != nil {
		core.LogError("Unable to send a ping to IRC server '%v', got error: %v", err)
		return serviceInfo, err
	}

	core.LogDebug("Successfully sent PING to service")

	err = connection.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		core.LogError("Unable to set deadline for connection, got error: %v", err)
		return serviceInfo, err
	}

	reader := bufio.NewReader(connection)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		err := scanner.Err()
		if err != nil {
			core.LogDebug("Unable to read line from IRC connection, got error: %v", err)
			return serviceInfo, err
		}

		core.LogDebug("Got response from IRC server: <%v>", line)

		// The service is considered up if it responds to a PING message
		if strings.Contains(line, "PONG") {
			serviceInfo.Status = core.StatusUp
			break
		}
	}

	return serviceInfo, nil
}

func sendPing(connection net.Conn, service *core.ServiceConfig) error {
	message := fmt.Sprintf("PING %v\r\n", service.Hostname)
	_, err := connection.Write([]byte(message))
	return err
}
