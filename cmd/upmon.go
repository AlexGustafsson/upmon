package main

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/AlexGustafsson/upmon/internal/commands"
	"github.com/AlexGustafsson/upmon/internal/version"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var appHelpTemplate = `Usage: {{.Name}} [global options] command [command options] [arguments]

{{.Usage}}

Version: {{.Version}}

Options:
  {{range .Flags}}{{.}}
  {{end}}
Commands:
  {{range .Commands}}{{.Name}}{{ "\t" }}{{.Usage}}
  {{end}}
Run '{{.Name}} help command' for more information on a command.
`

var commandHelpTemplate = `Usage: upmon {{.Name}} [options] {{if .ArgsUsage}}{{.ArgsUsage}}{{end}}

{{.Usage}}{{if .Description}}

Description:
   {{.Description}}{{end}}{{if .Flags}}

Options:{{range .Flags}}
   {{.}}{{end}}{{end}}
`

func setDebugOutputLevel() {
	for _, flag := range os.Args {
		if flag == "-v" || flag == "--verbose" {
			log.SetLevel(log.DebugLevel)
		}
	}
}

func commandNotFound(context *cli.Context, command string) {
	log.Errorf(
		"%s: '%s' is not a %s command. See '%s help'.",
		context.App.Name,
		command,
		context.App.Name,
		os.Args[0],
	)
	os.Exit(1)
}

func main() {
	setDebugOutputLevel()

	cli.AppHelpTemplate = appHelpTemplate
	cli.CommandHelpTemplate = commandHelpTemplate

	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Usage = "A cloud-native, distributed uptime monitor"
	app.Version = version.FullVersion()
	app.CommandNotFound = commandNotFound
	app.EnableBashCompletion = true
	app.Commands = commands.Commands
	app.HideVersion = true
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "Enable verbose logging",
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
