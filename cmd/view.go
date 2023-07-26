package cmd

import (
	"fmt"
	"os"

	"github.com/GLAD-DEV/dn/note"

	"github.com/spf13/cobra"
)

func newCmdView() *cobra.Command {
	return &cobra.Command{
		Use:     "view {date}",
		Aliases: []string{"v"},
		Short:   "If no date is passed, displays all notes. Otherwise displays notes matching passed date slug.",
		Long:    "The accepted date formats are 'YYYY-MM-DD' and 'DD-MM-YYYY'",
		Example: "dn view 2023-05-01",
		Args:    cobra.MaximumNArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			slug := ""
			if len(args) == 1 {
				slug = args[0]
			}

			err := note.View(slug)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
		},
	}
}
