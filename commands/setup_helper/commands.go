package setup_helper

import "github.com/spf13/cobra"

type commandsListContext struct {
	commands []*cobra.Command
}

func NewCommandsList(commands ...*cobra.Command) *commandsListContext {
	var commandsList = commandsListContext{
		commands: commands,
	}

	return &commandsList
}

func (commandsContext *commandsListContext) SetupCommands(rootCommand *cobra.Command) {
	for _, command := range commandsContext.commands {
		if command != nil {
			rootCommand.AddCommand(command)
		}
	}
}
