package main

import (
	"fmt"
	"github.com/AlexGustafsson/upmon/core"
	"github.com/AlexGustafsson/upmon/rpc"
	"google.golang.org/grpc/credentials"
	"github.com/golang/protobuf/ptypes"
	"context"
	"google.golang.org/grpc"
	"github.com/google/uuid"
	"plugin"
	"time"
	"os"
	"crypto/tls"
	"sync"
	"net"
	"crypto/sha1"
	"encoding/base64"
)

func main() {
	if len(os.Args) <= 1 {
		core.LogError("Expected a command to be given")
		printHelp()
		os.Exit(1)
	}

	command := os.Args[1]

	if command == "generate-certificate" {
		generateCertificates()
	} else if command == "start" {
		start()
	} else if command == "help" {
		printHelp()
	} else {
		core.LogError("Unknown command: %v", command)
		printHelp()
		os.Exit(1)
	}
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
			logLevel = "debug"
		}
	}

	if err != nil {
		core.LogError("Unable to parse arguments, got error %v", err)
		os.Exit(1)
	}

	if logLevel != "" {
		config.LogLevel = logLevel
	} else if config.LogLevel == "" {
		config.LogLevel = "error"
	}
	core.SetLogLevel(config.LogLevel)

	self := config.Peers["self"]
	core.LogDebug("Loading certificate from '%s' and key from '%s'", self.Certificate, self.Key)
	certificate, err := tls.LoadX509KeyPair(self.Certificate, self.Key)
	if err != nil {
		core.LogError("Unable to load TLS certificate and key, got error: %v", err)
		os.Exit(1)
	}

	tlsConfig := tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth: tls.RequireAnyClientCert,
		MinVersion: tls.VersionTLS13,
		InsecureSkipVerify: true,
	}

  var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go listen(tlsConfig, self.Hostname, self.Port, waitGroup)

	waitGroup.Add(1)
	go connect(tlsConfig, self.Hostname, self.Port, waitGroup)

	waitGroup.Wait()
}

func connect(tlsConfig tls.Config, hostname string, port int, waitGroup sync.WaitGroup) {
	defer waitGroup.Done()

	connection, err := grpc.Dial(fmt.Sprintf("%s:%d", hostname, port), grpc.WithTransportCredentials(credentials.NewTLS(&tlsConfig)))
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
		Status: rpc.ServicePing_UP,
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

	server := grpc.NewServer(grpc.Creds(credentials.NewTLS(&tlsConfig)))

	// Register services
	rpc.RegisterUpmonServer(server, &upmonServer{})

	if err = server.Serve(listener); err != nil {
		core.LogError("Failed to start server instance: %v", err)
		return
	}
/*
	p, ok := peer.FromContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "no peer found")
	}

	tlsAuth, ok := p.AuthInfo.(credentials.TLSInfo)
	if !ok {
	    return status.Error(codes.Unauthenticated, "unexpected peer transport credentials")
	}

	state := tlsAuth.State
	if len(state.PeerCertificates) == 0 {
		core.LogWarning("Peer from %s sent no certificates, closing connection", connection.RemoteAddr())
		return
	} else if len(state.PeerCertificates) != 1 {
		core.LogWarning("Peer from %s has multiple certificates, closing connection", connection.RemoteAddr())
		return
	}

	certificate := state.PeerCertificates[0]

	shasum := sha1.Sum(certificate.Raw)
	fingerprint := base64.StdEncoding.EncodeToString(shasum[:])
	core.LogDebug("Peer has the fingerprint fingerprint: %v", fingerprint)*/


	core.LogNotice("Listening on port %s:%d", hostname, port)
}

func printHelp() {
	fmt.Println("Upmon ( \xF0\x9D\x9C\xB6 ) - A cloud-native, distributed uptime monitor written in Go");
	fmt.Println("");
	fmt.Println("\x1b[1mVERSION\x1b[0m");
	fmt.Println("Upmon v0.1.0");
	fmt.Println("");
	fmt.Println("\x1b[1mUSAGE\x1b[0m");
	fmt.Println("$ upmon <command> [arguments]");
	fmt.Println("");
	fmt.Println("\x1b[1mCOMMANDS\x1b[0m");
	fmt.Println("start                   Start Upmon");
	fmt.Println("help                    Show this help text");
	fmt.Println("version                 Show current version");
	fmt.Println("generate-certificate    Generate a strong certificate and private key for TLS 1.3");
	fmt.Println("");
	fmt.Println("\x1b[1mARGUMENTS\x1b[0m");
	fmt.Println("-c    --config          Specify the config file");
	fmt.Println("-d    --debug           Run Upmon with debug logging enabled");
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
