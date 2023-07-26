package cmd

import (
	"github.com/spf13/cobra"

	"github.com/GLAD-DEV/dn/constants"
)

// Execute parses the command line parameters and starts the program.
func Execute() error {
	rootCmd := newRootCmd()

	rootCmd.Commands()[0].Name()

	return rootCmd.Execute()
}

func GetCommandNames() []string {
	cmds := newRootCmd().Commands()
	commands := make([]string, len(cmds))

	for i, cmd := range cmds {
		commands[i] = cmd.Name() // We can ignore alias since they are shorter than 3 chars
	}

	return commands
}

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "dn",
		Version: constants.Version,
	}

	cmdAdd := newCmdAdd()
	cmdEdit := newCmdEdit()
	cmdEditToday := newCmdEditToday()
	cmdOn := newCmdOn()
	cmdRemove := newCmdRemove()
	cmdSearch := newCmdSearch()
	cmdToday := newCmdToday()
	cmdView := newCmdView()
	cmdConfig := newCmdConfig()

	rootCmd.AddCommand(cmdAdd, cmdEdit, cmdEditToday, cmdOn, cmdRemove, cmdSearch, cmdToday, cmdView, cmdConfig)

	return rootCmd
}
