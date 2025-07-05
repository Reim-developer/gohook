package webhook

import (
	"gohook/commands/handle"

	"github.com/spf13/cobra"
)

type CobraClosure = func(cmd *cobra.Command, args []string)

type WebhookSendParameters struct {
	tomlConfigPath string
	envUrl         string
	verbose        bool
	dryMode        bool
	threads        int
	loop           int
	delay          int
	explicit       bool
	toJson         bool
}

func handleClosure(commands *WebhookSendParameters) CobraClosure {
	function := func(cmd *cobra.Command, args []string) {
		commands.tomlConfigPath = args[0]

		var arguments = handle.CommandParameters{
			TomlConfigPath: commands.tomlConfigPath,
			Verbose:        commands.verbose,
			DryMode:        commands.dryMode,
			Threads:        commands.threads,
			Loop:           commands.loop,
			Delay:          commands.delay,
			EnvWebhookUrl:  commands.envUrl,
			Explicit:       commands.explicit,
			ToJson:         commands.toJson,
		}

		handle.HandleWebhookSendCommand(&arguments)
	}

	return function
}

func setupCommand(webhookCommand *cobra.Command, commands *WebhookSendParameters) {
	webhookCommand.Flags().BoolVarP(&commands.verbose, "verbose", "v", false, "Enable verbose mode")
	webhookCommand.Flags().BoolVarP(&commands.dryMode, "dry-run", "", false, "Enable dry-run mode")
	webhookCommand.Flags().IntVarP(&commands.threads, "thread", "", 1, "Enable thread run")
	webhookCommand.Flags().IntVarP(&commands.loop, "loop", "l", 1, "Enable loop run")
	webhookCommand.Flags().IntVarP(&commands.delay, "delay", "", 2, "Enable delay")
	webhookCommand.Flags().BoolVarP(&commands.explicit, "explicit", "e", false, "Enable explicit mode")
	webhookCommand.Flags().StringVarP(&commands.envUrl, "use-env-url", "", "", "Enable environment support")
	webhookCommand.Flags().BoolVarP(&commands.toJson, "to-json", "", false, "Export payload to JSON")
}

func WebhookCommand() *cobra.Command {
	var commands = WebhookSendParameters{}
	var function = handleClosure(&commands)

	var webhookCommand = &cobra.Command{
		Use:   "wh-send <TOML Config>",
		Short: "Send content to Discord webhook",
		Args:  cobra.ExactArgs(1),
		Run:   function,
	}
	setupCommand(webhookCommand, &commands)

	return webhookCommand

}
