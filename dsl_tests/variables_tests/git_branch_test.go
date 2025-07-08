package variables_tests

import (
	"gohook/dsl"
	"gohook/dsl/variables/git_vars"
	"testing"
)

func TestGitBranchVariables(t *testing.T) {
	modeContext := dsl.ModeContext{
		StrictMode: true,
	}
	result := git_vars.GetBranchName(&modeContext)

	t.Log("Current Git branch is: ", result())
}
