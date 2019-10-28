package main

import (
	"context"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/AlexGustafsson/upmon/cli"
	"github.com/AlexGustafsson/upmon/core"
	"github.com/AlexGustafsson/upmon/rpc"
	"github.com/AlexGustafsson/upmon/transport"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"net"
	"os"
	"plugin"
	"sync"
	"time"
)

func main() {
	config, err := cli.ParseArguments()
	if err != nil {
		os.Exit(1)
	}

	cli.SetVersion(upmonVersion)
	cli.SetGoVersion(goVersion)
	cli.SetCompileTime(compileTime)

	cli.RegisterStandaloneCommand(
		"generate-certificate",
		generateCertificates,
		"Generate a strong certificate and private key for TLS 1.3",
	)

	cli.RegisterCommand(
		"start",
		start,
		"Start Upmon",
	)

	cli.RegisterCommand(
		"check",
		check,
		"Run checks",
	)

	var allowedFingerprints []string
	for _, peer := range config.Peers {
		core.LogDebug(peer.Fingerprint)
		allowedFingerprints = append(allowedFingerprints, peer.Fingerprint)
	}

	transport.SetAllowedFingerprints(allowedFingerprints)

	cli.Run(config)
}

func generateCertificates() {
	privateKey, _, certificateBytes, err := createSelfSignedCertificate("localhost")
	if err != nil {
		core.LogError("Unable to create self-signed server certificate, got error: %v", err)
		os.Exit(1)
	}

	err = writeCertificate(certificateBytes, "./server.crt")
	if err != nil {
		core.LogError("%v", err)
		os.Exit(1)
	}
	fmt.Println("Stored certificate in: server.crt")

	err = writeKey(privateKey, "./server.pem")
	if err != nil {
		core.LogError("%v", err)
		os.Exit(1)
	}
	fmt.Println("Stored private key in: server.pem")

	shasum := sha1.Sum(certificateBytes)
	fingerprint := base64.StdEncoding.EncodeToString(shasum[:])
	fmt.Println("Fingerprint is:", fingerprint)
}

func start(config *core.Config) {
	self := config.GetPeerByName("self")
	core.LogDebug("Loading certificate from '%s' and key from '%s'", self.Certificate, self.Key)
	certificate, err := tls.LoadX509KeyPair(self.Certificate, self.Key)
	if err != nil {
		core.LogError("Unable to load TLS certificate and key, got error: %v", err)
		os.Exit(1)
	}

	tlsConfig := tls.Config{
		Certificates:       []tls.Certificate{certificate},
		ClientAuth:         tls.RequireAnyClientCert,
		MinVersion:         tls.VersionTLS13,
		InsecureSkipVerify: true,
	}

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go listen(tlsConfig, self.Hostname, self.Port, waitGroup)

	waitGroup.Add(1)
	go connect(tlsConfig, self.Hostname, self.Port, waitGroup)

	waitGroup.Wait()
}

func check(config *core.Config) {
	modules := make(map[string]*core.Module)
	for _, moduleName := range config.GetUsedModules() {
		module, err := loadModule(fmt.Sprintf("./build/modules/%v.so", moduleName))
		modules[moduleName] = module
		if err != nil {
			core.LogError("Unable to load module '%v', got error: %v", moduleName, err)
			os.Exit(1)
		}
	}

	if len(config.Services) == 0 {
		core.LogWarning("No services configured, nothing to check")
		return
	}

	var result core.Result
	result.Services = make(map[string]*core.ServiceResult)

	for _, service := range config.Services {
		core.LogDebug("Checking status for service '%v'", service.Name)
		serviceResult := core.ServiceResult{
			Name:   service.Name,
			Status: core.StatusUnknown,
		}
		result.Services[service.Name] = &serviceResult

		for _, moduleName := range service.Checks {
			core.LogDebug("Using module '%v'", moduleName)
			module := modules[moduleName]
			serviceResult.Timestamp = time.Now().Unix()
			serviceInfo, err := module.CheckService(&service)
			if err != nil {
				core.LogError("Unable to check status of service '%v' using module '%v', got error: %v", service.Name, moduleName, err)
				os.Exit(1)
			}

			core.LogDebug("The IRC service had the status %v", serviceInfo.Status)

			if serviceInfo.Status != core.StatusUnknown {
				core.LogDebug("Got a result from module '%v', assuming accurate", moduleName)
				serviceResult.Status = serviceInfo.Status
				break
			}
		}
	}

	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		core.LogError("Unable to marshal result to JSON, got error: %v", err)
		os.Exit(1)
	}

	os.Stdout.Write(jsonBytes)
	fmt.Println()
}

func connect(tlsConfig tls.Config, hostname string, port int, waitGroup sync.WaitGroup) {
	defer waitGroup.Done()

	connection, err := grpc.Dial(fmt.Sprintf("%s:%d", hostname, port), grpc.WithTransportCredentials(transport.WithConfig(&tlsConfig)))
	if err != nil {
		core.LogError("Unable to connect to peer %v:%v, got error: %v", hostname, port, err)
		return
	}
	defer connection.Close()

	client := rpc.NewUpmonClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Generate UUIDv4
	serviceID, err := uuid.NewRandom()
	if err != nil {
		core.LogError("Unable to generate UUIDv4")
		return
	}

	// Format UUIDv4
	formattedServiceID := serviceID.URN()

	timestamp := ptypes.TimestampNow()

	result, err := client.SendServicePing(ctx, &rpc.ServicePing{
		ServiceId: formattedServiceID,
		Status:    rpc.ServicePing_UP,
		Timestamp: timestamp,
	})
	if err != nil {
		core.LogError("Unable to send service ping")
		return
	}

	core.LogDebug("Got result from server: %v", result)
}

func listen(tlsConfig tls.Config, hostname string, port int, waitGroup sync.WaitGroup) {
	defer waitGroup.Done()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		core.LogError("Unable to create listener instance, got error: %v", err)
		return
	}

	server := grpc.NewServer(grpc.Creds(transport.WithConfig(&tlsConfig)))

	// Register services
	rpc.RegisterUpmonServer(server, &upmonServer{})

	err = server.Serve(listener)
	if err != nil {
		core.LogError("Failed to start server instance: %v", err)
		return
	}

	core.LogNotice("Listening on port %s:%d", hostname, port)
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
