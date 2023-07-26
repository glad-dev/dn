package note

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/GLAD-DEV/dn/config"
)

func Edit(fileName string) error {
	path, err := config.BasePath()
	if err != nil {
		return err
	}

	filePath := filepath.Join(path, fileName)

	// Check if file exists
	f, err := os.Open(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("no note for the date '%s' exists", fileName)
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
		return fmt.Errorf("failed to run the editor: %w", err)
	}

	// Delete note if it contains no content
	info, err := os.Stat(filePath)
	if err != nil {
		// Since this is an optional step, we silently ignore the error.
		return nil // nolint:nilerr
	}

	if info.Size() == 0 {
		return Remove(fileName)
	}

	return nil
}
