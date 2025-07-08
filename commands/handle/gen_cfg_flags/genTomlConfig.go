package gen_cfg_flags

import (
	"bufio"
	"gohook/core"
	"gohook/utils"
	"os"
)

const BasedTomlConfig = `# Webhook URL is required.
[Webhook]
# For more information, how to get Discord webhook URL, please visit
# https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks
url = "My Discord Webhook URL"

# Base information setup for webhook
[Base]
username = ""   # Your webhook username. Leave blank for default username
avatar_url = "" # Your webhook avatar. Leave blank for default avatar

# Webhook simple content. With no embeds.
[Message]
content = "Hello World. From GoHook !" # Your first content. Awesome!
`

type genTomlConfigContext struct {
	tomlConfigName string
}

func NewGenTomlConfig(tomlConfigName string) *genTomlConfigContext {
	var genTomlConfigContext = genTomlConfigContext{
		tomlConfigName: tomlConfigName,
	}

	return &genTomlConfigContext
}

func (context *genTomlConfigContext) GenTomlConfig() {
	file, createFailed := os.Create(context.tomlConfigName)
	if createFailed != nil {
		utils.CriticalShow("Could not create %s", context.tomlConfigName)
		utils.CriticalShow("Full error message: %s", createFailed)

		os.Exit(core.CreateFileFailed)
	}
	defer file.Close()

	var writer = bufio.NewWriter(file)
	bytesWrite, writeFailed := writer.WriteString(BasedTomlConfig)
	if writeFailed != nil {
		utils.CriticalShow("Could not create %s", context.tomlConfigName)
		utils.CriticalShow("Full error message: %s", writeFailed)

		os.Exit(core.WriteFileFailed)
	}

	flushFailed := writer.Flush()
	if flushFailed != nil {
		utils.CriticalShow("Could not flush data to %s", context.tomlConfigName)
		utils.CriticalShow("Full error message: %s", flushFailed)

		os.Exit(core.FlushFileFailed)
	}

	utils.InfoShow("Successfully generated: %s (%d bytes written)", context.tomlConfigName, bytesWrite)
}
