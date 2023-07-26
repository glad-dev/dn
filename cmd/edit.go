package cmd

import (
	"github.com/GLAD-DEV/dn/note"
	"github.com/GLAD-DEV/dn/util"

	"github.com/spf13/cobra"
)

func newCmdEdit() *cobra.Command {
	return &cobra.Command{
		Use:     "edit {date}",
		Aliases: []string{"e"},
		Short:   "If no date is passed, $EDITOR's file selection is opened. Otherwise the corresponding note is opened in $EDITOR.",
		Long:    "Accepted date formats are 'YYYY-MM-DD' and 'DD-MM-YYYY'.",
		Example: "dn edit 2023-01-02",
		Args:    cobra.MaximumNArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			date := ""
			var err error
			if len(args) == 1 {
				date, err = parseDate(args[0])
				if err != nil {
					util.PrintErrAndExit(err.Error())
				}
			}

			err = note.Edit(date)
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}
		},
	}
}
