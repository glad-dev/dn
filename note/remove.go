package note

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/GLAD-DEV/dn/config"
)

func Remove(date string) error {
	path, err := config.BasePath()
	if err != nil {
		return err
	}

	path = filepath.Join(path, date)
	// Check if file exists
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("there is no note for the date %s", date)
	}

	if info.IsDir() {
		return fmt.Errorf("the passed path ('%s') is a directory, not a file", path)
	}

	err = os.Remove(path)
	if err != nil {
		return fmt.Errorf("failed to remove file: %w", err)
	}

	return nil
}
