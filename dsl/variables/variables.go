package variables

import (
	"gohook/core"
	"gohook/dsl/variables/utils_vars"
)

var GoHookVariables = core.VariablesFunc{
	"$TIME_NOW": utils_vars.TimeNow(),
}
