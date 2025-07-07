// This package for handle 'wh-send' command
package handle

import (
	"gohook/commands/handle/wh_send_flags"
	"gohook/commands/handle/wh_send_flags/embeds_manager"
	"gohook/core"
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

func setupFlags(params *CommandParameters) {
	var config = getUserConfig(params.TomlConfigPath)
	var embeds = embeds_manager.GetEmbedsSetting(&config)
	var payload = core.DiscordWebhook{
		Content:  config.Message.Content,
		Username: config.Base.Username,
		Avatar:   config.Base.Avatar,
		Embeds:   embeds,
	}

	var dryRunContext = wh_send_flags.DryRunContext{
		EnableMode: params.DryMode,
	}
	dryRunContext.HandleDryRun(&payload)

	var webhookSendOnceContext = wh_send_flags.WebhookSendContext{
		IsDryMode:      params.DryMode,
		IsExplicitMode: params.Explicit,
		EnvURL:         params.EnvWebhookUrl,
		ConfigToml:     &config,
		LoopCount:      params.Loop,
	}
	webhookSendOnceContext.HandleWebhookSendOnce(&payload)

	var explicitContext = wh_send_flags.ExplicitContext{
		EnableExplicit: params.Explicit,
		EnvUrlName:     params.EnvWebhookUrl,
		Config:         &config,
	}
	explicitContext.HandleExplicitMode(&payload)

	var loopSendContext = wh_send_flags.LoopSendContext{
		DryMode:      params.DryMode,
		ExplicitMode: params.Explicit,
		LoopCount:    params.Loop,
		DelayTime:    params.Delay,
		EnvUrlName:   params.EnvWebhookUrl,
		Config:       &config,
	}
	loopSendContext.HandleLoopSend(&payload)

	var verboseContext = wh_send_flags.VerboseContext{
		EnableVerbose: params.Verbose,
		EnableDryRun:  params.DryMode,
	}
	verboseContext.HandleVerbose(&payload)

	var toJsonContext = wh_send_flags.ToJsonContext{
		IsEnableMode: params.ToJson,
	}
	toJsonContext.HandleExportToJson(&payload)
}

func HandleWebhookSendCommand(params *CommandParameters) {
	if !utils.FileExists(params.TomlConfigPath) {

		utils.CriticalShow("File %s does not exists.", params.TomlConfigPath)
		os.Exit(core.FileNotFoundError)
	}

	setupFlags(params)
}
