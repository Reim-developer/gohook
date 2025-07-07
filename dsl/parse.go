package dsl

import (
	"gohook/core"
	"regexp"
	"strings"
)

func ReplaceVariables(str string, variables core.VariablesFunc) string {
	pattern := regexp.MustCompile(`\$[A-Z_][A-Z0-9_]*`)
	matches := pattern.FindAllString(str, -1)

	unique := make(map[string]struct{})
	for _, match := range matches {
		unique[match] = struct{}{}
	}

	for variable := range unique {
		if generateValue, exists := variables[variable]; exists {
			replacement := generateValue()

			str = strings.ReplaceAll(str, variable, replacement)
		}
	}

	return str
}

func ParseVarsDiscordEmbed(embed *core.DiscordEmbedConfig, variables core.VariablesFunc) {
	embed.Title = ReplaceVariables(embed.Title, variables)
	embed.Description = ReplaceVariables(embed.Description, variables)
	embed.Footer.Text = ReplaceVariables(embed.Footer.Text, variables)
}

func ParseVarsDiscordMessage(config *core.DiscordWebhookConfig, variables core.VariablesFunc) {
	config.Message.Content = ReplaceVariables(config.Message.Content, variables)
}
