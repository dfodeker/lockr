package main

import (
	"flag"
	"fmt"
	command "lockr/cmd"
	"os"
)

// const lockrDir string = ".lockr"

// Name of config file for Lockr

var usage = `Usage lockr command [options] a simple tool to allow you manage the differnt env's within a project without a need for multiple .env files in a project`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}
	var cmd *command.Command

	if len(os.Args) < 2 {
		usageAndExit(usage)
	}

	switch os.Args[1] {
	case "init":
		cmd = command.InitCommand()
	case "version":
		cmd = command.VersionCommand()
	case "list":
		cmd = command.NewListCommand()
	case "env":
		cmd = command.EnviromentCommand()
	default:
		usageAndExit(fmt.Sprintf("Lockr: '%s' is not a lockr command.\n", os.Args[1]))
	}

	cmd.Init(os.Args[2:])
	cmd.Run()

}

func usageAndExit(msg string) {
	if len(msg) > 1 {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	os.Exit(0)
}
