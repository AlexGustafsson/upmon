package main

import (
	"fmt"
	"github.com/AlexGustafsson/upmon/core"
	"plugin"
	"os"
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

	module, err := loadModule("./build/modules/ping.so")
	if err != nil {
		core.LogError("Unable to load module, got error: %v", err)
		os.Exit(1)
	}

	host := new(core.Host)
	host.Name = "TestHost"
	host.IP = "localhost"
	module.CheckStatus(host)
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