package commands

import (
	"gohook/commands/config_manager"
	"gohook/commands/webhook"
	"gohook/core"
	"gohook/utils"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "gohook",
	Short: "Webhook sender with Golang",
	Long:  "GoHook is a Command Line Interface for Discord webhook sender",
}

func Execute() {
	rootCommand.AddCommand(webhook.WebhookCommand())
	rootCommand.AddCommand(config_manager.GenTomlConfigCommand())

	if err := rootCommand.Execute(); err != nil {
		utils.CriticalShow("Details error: %s", err)
		os.Exit(core.CommandRunFailed)
	}
}
