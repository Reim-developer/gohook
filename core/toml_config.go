package core

type Config struct {
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

	Embed struct {
		Title string `toml:"title"`
	} `toml:"Embed"`
}
