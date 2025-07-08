package git_vars

import (
	"gohook/core"
	"gohook/dsl"
	"gohook/dsl/variables/dsl_helper"
	"gohook/utils"
)

func handleClosure(modeContext *dsl.ModeContext) core.StrFunc {
	var function = func() string {

		result, err := utils.RunProgram(Git, RevParse, "--abbrev-ref", Head)
		dsl_helper.NewProgramError(err, modeContext, Git).HandleProgramError()

		return result
	}

	return function
}

func GetBranchName(modeContext *dsl.ModeContext) core.StrFunc {
	var function = handleClosure(modeContext)

	return function
}
