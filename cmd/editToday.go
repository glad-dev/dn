package cmd

import (
	"time"

	"github.com/GLAD-DEV/dn/constants"
	"github.com/GLAD-DEV/dn/note"
	"github.com/GLAD-DEV/dn/util"

	"github.com/spf13/cobra"
)

func newCmdEditToday() *cobra.Command {
	return &cobra.Command{
		Use:     "editToday",
		Aliases: []string{"et"},
		Short:   "Edit today's note in $EDITOR.",
		Example: "dn editToday",
		Args:    cobra.ExactArgs(0),
		Run: func(_ *cobra.Command, args []string) {
			err := note.Edit(time.Now().Format(constants.DateFormat))
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}
		},
	}
}
