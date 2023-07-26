package util

import (
	"fmt"
	"os"
)

func PrintErrAndExit(str string) {
	_, err := fmt.Fprintf(os.Stderr, "Error: %s\n", str)
	if err != nil {
		// Failed to write to stderr => write to stdout
		fmt.Printf("Error: %s\n", err)
	}

	os.Exit(1)
}
