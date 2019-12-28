package plugin_parser

import (
	"os"

	"code.cloudfoundry.org/cli/util/configv3"
	"code.cloudfoundry.org/cli/util/ui"
	"code.cloudfoundry.org/cli/plugin/transition"
	"fmt"

	"code.cloudfoundry.org/cli/command/common"
	"code.cloudfoundry.org/cli/command/translatableerror"
)

func runPlugin(plugin configv3.Plugin) int {
	_, commandUI, err := getCFConfigAndCommandUIObjects()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return 1
	}
	defer commandUI.FlushDeferred()
	pluginErr := plugin_transition.RunPlugin(plugin, commandUI)
	if pluginErr != nil {
		handleError(pluginErr, commandUI) //nolint: errcheck
		return 1
	}
	return 0
}

func getCFConfigAndCommandUIObjects() (*configv3.Config, *ui.UI, error) {
	cfConfig, configErr := configv3.LoadConfig(configv3.FlagOverride{
		Verbose: common.Commands.VerboseOrVersion,
	})
	if configErr != nil {
		if _, ok := configErr.(translatableerror.EmptyConfigError); !ok {
			return nil, nil, configErr
		}
	}
	commandUI, err := ui.NewUI(cfConfig)
	return cfConfig, commandUI, err
}

func isPluginCommand(command string, plugins []configv3.Plugin) (configv3.Plugin, bool) {
	for _, plugin := range plugins {
		for _, pluginCommand := range plugin.Commands {
			if command == pluginCommand.Name || command == pluginCommand.Alias {
				return plugin, true
			}
		}
	}

	return configv3.Plugin{}, false
}

func handleError(passedErr error, commandUI UI) error {
	if passedErr == nil {
		return nil
	}

	translatedErr := translatableerror.ConvertToTranslatableError(passedErr)
	commandUI.DisplayError(translatedErr)

	if _, ok := translatedErr.(DisplayUsage); ok {
		return ParseErr
	}

	return ErrFailed
}
