// This package for handle 'wh-send' command
package handle

import (
	"bytes"
	"encoding/json"
	"gohook/core"
	"os"
	"time"
	"unicode/utf8"

	"github.com/BurntSushi/toml"
)

const (
	EmbedTitleMaxLen       = 256
	EmbedDescriptionMaxLen = 2048
	EmbedFooterMaxLen      = 2048
)

type CommandParameters struct {
	TomlConfigPath string
	Verbose        bool
	DryMode        bool
	Threads        int
	Loop           int
	Delay          int
	Explicit       bool
}

func HandleDryRun(dryMode bool, payload *core.DiscordWebhook) {
	if dryMode {
		var buffer bytes.Buffer
		var encoder = json.NewEncoder(&buffer)

		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", " ")

		err := encoder.Encode(payload)
		if err != nil {

			core.CriticalShow("Could not decode JSON: %s", err)
			os.Exit(core.JsonDecodeError)
		}

		core.InfoShow("Running in Dry Mode:")
		core.InfoShow("Your webhook payload:\n%s", buffer.String())
		os.Exit(core.Success)
	}
}

func HandleVerbose(verbose bool, payload *core.DiscordWebhook) {
	if verbose {
		var buffer bytes.Buffer
		var encoder = json.NewEncoder(&buffer)

		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", " ")

		err := encoder.Encode(payload)
		if err != nil {

			core.CriticalShow("Could not decode JSON: %s", err)
			os.Exit(core.JsonDecodeError)
		}

		core.InfoShow("Your webhook payload:\n%s", buffer.String())
	}
}

func HandleLoopSend(url *string, payload *core.DiscordWebhook, params *CommandParameters) {
	if params.Loop > 1 {
		var successCount = 0
		var failedCount = 0

		for i := range params.Loop {
			err := core.SendWebhook(url, payload)

			if err != nil {
				core.CriticalShow("Send webhook failed (%d) time(s) %s\n", i, err)

				failedCount += 1
				time.Sleep(time.Duration(params.Delay) * time.Second)
				continue
			}
			core.InfoShow("Send webhook success (%d) time(s), delay time: %d", i+1, params.Delay)

			successCount += 1
			time.Sleep(time.Duration(params.Delay) * time.Second)
		}

		core.InfoShow("Success count: %d time(s)", successCount)
		core.InfoShow("Failed count: %d time(s)", failedCount)

		HandleVerbose(params.Verbose, payload)
		os.Exit(core.Success)
	}
}

func copyOptionalEmbedFields(src *core.DiscordEmbedConfig, dst *core.Embed) {
	if core.IsNonEmpty(src.Footer.Text) {
		dst.Footer = &core.EmbedFooter{
			Text: src.Footer.Text,
		}
	}

	if core.IsNonEmpty(src.Image.URL) {
		dst.Image = &core.EmbedImage{
			URL: src.Image.URL,
		}
	}

	if core.IsNonEmpty(src.Thumbnail.URL) {
		dst.Thumbnail = &core.EmbedThumbnail{
			URL: src.Image.URL,
		}
	}
}

type maxLenCheck struct {
	fieldName string
	value     string
	maxLen    int
	exitCode  int
}

func assertMaxLen(max *maxLenCheck) {
	var length = utf8.RuneCountInString(max.value)

	if length > max.maxLen {
		core.CriticalShow("Embed [%s] too long (%d characters).", max.fieldName, length)
		core.InfoShow("Discord only allows up to %d characters.", max.maxLen)
		os.Exit(max.exitCode)
	}
}

func GetEmbedsSetting(config *core.DiscordWebhookConfig) []core.Embed {
	embeds := make([]core.Embed, 0, len(config.Embeds))

	for index := range config.Embeds {
		var embed = &config.Embeds[index]

		if embed.Title == "" && embed.Description == "" {
			continue
		}

		color, err := core.GetHexColor(embed.Color)
		if err != nil {
			core.InfoShow("Invalid color %s. Fallback to '0'\n", embed.Color)
			color = 0
		}

		assertMaxLen(&maxLenCheck{
			fieldName: "Descripton",
			value:     embed.Description,
			maxLen:    EmbedDescriptionMaxLen,
			exitCode:  core.DescriptionMaxLenError,
		})

		assertMaxLen(&maxLenCheck{
			fieldName: "Title",
			value:     embed.Title,
			maxLen:    EmbedTitleMaxLen,
			exitCode:  core.TitleMaxLenError,
		})

		assertMaxLen(&maxLenCheck{
			fieldName: "Footer",
			value:     embed.Footer.Text,
			maxLen:    EmbedFooterMaxLen,
			exitCode:  core.FooterMaxLenError,
		})

		var newEmbed = core.Embed{
			Title:       embed.Title,
			Description: embed.Description,
			Color:       color,
		}

		copyOptionalEmbedFields(embed, &newEmbed)
		embeds = append(embeds, newEmbed)
	}

	return embeds
}

func HandleWebhookSendOnce(config *core.DiscordWebhookConfig, payload *core.DiscordWebhook, params *CommandParameters) {
	var err = core.SendWebhook(config.Webhook.URL, payload)

	if err != nil {
		core.CriticalShow("Critical error: %s\n", err)
		os.Exit(core.WebhookSendFailed)
	}

	core.InfoShow("Successfully send webhook")

	HandleVerbose(params.Verbose, payload)
	os.Exit(core.Success)
}

func HandleExplicitMode(config *core.DiscordWebhookConfig, payload *core.DiscordWebhook, params *CommandParameters) {
	var result, err = core.ExplicitSendWebhook(config.Webhook.URL, payload)
	if err != nil {
		core.CriticalShow("Could not send webhook with error: %s", err)
		os.Exit(core.WebhookSendFailed)
	}

	core.InfoShow("Successfully send webhook")
	core.InfoShow("Message ID: %s", result.MessageID)
	core.InfoShow("Channel ID: %s", result.ChannelID)

	HandleVerbose(params.Verbose, payload)
	os.Exit(core.Success)
}

func HandleCommand(params *CommandParameters) {
	if !core.FileExists(params.TomlConfigPath) {

		core.CriticalShow("File %s does not exists.", params.TomlConfigPath)
		os.Exit(core.FileNotFoundError)
	}

	var config core.DiscordWebhookConfig
	_, err := toml.DecodeFile(params.TomlConfigPath, &config)
	if err != nil {
		core.CriticalShow("Could not decode your TOML file: %s\n", err)
		os.Exit(core.TomlDecodeError)
	}

	var embeds = GetEmbedsSetting(&config)
	payload := core.DiscordWebhook{
		Content:  config.Message.Content,
		Username: config.Base.Username,
		Avatar:   config.Base.Avatar,
		Embeds:   embeds,
	}

	HandleDryRun(params.DryMode, &payload)
	HandleExplicitMode(&config, &payload, params)
	HandleLoopSend(config.Webhook.URL, &payload, params)
	HandleWebhookSendOnce(&config, &payload, params)
}
