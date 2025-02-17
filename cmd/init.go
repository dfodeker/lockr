package command

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var usage = `lockr => Initialize new Lockr env with Lockr init`
var initFunc = func(cmd *Command, args []string) {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal("Failed to get file path")
	}
	successMsg := fmt.Sprintf("Initialized Lockr 🚀 in %v", cwd)
	if _, err := os.Stat(LockrDir); err == nil {
		// for redabliity file = dir,

		//remove file and recreate it
		//using the os.MkdirAll will not let the user know if file exists
		//and will do nth if it does
		err := os.RemoveAll(LockrDir)
		if err != nil {
			log.Fatal("Failed to reinitialize lockr in ", cwd)
		}
		fmt.Printf("Reinitialized existing Lockr vault in %v\n", cwd)
		reinitMsg := fmt.Sprintf("Reinitialized existing Lockr vault in %v\n", cwd)
		successMsg = reinitMsg
	}
	CreateFile(LockrDir)
	createConfig()
	appendGitIgnore()
	fmt.Println(successMsg)
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
