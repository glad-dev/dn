package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func addNote(noteInstance *note, basePath string) {
	fullPath := filepath.Join(basePath, noteInstance.date)
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()

	if err != nil {
		fmt.Printf("Error: Opening the note failed: %s\n", err)
		file.Close() // Needs to be added since defer is not called when calling os.Exit
		os.Exit(1)
	}

	_, err = file.WriteString(fmt.Sprintf("* %s\n", noteInstance.content))
	if err != nil {
		fmt.Printf("Error: Writing the note failed: %s\n", err)
		file.Close() // Ditto
		os.Exit(1)
	}
}
