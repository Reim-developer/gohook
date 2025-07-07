package core

import (
	"gohook/core/discord_api"

	"github.com/spf13/cobra"
)

// [!] Type alias of Cobra closure.
type CobraClosure = func(cmd *cobra.Command, args []string)

// [!] Type alias of string closure function.
type StrFunc = func() string

// [!] Type alias of map string closure function.
type VariablesFunc = map[string]func() string

// [!] Discord Embed is a struct of JSON data of Discord embed payload.
type DiscordEmbed = discord_api.Embed

// [!] Discord Webhook is a struct of JSON data of Discord webhook payload.
type DiscordWebhook = discord_api.DiscordWebhook

// [!] Discord Embed Config is a struct TOML data of user-defined.
type DiscordEmbedConfig = discord_api.DiscordEmbedConfig

// [!] Discord Webhook Config is a struct TOML data of user-defined.
type DiscordWebhookConfig = discord_api.DiscordWebhookConfig
