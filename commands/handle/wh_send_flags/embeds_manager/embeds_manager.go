package embedsmanager

import (
	"gohook/commands/handle/wh_send_flags/helper"
	"gohook/core"
)

const (
	EmbedTitleMaxLen       = 256
	EmbedDescriptionMaxLen = 2048
	EmbedFooterMaxLen      = 2048
)

func CopyOptionalEmbedFields(src *core.DiscordEmbedConfig, dst *core.Embed) {
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
			URL: src.Thumbnail.URL,
		}
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

		var newEmbed = core.Embed{
			Title:       embed.Title,
			Description: embed.Description,
			Color:       color,
		}

		CopyOptionalEmbedFields(embed, &newEmbed)
		embeds = append(embeds, newEmbed)
	}

	return embeds
}
