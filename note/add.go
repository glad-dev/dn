// Package note adds, edits, removes, searches and displays notes.
package note

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/GLAD-DEV/dn/config"
)

func Add(content string, date string) error {
	path, err := config.BasePath()
	if err != nil {
		return err
	}

	fullPath := filepath.Join(path, date)

	//nolint:gosec
	// date is a formated date or empty.
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o600)
	if err != nil {
		return fmt.Errorf("failed to open note: %w", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = fmt.Fprintf(file, "* %s\n", content)
	if err != nil {
		return fmt.Errorf("failed to write note: %w", err)
	}

	return nil
}
