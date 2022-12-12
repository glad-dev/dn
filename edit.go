package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func editNote(fileName string) {
	filePath := path.Join(getBasePath(), fileName)

	// Check if file exists
	f, err := os.Open(filePath)
	f.Close()
	if os.IsNotExist(err) {
		fmt.Printf("Error: No note dated %s exists\n", fileName)
		os.Exit(1)
	}

	cmd := exec.Command("vim", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error: Running vim failed: %s\n", err)
		os.Exit(1)
	}
}
