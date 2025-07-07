package discord_api

type DiscordWebhookConfig struct {
	Webhook struct {
		URL  *string  `toml:"url"`
		URLs []string `toml:"urls"`
	} `toml:"Webhook"`

	Base struct {
		Username string `toml:"username"`
		Avatar   string `toml:"avatar_url"`
	} `toml:"Base"`

	Message struct {
		Content string `toml:"content"`
	} `toml:"Message"`

	Embeds []DiscordEmbedConfig `toml:"Embeds"`
}

type DiscordEmbedConfig struct {
	Title       string `toml:"title"`
	Description string `toml:"description"`
	Color       string `toml:"color"`

	Footer struct {
		Text string `toml:"text"`
	} `toml:"footer"`

	Image struct {
		URL string `toml:"url"`
	} `toml:"image"`

	Thumbnail struct {
		URL string `toml:"url"`
	} `toml:"thumbnail"`
}
