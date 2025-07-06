package core_test

import (
	"encoding/json"
	"fmt"
	"gohook/core"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
)

type StrFunc = func() string
type VariablesFunc = map[string]func() string

func TimeNow() StrFunc {
	var function = func() string {
		return time.Now().Format("2006-01-02 15:04:05")
	}

	return function
}

var ConfigVariables = VariablesFunc{
	"$TIME_NOW": TimeNow(),
}

func ReplaceVariables(str string, variables VariablesFunc) string {
	pattern := regexp.MustCompile(`\$[A-Z_][A-Z0-9_]*`)
	matches := pattern.FindAllString(str, -1)

	unique := make(map[string]struct{})
	for _, match := range matches {
		unique[match] = struct{}{}
	}

	for variable := range unique {
		if generateValue, exists := variables[variable]; exists {
			replacement := generateValue()

			str = strings.ReplaceAll(str, variable, replacement)
		}
	}

	return str
}

func ReplaceVariablesInEmbeds(embed *core.DiscordEmbedConfig, variables VariablesFunc) {
	embed.Title = ReplaceVariables(embed.Title, variables)
	embed.Description = ReplaceVariables(embed.Description, variables)
	embed.Footer.Text = ReplaceVariables(embed.Footer.Text, variables)
}

func ReplaceVariableInMessage(config *core.DiscordWebhookConfig, variables VariablesFunc) {
	config.Message.Content = ReplaceVariables(config.Message.Content, variables)
}

func TestMetadataToml(t *testing.T) {
	var config core.DiscordWebhookConfig
	_, err := toml.DecodeFile("settings_test.toml", &config)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[Core_Test Package] Error: %s\n", err)
		os.Exit(1)
	}

	ReplaceVariableInMessage(&config, ConfigVariables)
	for index := range config.Embeds {
		ReplaceVariablesInEmbeds(&config.Embeds[index], ConfigVariables)
	}

	jsonData, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		t.Fatal("[Core_Test Package] In Function 'TestMetaDataTOML':")
		t.Fatal(">>>", err)
	}

	t.Log(string(jsonData))
}
