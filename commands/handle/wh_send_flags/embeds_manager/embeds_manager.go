package embeds_manager

import (
	"gohook/commands/handle/wh_send_flags/helper"
	"gohook/core"
	"gohook/core/discord_api"
	"gohook/core/status_code"
	"gohook/dsl"
	"gohook/dsl/variables"
	"gohook/utils"
)

const (
	EmbedTitleMaxLen       = 256
	EmbedDescriptionMaxLen = 2048
	EmbedFooterMaxLen      = 2048
)

const (
	Description = "Description"
	Title       = "Title"
	Footer      = "Footer"
)

func CopyOptionalEmbedFields(src *core.DiscordEmbedConfig, dst *core.DiscordEmbed) {
	if utils.IsNonEmpty(src.Footer.Text) {
		dst.Footer = &discord_api.EmbedFooter{
			Text: src.Footer.Text,
		}
	}

	if utils.IsNonEmpty(src.Image.URL) {
		dst.Image = &discord_api.EmbedImage{
			URL: src.Image.URL,
		}
	}

	if utils.IsNonEmpty(src.Thumbnail.URL) {
		dst.Thumbnail = &discord_api.EmbedThumbnail{
			URL: src.Thumbnail.URL,
		}
	}
}

func GetEmbedsSetting(strictMode bool, config *core.DiscordWebhookConfig) []core.DiscordEmbed {
	var modeContext = dsl.ModeContext{
		StrictMode: strictMode,
	}
	var vars = variables.ParseVariables(&modeContext)

	for index := range config.Embeds {
		dsl.ParseVarsDiscordEmbed(&config.Embeds[index], vars)
	}

	embeds := make([]discord_api.Embed, 0, len(config.Embeds))

	for index := range config.Embeds {
		var embed = &config.Embeds[index]

		if embed.Title == "" && embed.Description == "" {
			continue
		}

		color, err := utils.GetHexColor(embed.Color)
		if err != nil {
			utils.InfoShow("Invalid color %s. Fallback to '0'\n", embed.Color)
			color = 0
		}

		helper.NewAssertEmbed(status_code.MaxLengthEmbedError).
			TryAssertEmbedLen(Description, embed.Description, EmbedFooterMaxLen).
			TryAssertEmbedLen(Title, embed.Title, EmbedTitleMaxLen).
			TryAssertEmbedLen(Footer, embed.Footer.Text, EmbedFooterMaxLen)

		var newEmbed = core.DiscordEmbed{
			Title:       embed.Title,
			Description: embed.Description,
			Color:       color,
		}

		CopyOptionalEmbedFields(embed, &newEmbed)
		embeds = append(embeds, newEmbed)
	}

	return embeds
}
