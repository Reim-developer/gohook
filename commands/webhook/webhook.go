package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gohook/core"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

type commandParameters struct {
	tomlConfigPath string
	verbose        bool
	dryMode        bool
	threads        int
	loop           int
	delay          int
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

func handlerLoopSend(url *string, payload *core.DiscordWebhook, params *commandParameters) {
	if params.loop > 1 {
		for i := range params.loop {
			err := core.SendWebhook(url, payload)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[CRITICAL] Send webhook failed at loop %d: %s\n", i, err)
				continue
			}
			fmt.Fprintf(os.Stdout, "[OK] Send webhook success (%d) time(s), delay time: %d\n", i+1, params.delay)
			time.Sleep(time.Duration(params.delay) * time.Second)
		}
		handlerVerbose(params.verbose, payload)
		os.Exit(core.SUCCESS)
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
	handlerLoopSend(config.Webhook.URL, &payload, params)
	err = core.SendWebhook(config.Webhook.URL, &payload)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[CRITICAL] Critical error: %s\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, "[OK] Successfully send webhook")

	handlerVerbose(params.verbose, &payload)
	os.Exit(core.SUCCESS)
}

func WebhookCommand() *cobra.Command {
	var toml_config_path string
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
			toml_config_path = args[0]

			var arguments = commandParameters{
				tomlConfigPath: toml_config_path,
				verbose:        verbose,
				dryMode:        dryMode,
				threads:        threads,
				loop:           loop,
				delay:          delay,
			}

			handlerCommand(&arguments)
		},
	}

	webhookCommand.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")
	webhookCommand.Flags().BoolVarP(&dryMode, "dry-run", "", false, "Enable dry-run mode")
	webhookCommand.Flags().IntVarP(&threads, "thread", "", 1, "Enable thread run")
	webhookCommand.Flags().IntVarP(&loop, "loop", "l", 1, "Enable loop run")
	webhookCommand.Flags().IntVarP(&delay, "delay", "", 2, "Enable delay")

	return webhookCommand

}
