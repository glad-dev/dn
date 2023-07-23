package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func remove(date string, basePath string) int {
	filePath := filepath.Join(basePath, date)
	// Check if file exists
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Printf("There is no note for the date '%s'\n", date)

		return exitFailure
	}

	if info.IsDir() {
		fmt.Printf("'%s' is a directory, not a a file. Exiting\n", date)

		return exitFailure
	}

	err = os.Remove(filePath)
	if err != nil {
		fmt.Printf("Failed to remove file: %s\n", err)

		return exitFailure
	}

	return exitSuccess
}
