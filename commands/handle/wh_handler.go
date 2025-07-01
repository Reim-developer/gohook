// This package for handle 'wh-send' command
package handle

import (
	"bytes"
	"encoding/json"
	"fmt"
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
}

func HandleDryRun(dryMode bool, payload *core.DiscordWebhook) {
	if dryMode {
		var buffer bytes.Buffer
		var encoder = json.NewEncoder(&buffer)

		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", " ")

		err := encoder.Encode(payload)
		if err != nil {

			fmt.Fprintf(os.Stderr, "[CRITICAL] Could not decode JSON: %s\n", err)
			os.Exit(core.JsonDecodeError)
		}

		fmt.Fprintln(os.Stdout, "[INFO] Running in Dry Mode:")
		fmt.Fprintf(os.Stdout, "[INFO] Your webhook payload:\n%s\n", buffer.String())
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

			fmt.Fprintf(os.Stderr, "[CRITICAL] Could not decode JSON: %s\n", err)
			os.Exit(core.JsonDecodeError)
		}

		fmt.Fprintf(os.Stdout, "[INFO] Your webhook payload:\n%s\n", buffer.String())
	}
}

func HandleLoopSend(url *string, payload *core.DiscordWebhook, params *CommandParameters) {
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
		fmt.Fprintf(os.Stderr, "[CRITICAL] Embed [%s] too long (%d characters).\n", max.fieldName, length)
		fmt.Fprintf(os.Stderr, "[HINT] Discord only allows up to %d character.\n", max.maxLen)
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
			fmt.Fprintf(os.Stderr, "[WARN] Invalid color %s, fallback to '0'\n", embed.Color)
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

func HandleCommand(params *CommandParameters) {
	if !core.FileExists(params.TomlConfigPath) {

		fmt.Fprintf(os.Stderr, "[CRITICAL] File %s does not exists.\n", params.TomlConfigPath)
		os.Exit(core.FileNotFoundError)
	}

	var config core.DiscordWebhookConfig
	_, err := toml.DecodeFile(params.TomlConfigPath, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CRITICAL] Could not decode your TOML file: %s\n", err)
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
	HandleLoopSend(config.Webhook.URL, &payload, params)
	err = core.SendWebhook(config.Webhook.URL, &payload)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[CRITICAL] Critical error: %s\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, "[OK] Successfully send webhook")

	HandleVerbose(params.Verbose, &payload)
	os.Exit(core.Success)
}
