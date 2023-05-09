package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func editNote(fileName string, basePath string) int {
	filePath := filepath.Join(basePath, fileName)

	// Check if file exists
	f, err := os.Open(filePath)
	if os.IsNotExist(err) {
		fmt.Printf("Error: No note dated %s exists\n", fileName)

		return exitFailure
	}
	_ = f.Close()

	editor, exists := os.LookupEnv("EDITOR")
	if !exists || len(editor) == 0 {
		editor = "vim"
	}

	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error: Running vim failed: %s\n", err)

		return exitFailure
	}

	return exitSuccess
}
