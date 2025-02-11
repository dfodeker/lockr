package command

import (
	"flag"
)

type Command struct {
	flags   *flag.FlagSet
	Execute func(cmd *Command, args []string)
}

func (c *Command) Init(args []string) error {
	return c.flags.Parse(args)
}

func (c *Command) Called() bool {
	return c.flags.Parsed()
}

func (c *Command) Run() {
	c.Execute(c, flag.Args())
}

const ConfigFile string = "config.json"

const GitIgnore string = ".gitignore"

const LockrDir string = ".lockr"
