// This package for handler 'wh-send' command
package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gohook/core"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type CommandParameters struct {
	TomlConfigPath string
	Verbose        bool
	DryMode        bool
	Threads        int
	Loop           int
	Delay          int
}

func HanderDryRun(dryMode bool, payload *core.DiscordWebhook) {
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

func HandlerVerbose(verbose bool, payload *core.DiscordWebhook) {
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

func HandlerLoopSend(url *string, payload *core.DiscordWebhook, params *CommandParameters) {
	if params.Loop > 1 {
		for i := range params.Loop {
			err := core.SendWebhook(url, payload)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[CRITICAL] Send webhook failed at loop %d: %s\n", i, err)
				continue
			}
			fmt.Fprintf(os.Stdout, "[OK] Send webhook success (%d) time(s), delay time: %d\n", i+1, params.Delay)
			time.Sleep(time.Duration(params.Delay) * time.Second)
		}

		HandlerVerbose(params.Verbose, payload)
		os.Exit(core.SUCCESS)
	}
}

func GetEmbedsSetting(config *core.Config) []core.Embed {
	embeds := make([]core.Embed, 0, len(config.Embeds))

	for _, embed := range config.Embeds {
		if embed.Title == "" && embed.Description == "" {
			continue
		}

		color, err := core.GetHexColor(embed.Color)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[WARN] Invalid color %s, fallback to '0'\n", embed.Color)
			color = 0
		}

		newEmbed := core.Embed{
			Title:       embed.Title,
			Description: embed.Description,
			Color:       color,
		}

		if embed.Footer.Text != "" {
			newEmbed.Footer = &core.EmbedFooter{
				Text: embed.Footer.Text,
			}
		}

		if embed.Image.URL != "" {
			newEmbed.Image = &core.EmbedImage{
				URL: embed.Image.URL,
			}
		}

		if embed.Thumbnail.URL != "" {
			newEmbed.Thumbnail = &core.EmbedThumbnail{
				URL: embed.Thumbnail.URL,
			}
		}

		embeds = append(embeds, newEmbed)
	}

	return embeds
}

func HandlerCommand(params *CommandParameters) {
	if !core.FileExists(params.TomlConfigPath) {

		fmt.Fprintf(os.Stderr, "[CRITICAL] File %s is not exists.\n", params.TomlConfigPath)
		os.Exit(core.FILE_NOT_FOUND)
	}

	var config core.Config
	_, err := toml.DecodeFile(params.TomlConfigPath, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CRITICAL] Could not decode your TOML file: %s\n", err)
		os.Exit(core.TOML_DECODE_ERROR)
	}

	var embeds = GetEmbedsSetting(&config)
	payload := core.DiscordWebhook{
		Content:  config.Message.Content,
		Username: config.Base.Username,
		Avatar:   config.Base.Avatar,
		Embeds:   embeds,
	}

	HanderDryRun(params.DryMode, &payload)
	HandlerLoopSend(config.Webhook.URL, &payload, params)
	err = core.SendWebhook(config.Webhook.URL, &payload)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[CRITICAL] Critical error: %s\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, "[OK] Successfully send webhook")

	HandlerVerbose(params.Verbose, &payload)
	os.Exit(core.SUCCESS)
}
