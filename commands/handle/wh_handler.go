// This package for handle 'wh-send' command
package handle

import (
	whsendflags "gohook/commands/handle/wh_send_flags"
	embedsmanager "gohook/commands/handle/wh_send_flags/embeds_manager"
	"gohook/core"
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

func HandleWebhookSendCommand(params *CommandParameters) {
	if !core.FileExists(params.TomlConfigPath) {

		core.CriticalShow("File %s does not exists.", params.TomlConfigPath)
		os.Exit(core.FileNotFoundError)
	}

	var config core.DiscordWebhookConfig
	_, err := toml.DecodeFile(params.TomlConfigPath, &config)
	if err != nil {
		core.CriticalShow("Could not decode your TOML file: %s\n", err)
		os.Exit(core.TomlDecodeError)
	}

	var embeds = embedsmanager.GetEmbedsSetting(&config)
	var payload = core.DiscordWebhook{
		Content:  config.Message.Content,
		Username: config.Base.Username,
		Avatar:   config.Base.Avatar,
		Embeds:   embeds,
	}

	var dryRunContext = whsendflags.DryRunContext{
		EnableMode: params.DryMode,
	}
	dryRunContext.HandleDryRun(&payload)

	var webhookSendOnceContext = whsendflags.WebhookSendContext{
		IsDryMode:      params.DryMode,
		IsExplicitMode: params.Explicit,
		EnvURL:         params.EnvWebhookUrl,
		ConfigToml:     &config,
		LoopCount:      params.Loop,
	}
	webhookSendOnceContext.HandleWebhookSendOnce(&payload)

	var explicitContext = whsendflags.ExplicitContext{
		EnableExplicit: params.Explicit,
		EnvUrlName:     params.EnvWebhookUrl,
		Config:         &config,
	}
	explicitContext.HandleExplicitMode(&payload)

	var loopSendContext = whsendflags.LoopSendContext{
		DryMode:      params.DryMode,
		ExplicitMode: params.Explicit,
		LoopCount:    params.Loop,
		DelayTime:    params.Delay,
		EnvUrlName:   params.EnvWebhookUrl,
		Config:       &config,
	}
	loopSendContext.HandleLoopSend(&payload)

	var verboseContext = whsendflags.VerboseContext{
		EnableVerbose: params.Verbose,
		EnableDryRun:  params.DryMode,
	}
	verboseContext.HandleVerbose(&payload)

	var toJsonContext = whsendflags.ToJsonContext{
		IsEnableMode: params.ToJson,
	}
	toJsonContext.HandleExportToJson(&payload)
}
