package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func searchNotes(needle string, caseSensitive bool) int {
	base := getBasePath()

	files, err := os.ReadDir(base)
	if err != nil {
		fmt.Printf("Error: Reading the base directory failed: %s\n", err)

		return exitFailure
	}

	out := ""
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("Warning: Encountered a directory: %s\n", file.Name())

			continue
		}

		contentByte, err := os.ReadFile(filepath.Join(base, file.Name()))
		if err != nil {
			fmt.Printf("Warning: Reading the content of %s failed: %s\n", file.Name(), err)
			fmt.Print("Skipped file\n\n")
		}

		if caseSensitive {
			out += sensitiveSearch(file.Name(), string(contentByte), needle)
		} else {
			out += insensitiveSearch(file.Name(), string(contentByte), needle)
		}
	}

	if len(out) == 0 {
		return exitFailure
	}

	fmt.Println(strings.TrimSpace(out))

	return exitSuccess
}

func sensitiveSearch(fileName string, content string, needle string) string {
	out := ""
	if strings.Contains(content, needle) {
		// Find out which line contains the needle
		for i, line := range strings.Split(content, "\n") {
			if strings.Contains(line, needle) {
				line = strings.ReplaceAll(line, needle, fmt.Sprintf("\033[0;31m%s\033[0m", needle))
				out += fmt.Sprintf("%s:%d: %s\n", fileName, i+1, line)
			}
		}
	}

	return out
}

func insensitiveSearch(fileName string, content string, needle string) string {
	out := ""
	if strings.Contains(strings.ToLower(content), needle) {
		// Find out which line contains the needle
		for i, line := range strings.Split(content, "\n") {
			lineLower := strings.ToLower(line)
			if strings.Contains(lineLower, needle) {
				index := strings.Index(lineLower, needle)
				toReplace := line[index : index+len(needle)]

				line = strings.ReplaceAll(line, toReplace, fmt.Sprintf("\033[0;31m%s\033[0m", toReplace))
				out += fmt.Sprintf("%s:%d: %s\n", fileName, i+1, line)
			}
		}
	}

	return out
}
