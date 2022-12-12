package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func searchNotes(searchInstance *search) {
	basePath := getBasePath()
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

		contentByte, err := os.ReadFile(path.Join(basePath, file.Name()))
		if err != nil {
			fmt.Printf("Warning: Reading the content of %s failed: %s\n", file.Name(), err)
			fmt.Print("Skipped file\n\n")
		}

		if searchInstance.caseSensitive {
			out += sensitiveSearch(file.Name(), string(contentByte), searchInstance)
		} else {
			out += insensitiveSearch(file.Name(), string(contentByte), searchInstance)
		}
	}

	if len(out) == 0 {
		os.Exit(1)
	}

	fmt.Println(strings.TrimSpace(out))
}

func sensitiveSearch(fileName string, content string, searchInstance *search) string {
	out := ""
	if strings.Contains(content, searchInstance.needle) {
		// Find out which line contains the needle
		for i, line := range strings.Split(content, "\n") {
			if strings.Contains(line, searchInstance.needle) {
				line = strings.ReplaceAll(line, searchInstance.needle, fmt.Sprintf("\033[0;31m%s\033[0m", searchInstance.needle))
				out += fmt.Sprintf("%s:%d: %s\n", fileName, i+1, line)
			}
		}
	}

	return out
}

func insensitiveSearch(fileName string, content string, searchInstance *search) string {
	out := ""
	if strings.Contains(strings.ToLower(content), searchInstance.needle) {
		// Find out which line contains the needle
		for i, line := range strings.Split(content, "\n") {
			lineLower := strings.ToLower(line)
			if strings.Contains(lineLower, searchInstance.needle) {
				index := strings.Index(lineLower, searchInstance.needle)
				toReplace := line[index : index+len(searchInstance.needle)]

				line = strings.ReplaceAll(line, toReplace, fmt.Sprintf("\033[0;31m%s\033[0m", toReplace))
				out += fmt.Sprintf("%s:%d: %s\n", fileName, i+1, line)
			}
		}
	}

	return out
}
