package commands

import (
	"gohook/commands/webhook"
	"gohook/core"
	"log"
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

	if err := rootCommand.Execute(); err != nil {
		log.Println(err)
		os.Exit(core.CommandRunFailed)
	}
}
