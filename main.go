package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	conf := parseFlags()

	// Creates the note directory if it does not yet exist
	err := os.Mkdir(getBasePath(), 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error: Creating the note directory failed: %s\n", err)
		os.Exit(1)
	}

	if conf.displayAll {
		displayNotes("")
	} else if len(conf.displayDate) > 0 {
		displayNotes(conf.displayDate)
	} else if conf.edit.wantToEdit {
		editNote(conf.edit.date)
	} else if len(conf.search.needle) > 0 {
		searchNotes(&conf.search)
	} else {
		addNote(&conf.addNote)
	}
}

func getBasePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error: Determining the home directory failed: %s\n", err)
		os.Exit(1)
	}

	return path.Join(homeDir, "dn")
}
