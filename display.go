package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func displayNotes(slug string, basePath string) {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Printf("Error: Reading the base directory failed: %s\n", err)
		os.Exit(1)
	}

	out := ""
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("Warning: Encountered a directory: %s\n", file.Name())
			continue
		}

		if !strings.HasPrefix(file.Name(), slug) {
			// All files have a prefix of ""
			continue
		}

		content, err := os.ReadFile(filepath.Join(basePath, file.Name()))
		if err != nil {
			fmt.Printf("Warning: Reading the content of %s failed: %s\n", file.Name(), err)
			fmt.Print("Skipped file\n\n")
		}

		// Add space before every bullet point
		split := strings.Split(string(content), "\n")
		for i, str := range split {
			split[i] = " " + str
		}

		out += strings.TrimSpace(fmt.Sprintf("%s\n%s", file.Name(), strings.Join(split, "\n"))) + "\n"
	}

	fmt.Print(strings.TrimSpace(out))
}
