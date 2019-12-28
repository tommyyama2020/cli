// +build go1.13

package main

import (
	"os"
	"code.cloudfoundry.org/cli/plugin_parser"
	command_parser "code.cloudfoundry.org/cli/parser"
	"code.cloudfoundry.org/cli/util/panichandler"
)

func main() {
	var exitCode int
	defer panichandler.HandlePanic()
	plugin, commandIsPlugin := plugin_parser.IsPluginCommand(os.Args[0])

	if commandIsPlugin == true {
		exitCode = plugin_parser.RunPlugin(plugin)
	} else {
		exitCode = command_parser.ParseCommandFromArgs(os.Args)
	}
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}
