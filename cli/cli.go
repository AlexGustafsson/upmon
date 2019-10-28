package cli

import (
	"fmt"
	"github.com/AlexGustafsson/upmon/core"
	"os"
)

type standaloneCommandHandler func()
type commandHandler func(*core.Config)

type argument struct {
	Short       string
	Long        string
	Description string
}

type standaloneCommand struct {
	Name        string
	Description string
	Handler     standaloneCommandHandler
}

type command struct {
	Name        string
	Description string
	Handler     commandHandler
}

var standaloneCommands []standaloneCommand
var commands []command
var arguments []argument

var appVersion string
var goVersion string
var compileTime string

// ParseArguments parses CLI arguments
func ParseArguments() (*core.Config, error) {
	// Parse config and arguments
	config := new(core.Config)
	var logLevel string
	var err error = nil
	for i := 1; i < len(os.Args) && err == nil; i++ {
		argument := os.Args[i]
		if argument == "-c" || argument == "--config" {
			i++
			path := os.Args[i]
			config, err = core.LoadConfig(path)
		} else if argument == "-d" || argument == "--debug" {
			logLevel = "debug"
		}
	}

	if err != nil {
		core.LogError("Unable to parse arguments, got error %v", err)
		return nil, err
	}

	if logLevel != "" {
		config.LogLevel = logLevel
	} else if config.LogLevel == "" {
		config.LogLevel = "error"
	}
	core.SetLogLevel(config.LogLevel)

	return config, nil
}

// Run parses arguments and runs the specified command
func Run(config *core.Config) {
	RegisterStandaloneCommand(
		"help",
		printHelp,
		"Print this help text",
	)

	RegisterStandaloneCommand(
		"version",
		printVersion,
		"Print the version of Upmon",
	)

	RegisterArgument(
		"c",
		"config",
		"Specify the config file",
	)

	RegisterArgument(
		"d",
		"debug",
		"Run Upmon with debug logging enabled",
	)

	if len(os.Args) <= 1 {
		core.LogError("Expected a command to be given")
		printHelp()
		os.Exit(1)
	}

	commandName := os.Args[1]

	// Try standalone commands
	for _, command := range standaloneCommands {
		if commandName == command.Name {
			command.Handler()
			return
		}
	}

	// Try regular commands
	for _, command := range commands {
		if commandName == command.Name {
			command.Handler(config)
			return
		}
	}

	core.LogError("Unknown command: %v", commandName)
	printHelp()
	os.Exit(1)
}

// RegisterStandaloneCommand registers a command without any surrounding context
func RegisterStandaloneCommand(name string, handler standaloneCommandHandler, description string) {
	command := standaloneCommand{
		Name:        name,
		Description: description,
		Handler:     handler,
	}

	standaloneCommands = append(standaloneCommands, command)
}

// RegisterCommand registers a command which needs surrounding context
func RegisterCommand(name string, handler commandHandler, description string) {
	command := command{
		Name:        name,
		Description: description,
		Handler:     handler,
	}

	commands = append(commands, command)
}

// RegisterArgument registers an argument
func RegisterArgument(short string, long string, description string) {
	argument := argument{
		Short:       short,
		Long:        long,
		Description: description,
	}

	arguments = append(arguments, argument)
}

// SetVersion sets the version of the application
func SetVersion(version string) {
	appVersion = version
}

// SetGoVersion sets the Go version used to compile the application
func SetGoVersion(version string) {
	goVersion = version
}

// SetCompileTime sets the time at which the application was compiled
func SetCompileTime(time string) {
	compileTime = time
}

func printVersion() {
	fmt.Println("Upmon ( \xF0\x9D\x9C\xB6 ) - A cloud-native, distributed uptime monitor written in Go")
	fmt.Println("")
	fmt.Println("\x1b[1mVERSION\x1b[0m")
	fmt.Println(fmt.Sprintf("Upmon v%v", appVersion))
	fmt.Println(fmt.Sprintf("Compiled %v with %v", compileTime, goVersion))
	fmt.Println("")
	fmt.Println("\x1b[1mOPEN SOURCE\x1b[0m")
	fmt.Println("The source code is hosted on https://github.com/AlexGustafsson/upmon")
}

func printHelp() {
	fmt.Println("Upmon ( \xF0\x9D\x9C\xB6 ) - A cloud-native, distributed uptime monitor written in Go")
	fmt.Println("")
	fmt.Println("\x1b[1mVERSION\x1b[0m")
	fmt.Println(fmt.Sprintf("Upmon v%v", appVersion))
	fmt.Println("")
	fmt.Println("\x1b[1mUSAGE\x1b[0m")
	fmt.Println("$ upmon <command> [arguments]")
	fmt.Println("")
	fmt.Println("\x1b[1mCOMMANDS\x1b[0m")

	padLength := getLongestCommandLength()

	for _, command := range commands {
		fmt.Printf("%-*v    %v\n", padLength, command.Name, command.Description)
	}
	for _, command := range standaloneCommands {
		fmt.Printf("%-*v    %v\n", padLength, command.Name, command.Description)
	}

	fmt.Println("")
	fmt.Println("\x1b[1mARGUMENTS\x1b[0m")

	for _, argument := range arguments {
		fmt.Printf("%-*v    %v\n", padLength, fmt.Sprintf("-%v    --%v", argument.Short, argument.Long), argument.Description)
	}
}

func getLongestCommandLength() int {
	length := 0
	for _, command := range commands {
		if len(command.Name) > length {
			length = len(command.Name)
		}
	}

	for _, command := range standaloneCommands {
		if len(command.Name) > length {
			length = len(command.Name)
		}
	}

	return length
}
