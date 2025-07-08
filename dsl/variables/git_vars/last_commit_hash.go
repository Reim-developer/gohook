package git_vars

import (
	"gohook/core"
	"gohook/dsl"
	"gohook/dsl/variables/dsl_helper"
	"gohook/utils"
)

func GetLastCommitHash(context *dsl.ModeContext) core.StrFunc {
	var function = func() string {

		result, err := utils.RunProgram(Git, RevParse, Head)
		dsl_helper.NewProgramError(err, context, Git).HandleProgramError()

		return string(result)
	}

	return function
}
