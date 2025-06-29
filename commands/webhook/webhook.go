package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gohook/core"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

type commandParameters struct {
	tomlConfigPath string
	verbose        bool
	dryMode        bool
}

func handerDryRun(dryMode bool, payload *core.DiscordWebhook) {
	if dryMode {
		var buffer bytes.Buffer
		var encoder = json.NewEncoder(&buffer)

		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", " ")

		err := encoder.Encode(payload)
		if err != nil {

			fmt.Fprintf(os.Stderr, "[CRITICAL] Could not decode JSON: %s\n", err)
			os.Exit(core.JSON_DECODE_ERROR)
		}

		fmt.Fprintln(os.Stdout, "[INFO] Running in Dry Mode:")
		fmt.Fprintf(os.Stdout, "[INFO] Your webhook payload:\n%s\n", buffer.String())
		os.Exit(core.SUCCESS)
	}
}

func handlerVerbose(verbose bool, payload *core.DiscordWebhook) {
	if verbose {
		var buffer bytes.Buffer
		var encoder = json.NewEncoder(&buffer)

		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", " ")

		err := encoder.Encode(payload)
		if err != nil {

			fmt.Fprintf(os.Stderr, "[CRITICAL] Could not decode JSON: %s\n", err)
			os.Exit(core.JSON_DECODE_ERROR)
		}

		fmt.Fprintf(os.Stdout, "[INFO] Your webhook payload:\n%s\n", buffer.String())
	}
}

func handlerCommand(params *commandParameters) {
	if !core.FileExists(params.tomlConfigPath) {

		fmt.Fprintf(os.Stderr, "[CRITICAL] File %s is not exists.\n", params.tomlConfigPath)
		os.Exit(core.FILE_NOT_FOUND)
	}

	var config core.Config
	_, err := toml.DecodeFile(params.tomlConfigPath, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CRITICAL] Could not decode your TOML file: %s\n", err)
		os.Exit(core.TOML_DECODE_ERROR)
	}

	payload := core.DiscordWebhook{
		Content:  config.Message.Content,
		Username: config.Base.Username,
		Avatar:   config.Base.Avatar,
	}

	handerDryRun(params.dryMode, &payload)
	core.SendWebhook(config.Webhook.URL, &payload)

	fmt.Fprintln(os.Stdout, "[OK] Successfully send webhook")

	handlerVerbose(params.verbose, &payload)
	os.Exit(core.SUCCESS)
}

func WebhookCommand() *cobra.Command {
	var toml_config_path string
	var verbose bool
	var dryMode bool

	var webhookCommand = &cobra.Command{
		Use:   "wh-send <TOML Config>",
		Short: "Send content to Discord webhook",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			toml_config_path = args[0]

			var arguments = commandParameters{
				tomlConfigPath: toml_config_path,
				verbose:        verbose,
				dryMode:        dryMode,
			}

			handlerCommand(&arguments)
		},
	}

	webhookCommand.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")
	webhookCommand.Flags().BoolVarP(&dryMode, "dry-run", "", false, "Enable dry-run mode")
	return webhookCommand

}
