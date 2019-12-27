// +build go1.13

package main

import (
	"os"

	command_parser "code.cloudfoundry.org/cli/parser"
	"code.cloudfoundry.org/cli/util/panichandler"
)

func main() {
	defer panichandler.HandlePanic()
	exitStatus := command_parser.CommandParser(os.Args)
	if exitStatus != 0 {
		os.Exit(exitStatus)
	}
}
