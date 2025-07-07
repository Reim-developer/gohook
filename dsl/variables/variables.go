package variables

import (
	"gohook/core"
	"gohook/dsl/variables/utils_vars"
)

var GoHookVariables = core.VariablesFunc{
	"$TIME_NOW":      utils_vars.TimeNow(),
	"$USER_HOME":     utils_vars.UserHome(),
	"$USER_OS":       utils_vars.UserOS(),
	"$USER_HOSTNAME": utils_vars.HostName(),
}
