package cmd

import (
	"time"

	"github.com/GLAD-DEV/dn/constants"
	"github.com/GLAD-DEV/dn/note"
	"github.com/GLAD-DEV/dn/util"

	"github.com/spf13/cobra"
)

func newCmdAdd() *cobra.Command {
	return &cobra.Command{
		Use:     "add [note]",
		Aliases: []string{"a"},
		Short:   "Adds note with today's date. Can be omitted if note is longer than three chars.",
		Example: "dn add \"This note will be added with today's date\"",
		Args:    cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			err := note.Add(args[0], time.Now().Format(constants.DateFormat))
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}
		},
	}
}
