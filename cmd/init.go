package command

import (
	"flag"
	"fmt"
	"os"
)

var usage = `lockr => Initialize new Lockr env with Lockr init`
var initFunc = func(cmd *Command, args []string) {
	fmt.Print("hi there thank you for initializing\n")
}

func InitCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("version", flag.ExitOnError),
		Execute: initFunc,
	}
	//cmd.flags.BoolVar(&short, "short", false, "")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, versionUsage)
	}
	return cmd
}
