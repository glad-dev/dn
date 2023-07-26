package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/GLAD-DEV/dn/config"
	"github.com/GLAD-DEV/dn/constants"
	"github.com/GLAD-DEV/dn/note"
	"github.com/GLAD-DEV/dn/util"

	"github.com/spf13/cobra"
)

// Execute parses the command line parameters and starts the program.
func Execute() error {
	rootCmd := newRootCmd()

	done := handleShortAdd(rootCmd)
	if done {
		return nil
	}

	return rootCmd.Execute()
}

// handleShortAdd handles cases like `dn "This is a note"`. Returns whether the program is done.
func handleShortAdd(rootCmd *cobra.Command) bool {
	// We only care about
	if len(os.Args) != 2 {
		return false
	}

	// Parameter could be a (mistyped) command or an alias
	if len(os.Args[1]) <= 3 {
		return false
	}

	// Get list of commands
	commands := make([]string, len(rootCmd.Commands()))
	for i, cmd := range rootCmd.Commands() {
		commands[i] = cmd.Name() // We can ignore alias since they are shorter than 3 chars
	}

	// Add the completion command
	commands = append(commands, "completion")

	// Add flags
	commands = append(commands, "--help")
	commands = append(commands, "--version")

	// Ensure that we don't accidentally interpret commands as notes
	for _, name := range commands {
		if os.Args[1] == name {
			return false
		}
	}

	err := note.Add(os.Args[1], time.Now().Format(constants.DateFormat))
	if err != nil {
		util.PrintErrAndExit(err.Error())
	}

	conf, err := config.Load()
	if err != nil {
		fmt.Printf("Added '%s' to today's note\n", os.Args[1])
		fmt.Println("If this was not intended, run `dn et` to edit today's note")
		fmt.Printf("Failed to read config: %s\n", err)
	}

	if !conf.SilentAdd {
		fmt.Printf("Added '%s' to today's note\n", os.Args[1])
		fmt.Println("If this was not intended, run `dn et` to edit today's note")
	}

	return true
}
