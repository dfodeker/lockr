package command

import (
	"encoding/json"
	"flag"
	"fmt"
	config "lockr/config"
	"os"
	"path/filepath"
)

var nicely bool

var envUsage = `The 'env' command is used to manage environments. Environments are isolated workspaces that allow you to switch between different configurations or contexts.

Usage:
  lockr env <subcommand> [flags]

Subcommands:
  create    Create a new environment.
  switch    Switch to a specified environment.

Examples:
  lockr env create myenv      Create a new environment named 'myenv'.
  lockr env switch myenv      Switch to the environment named 'myenv'.`

var envFunc = func(cmd *Command, args []string) {
	configPath := filepath.Join(LockrDir, ConfigFile)
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to load config file: %v", err)
		os.Exit(1)
	}
	var config config.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse config file: %v", err)
		os.Exit(1)
	}

	if nicely {
		fmt.Printf("Active Env: %s\n", config.ActiveEnv)
	} else {
		fmt.Printf("Active enviroment  is '%s'\n", config.ActiveEnv)
	}
	os.Exit(0)
}

func EnviromentCommand() *Command {

	cmd := &Command{
		flags:   flag.NewFlagSet("env", flag.ExitOnError),
		Execute: envFunc,
	}

	cmd.flags.BoolVar(&nicely, "nicely", false, "")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, envUsage)
	}

	return cmd
}
