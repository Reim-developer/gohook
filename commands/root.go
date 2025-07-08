package commands

import (
	"gohook/commands/config_manager"
	"gohook/commands/setup_helper"
	"gohook/commands/webhook"
	"gohook/core/status_code"
	"gohook/utils"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "gohook",
	Short: "Webhook sender with Golang",
	Long:  "GoHook is a Command Line Interface for Discord webhook sender",
}

func setupCommands() {
	var webhookCommand = webhook.WebhookCommand()
	var genTomlConfigCommand = config_manager.GenTomlConfigCommand()

	setup_helper.NewCommandsList(webhookCommand, genTomlConfigCommand).SetupCommands(rootCommand)
}

func Execute() {
	setupCommands()

	if err := rootCommand.Execute(); err != nil {
		utils.CriticalShow("Details error: %s", err)
		os.Exit(status_code.CommandRunFailed)
	}
}
