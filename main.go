package main

import (
	"flag"
	"fmt"
	"os"
)

var usage = `Usage lockr command [options] a simple tool to allow you manage the differnt env's within a project without a need for multiple .env files in a project`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}
	usageAndExit("")
}

func usageAndExit(msg string) {
	if len(msg) > 1 {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	os.Exit(0)
}
