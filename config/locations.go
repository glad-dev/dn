package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateDir() error {
	path, err := BasePath()
	if err != nil {
		return err
	}

	// Creates the note directory if it does not yet exist
	err = os.Mkdir(path, 0o755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create note directory at '%s': %w", path, err)
	}

	return nil
}

func BasePath() (string, error) {
	homeDir, ok := os.LookupEnv("DN_HOME")
	if ok && len(homeDir) > 0 {
		// DN_HOME is set
		// Check if path is absolute
		if !filepath.IsAbs(homeDir) {
			return "", fmt.Errorf("$DN_HOME (%s) is not an absoulte path", homeDir)
		}

		// Check if path is valid
		_, err := os.Stat(homeDir)
		if err == nil || os.IsNotExist(err) {
			return homeDir, nil
		}

		// os.Stat failed with an unknown error
		return "", fmt.Errorf("$DN_HOME ('%s') is not a valid path: %w", homeDir, err)
	}

	// No custom home directory is set => default to $HOME/dn/
	var err error
	homeDir, err = os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to resolve home directory: %w", err)
	}

	return filepath.Join(homeDir, "dn"), nil
}

func configPath() (string, error) {
	base, err := BasePath()
	if err != nil {
		return "", err
	}

	return filepath.Join(base, "config.toml"), nil
}
