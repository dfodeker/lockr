package command

import (
	"encoding/json"
	"flag"
	"fmt"
	config "lockr/config"
	"os"
	"path/filepath"
)

var listUsage = `List all currently avaible enviroments. Usage: brief list Options: `
var listFunc = func(cmd *Command, args []string) {
	configPath := filepath.Join(LockrDir, ConfigFile)
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: unable to load config file: %v\n", err)
		os.Exit(1) // Exit if the config file can't be read
	}

	var config config.Config
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Error: unable to parse config file: %v\n", err)
		os.Exit(1) // Exit if JSON parsing fails
	}

	// If needed, interact with cmd (e.g., check flags)
	if cmd.flags.Lookup("short").Value.String() == "true" {
		fmt.Println("Short mode enabled")
	}
	for _, env := range config.Enviroments {
		fmt.Printf("%s\n", env)

	}
	// Process the config
	//fmt.Printf("Loaded config: %+v\n", config)
}

func NewListCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("List", flag.ExitOnError),
		Execute: listFunc,
	}
	cmd.flags.BoolVar(&short, "short", false, "")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, versionUsage)
	}

	return cmd
}
