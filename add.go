package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func addNote(noteInstance *note, basePath string) int {
	fullPath := filepath.Join(basePath, noteInstance.date)
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Printf("Error: Opening the note failed: %s\n", err)

		return exitFailure
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("* %s\n", noteInstance.content))
	if err != nil {
		fmt.Printf("Error: Writing the note failed: %s\n", err)

		return exitFailure
	}

	return exitSuccess
}
