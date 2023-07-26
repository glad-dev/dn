package cmd

import (
	"time"

	"github.com/GLAD-DEV/dn/constants"
	"github.com/GLAD-DEV/dn/note"
	"github.com/GLAD-DEV/dn/util"

	"github.com/spf13/cobra"
)

func newCmdToday() *cobra.Command {
	return &cobra.Command{
		Use:     "today",
		Aliases: []string{"t"},
		Short:   "Display today's notes.",
		Example: "dn today",
		Args:    cobra.ExactArgs(0),
		Run: func(_ *cobra.Command, args []string) {
			err := note.View(time.Now().Format(constants.DateFormat))
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}
		},
	}
}
