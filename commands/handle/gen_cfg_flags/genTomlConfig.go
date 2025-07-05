package gen_cfg_flags

import "gohook/core"

type GenTomlConfigContext struct {
	TomlConfigName string
}

func (context *GenTomlConfigContext) GenTomlConfig() {
	core.InfoShow("ok %s", context.TomlConfigName)
}
