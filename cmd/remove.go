package cmd

import (
	"fmt"

	"github.com/GLAD-DEV/dn/util"

	"github.com/spf13/cobra"

	"github.com/GLAD-DEV/dn/config"
	"github.com/GLAD-DEV/dn/note"
)

func newCmdRemove() *cobra.Command {
	return &cobra.Command{
		Use:     "remove [date]",
		Aliases: []string{"r"},
		Short:   "Removes the note dated to the passed date.",
		Long:    "Accepted date formats are 'YYYY-MM-DD' and 'DD-MM-YYYY'.",
		Example: "dn remove 2023-05-01",
		Args:    cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			date, err := parseDate(args[0])
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}

			err = note.Remove(date)
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}

			conf, err := config.Load()
			if err != nil {
				fmt.Printf("Sucessfully deleted the note.\n")
				fmt.Printf("But failed to read config: %s\n", err)

				return
			}

			if !conf.SilentRemove {
				fmt.Printf("Successfully deleted the note for '%s'\n", date)
			}
		},
	}
}
