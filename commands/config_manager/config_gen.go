package config_manager

import (
	"gohook/commands/handle"

	"github.com/spf13/cobra"
)

type CobraClosure = func(cmd *cobra.Command, args []string)

func handleClosure() CobraClosure {
	function := func(cmd *cobra.Command, args []string) {
		var tomlConfigName = args[0]

		var arguments = handle.GenCfgCommand{
			TomlConfigName: tomlConfigName,
		}
		handle.HandleGenCfgCommand(&arguments)
	}

	return function
}

func GenTomlConfigCommand() *cobra.Command {
	var function = handleClosure()

	var genTomlCommand = &cobra.Command{
		Use:   "gen-cfg <TOML_CONFIG_PATH>",
		Short: "Generate TOML configuration for webhook with the given name.",
		Args:  cobra.ExactArgs(1),

		Run: function,
	}

	return genTomlCommand
}
