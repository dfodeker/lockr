package main

import (
	"flag"
	"fmt"
	command "lockr/cmd"
	"os"
)

// const lockrDir string = ".lockr"

// Name of config file for Lockr
const ConfigFile string = "config.json"

const GitIgnore string = ".gitignore"

const LockrDir string = ".lockr"

var usage = `Usage lockr command [options] a simple tool to allow you manage the differnt env's within a project without a need for multiple .env files in a project`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}
	var cmd *command.Command
	switch os.Args[1] {
	case "version":
		cmd = command.VersionCommand()
	case "list":
		cmd = command.NewListCommand()
	default:
		usageAndExit(fmt.Sprintf("Lockr: '%s' is not a lockr command.\n", os.Args[1]))
	}
	cmd.Init(os.Args[2:])
	cmd.Run()
	usageAndExit("")
}

func usageAndExit(msg string) {
	if len(msg) > 1 {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	os.Exit(0)
}
