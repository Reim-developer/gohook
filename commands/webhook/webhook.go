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

type command_parameters struct {
	toml_config_path string
	verbose          bool
}

func handler_command_parameters(params *command_parameters) {
	if !core.FileExists(params.toml_config_path) {

		fmt.Fprintf(os.Stderr, "[CRITICAL] File %s is not exists.\n", params.toml_config_path)
		os.Exit(core.FILE_NOT_FOUND)
	}

	var config core.Config
	_, err := toml.DecodeFile(params.toml_config_path, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CRITICAL] Could not decode your TOML file: %s\n", err)
		os.Exit(core.TOML_DECODE_ERROR)
	}

	payload := core.DiscordWebhook{
		Content:  config.Message.Content,
		Username: config.Base.Username,
		Avatar:   config.Base.Avatar,
	}

	core.SendWebhook(config.Webhook.URL, &payload)
	fmt.Fprintln(os.Stdout, "[OK] Successfully send webhook")

	if params.verbose {
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
	os.Exit(core.SUCCESS)
}

func WebhookCommand() *cobra.Command {
	var toml_config_path string
	var verbose bool

	var webhook_command = &cobra.Command{
		Use:   "wh-send <TOML Config>",
		Short: "Send content to Discord webhook",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			toml_config_path = args[0]

			var arguments = command_parameters{
				toml_config_path: toml_config_path,
				verbose:          verbose,
			}

			handler_command_parameters(&arguments)
		},
	}

	webhook_command.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")
	return webhook_command

}
