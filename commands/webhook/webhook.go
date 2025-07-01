package webhook

import (
	"gohook/commands/handle"

	"github.com/spf13/cobra"
)

func WebhookCommand() *cobra.Command {
	var tomlConfigPath string
	var verbose bool
	var dryMode bool
	var threads int
	var loop int
	var delay int

	var webhookCommand = &cobra.Command{
		Use:   "wh-send <TOML Config>",
		Short: "Send content to Discord webhook",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tomlConfigPath = args[0]

			var arguments = handle.CommandParameters{
				TomlConfigPath: tomlConfigPath,
				Verbose:        verbose,
				DryMode:        dryMode,
				Threads:        threads,
				Loop:           loop,
				Delay:          delay,
			}

			handle.HandleCommand(&arguments)
		},
	}

	webhookCommand.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")
	webhookCommand.Flags().BoolVarP(&dryMode, "dry-run", "", false, "Enable dry-run mode")
	webhookCommand.Flags().IntVarP(&threads, "thread", "", 1, "Enable thread run")
	webhookCommand.Flags().IntVarP(&loop, "loop", "l", 1, "Enable loop run")
	webhookCommand.Flags().IntVarP(&delay, "delay", "", 2, "Enable delay")

	return webhookCommand

}
