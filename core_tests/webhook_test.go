package core_test

import (
	"fmt"
	"gohook/core"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
)

func TestDecodeToml(t *testing.T) {
	var config core.Config
	_, err := toml.DecodeFile("settings_test.toml", &config)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[Core_Test Package] Error: %s\n", err)
		os.Exit(1)
	}
}
