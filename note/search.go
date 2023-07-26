package note

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GLAD-DEV/dn/config"
)

func Search(needle string, caseSensitive bool) error {
	path, err := config.BasePath()
	if err != nil {
		return err
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	out := ""
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("Warning: Encountered a directory: %s\n", file.Name())

			continue
		}

		contentByte, err := os.ReadFile(filepath.Join(path, file.Name()))
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
		return errors.New("no match found")
	}

	fmt.Println(strings.TrimSpace(out))

	return nil
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
