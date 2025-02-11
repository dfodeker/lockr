package command

import (
	"flag"
	"fmt"
	"os"
)

var versionUsage = `Print the app version and build info for the current context. Usage: lockr version [options] Options: --short If true, print just the version number. Default false. `
var (
	build   = "???"
	version = "???"
	short   = false
)

var versionFunc = func(cmd *Command, args []string) {
	if short {
		fmt.Printf("brief version: v%s\n", version)
	} else {
		fmt.Printf("brief version: v%s, build: %s \n", version, build)
	}
	os.Exit(0)
}

func VersionCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("version", flag.ExitOnError),
		Execute: versionFunc,
	}

	cmd.flags.BoolVar(&short, "short", false, "")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, versionUsage)
	}

	return cmd
}
