package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	conf := parseFlags()
	basePath := getBasePath()

	// Creates the note directory if it does not yet exist
	err := os.Mkdir(basePath, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error: Creating the note directory failed: %s\n", err)
		os.Exit(1)
	}

	if conf.displayAll {
		displayNotes("", basePath)
	} else if len(conf.displayDate) > 0 {
		displayNotes(conf.displayDate, basePath)
	} else if conf.edit.wantToEdit {
		editNote(conf.edit.date, basePath)
	} else if len(conf.search.needle) > 0 {
		searchNotes(&conf.search, basePath)
	} else {
		addNote(&conf.addNote, basePath)
	}
}

func getBasePath() string {
	homeDir, ok := os.LookupEnv("DN_HOME")
	if ok && len(homeDir) > 0 {
		// DN_HOME is set
		// Check if path is absolute
		if !filepath.IsAbs(homeDir) {
			fmt.Printf("ERROR: DN_HOME (%s) is not an absolute path\n", homeDir)
			os.Exit(1)
		}

		// Check if path is
		_, err := os.Stat(homeDir)
		if err == nil || os.IsNotExist(err) {
			return homeDir
		}

		// os.Stat failed with an unknown error
		fmt.Printf("ERROR: DN_HOME is not a valid path: %s\n", err)
		os.Exit(1)
	}

	// No custom home directory is set => default to $HOME/dn/
	var err error
	homeDir, err = os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error: Determining the home directory failed: %s\n", err)
		os.Exit(1)
	}

	return filepath.Join(homeDir, "dn")
}
