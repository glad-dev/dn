package main

import (
	"fmt"
	"github.com/GLAD-DEV/dn/util"
	"os"
	"time"

	"github.com/GLAD-DEV/dn/cmd"
	"github.com/GLAD-DEV/dn/config"
	"github.com/GLAD-DEV/dn/constants"
	"github.com/GLAD-DEV/dn/note"
)

func main() {
	err := config.CreateDir()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	done := handleShortAdd()
	if done {
		return
	}

	_ = cmd.Execute()
}

// handleShortAdd handles cases like `dn "This is a note"`. Returns whether the program is done.
func handleShortAdd() bool {
	// We only care about
	if len(os.Args) != 2 {
		return false
	}

	// Parameter could be a (mistyped) command or an alias
	if len(os.Args[1]) <= 3 {
		return false
	}

	names := cmd.GetCommandNames()
	// Add flags
	names = append(names, "--help")
	names = append(names, "--version")

	// Ensure that we don't accidentally interpret commands as notes
	for _, name := range names {
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
