package cmd

import (
	"github.com/GLAD-DEV/dn/note"
	"github.com/GLAD-DEV/dn/util"

	"github.com/spf13/cobra"
)

func newCmdOn() *cobra.Command {
	return &cobra.Command{
		Use:     "on [date] [note]",
		Aliases: []string{"o"},
		Short:   "Add note dated at passed date.",
		Long:    "Accepted date formats are 'YYYY-MM-DD' and 'DD-MM-YYYY'.",
		Example: "dn on 2023-05-01 \"This note will be added for the May 1 2005\"",
		Args:    cobra.ExactArgs(2),
		Run: func(_ *cobra.Command, args []string) {
			date, err := parseDate(args[0])
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}

			err = note.Add(args[1], date)
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}
		},
	}
}
