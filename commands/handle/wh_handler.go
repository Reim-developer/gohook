// This package for handle 'wh-send' command
package handle

import (
	"gohook/commands/handle/wh_send_flags"
	"gohook/commands/handle/wh_send_flags/embeds_manager"
	"gohook/core"
	"gohook/dsl"
	"gohook/dsl/variables"
	"gohook/utils"
	"os"

	"github.com/BurntSushi/toml"
)

type CommandParameters struct {
	TomlConfigPath string
	EnvWebhookUrl  string
	Verbose        bool
	DryMode        bool
	Threads        int
	Loop           int
	Delay          int
	Explicit       bool
	ToJson         bool
	StrictMode     bool
}

func getUserConfig(tomlConfigPath string) core.DiscordWebhookConfig {
	var config core.DiscordWebhookConfig

	_, err := toml.DecodeFile(tomlConfigPath, &config)
	if err != nil {
		utils.CriticalShow("Could not decode your TOML file: %s\n", err)
		os.Exit(core.TomlDecodeError)
	}

	return config
}

func loadVariables(strictMode bool, config *core.DiscordWebhookConfig) {
	var modeContext = dsl.ModeContext{
		StrictMode: strictMode,
	}
	var vars = variables.ParseVariables(&modeContext)
	dsl.ParseVarsDiscordMessage(config, vars)
}

func getDiscordPayload(config *core.DiscordWebhookConfig, embeds []core.DiscordEmbed) core.DiscordWebhook {
	var payload = core.DiscordWebhook{
		Content:  config.Message.Content,
		Username: config.Base.Username,
		Avatar:   config.Base.Avatar,
		Embeds:   embeds,
	}

	return payload
}

func setupFlags(params *CommandParameters) {
	var config = getUserConfig(params.TomlConfigPath)
	var embeds = embeds_manager.GetEmbedsSetting(params.StrictMode, &config)

	loadVariables(params.StrictMode, &config)
	var payload = getDiscordPayload(&config, embeds)

	wh_send_flags.NewDryRun(params.DryMode).HandleDryRun(&payload)
	wh_send_flags.NewWebhookSendOnce(
		params.DryMode, params.Explicit,
		params.EnvWebhookUrl, params.Loop, &config).HandleWebhookSendOnce(&payload)

	wh_send_flags.NewExplicit(params.Explicit, params.EnvWebhookUrl, &config).HandleExplicitMode(&payload)
	wh_send_flags.NewLoopSend(
		params.DryMode, params.Explicit, params.Loop,
		params.EnvWebhookUrl, params.Delay, &config).HandleLoopSend(&payload)

	wh_send_flags.NewVerbose(params.Verbose, params.DryMode).HandleVerbose(&payload)
	wh_send_flags.NewToJson(params.ToJson).HandleExportToJson(&payload)
}

func HandleWebhookSendCommand(params *CommandParameters) {
	if !utils.FileExists(params.TomlConfigPath) {

		utils.CriticalShow("File %s does not exists.", params.TomlConfigPath)
		os.Exit(core.FileNotFoundError)
	}

	setupFlags(params)
}
