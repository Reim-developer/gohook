package handle

import "gohook/commands/handle/gen_cfg_flags"

type GenCfgCommand struct {
	TomlConfigName string
}

func HandleGenCfgCommand(params *GenCfgCommand) {
	gen_cfg_flags.NewGenTomlConfig(params.TomlConfigName).GenTomlConfig()
}
