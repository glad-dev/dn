package config

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	SilentAdd    bool
	SilentRemove bool
}

func Load() (*Config, error) {
	config := Config{}

	configLocation, err := configPath()
	if err != nil {
		return nil, err
	}

	_, err = toml.DecodeFile(configLocation, &config)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &Config{
				SilentAdd:    false,
				SilentRemove: false,
			}, nil
		}
	}

	return &config, nil
}

func Write(config *Config) error {
	configLocation, err := configPath()
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = toml.NewEncoder(buf).Encode(config)
	if err != nil {
		return fmt.Errorf("failed to encode the config: %w", err)
	}

	// Delete the old config
	f, err := os.OpenFile(configLocation, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o600)
	if err != nil {
		return fmt.Errorf("failed to open config: %w", err)
	}
	defer f.Close()

	_, err = f.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("Failed to write config: %w", err)
	}

	return nil
}

func Print(config *Config) {
	fmt.Printf("silentAdd: %v\n", config.SilentAdd)
	fmt.Printf("silentRemove: %v\n", config.SilentRemove)
}
