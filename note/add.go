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
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o600)
	if err != nil {
		return fmt.Errorf("failed to open note: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("* %s\n", content))
	if err != nil {
		return fmt.Errorf("failed to write note: %w", err)
	}

	return nil
}
