package cmd

import (
	"unicode"

	"github.com/GLAD-DEV/dn/note"
	"github.com/GLAD-DEV/dn/util"

	"github.com/spf13/cobra"
)

func newCmdSearch() *cobra.Command {
	return &cobra.Command{
		Use:     "search [slug]",
		Aliases: []string{"s"},
		Short:   "Smart-case search",
		Long:    "If the argument contains at least one capital letter, the search is case-sensitive. Otherwise, it is case-insensitive.",
		Example: "dn search \"This search will be case-sensitive\"",
		Args:    cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			caseSensitive := false
			for _, c := range args[0] {
				if unicode.IsUpper(c) {
					caseSensitive = true

					break
				}
			}

			err := note.Search(args[0], caseSensitive)
			if err != nil {
				util.PrintErrAndExit(err.Error())
			}
		},
	}
}
