package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Execute parses the command line parameters and starts the program.
func Execute() error {
	rootCmd := newRootCmd()

	if isShortAdd(rootCmd, os.Args[1:]) {
		// `dn "A note"` -> `dn add "A note"`
		rootCmd.SetArgs(append([]string{"add"}, os.Args[1:]...))
	}

	return rootCmd.Execute()
}

// isShortAdd reports whether args is a note without the add command, i.e. `dn "A note"`.
func isShortAdd(rootCmd *cobra.Command, args []string) bool {
	if len(args) != 1 || args[0] == "" {
		return false
	}

	// cobra only adds these commands during Execute, so Find would not know them otherwise
	rootCmd.InitDefaultHelpCmd()
	rootCmd.InitDefaultCompletionCmd()

	// Find fails for unknown commands. Flags like --help resolve to the root command without an error.
	_, _, err := rootCmd.Find(args)

	return err != nil
}
