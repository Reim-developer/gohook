package handle

import "gohook/commands/handle/gen_cfg_flags"

type GenCfgCommand struct {
	TomlConfigName string
}

func HandleGenCfgCommand(params *GenCfgCommand) {
	var genTomlConfigContext = gen_cfg_flags.GenTomlConfigContext{
		TomlConfigName: params.TomlConfigName,
	}

	genTomlConfigContext.GenTomlConfig()
}
