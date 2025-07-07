package embeds_manager

import (
	"gohook/commands/handle/wh_send_flags/helper"
	"gohook/core"
	"gohook/core/discord_api"
	"gohook/dsl"
	"gohook/dsl/variables"
	"gohook/utils"
)

const (
	EmbedTitleMaxLen       = 256
	EmbedDescriptionMaxLen = 2048
	EmbedFooterMaxLen      = 2048
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

func GetEmbedsSetting(config *core.DiscordWebhookConfig) []core.DiscordEmbed {
	for index := range config.Embeds {
		dsl.ParseVarsDiscordEmbed(&config.Embeds[index], variables.GoHookVariables)
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

		helper.AssertMaxLen(&helper.MaxLenCheck{
			FieldName: "Descripton",
			Value:     embed.Description,
			MaxLen:    EmbedDescriptionMaxLen,
			ExitCode:  core.DescriptionMaxLenError,
		})

		helper.AssertMaxLen(&helper.MaxLenCheck{
			FieldName: "Title",
			Value:     embed.Title,
			MaxLen:    EmbedTitleMaxLen,
			ExitCode:  core.TitleMaxLenError,
		})

		helper.AssertMaxLen(&helper.MaxLenCheck{
			FieldName: "Footer",
			Value:     embed.Footer.Text,
			MaxLen:    EmbedFooterMaxLen,
			ExitCode:  core.FooterMaxLenError,
		})

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
