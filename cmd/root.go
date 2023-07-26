package cmd

import (
	"github.com/GLAD-DEV/dn/constants"

	"github.com/spf13/cobra"
)

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
