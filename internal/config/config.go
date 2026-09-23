package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func Read() (Config, error) {
	cfgPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error: %v", err)
	}

	content, err := os.ReadFile(cfgPath)
	if err != nil {
		return Config{}, fmt.Errorf("error: %v", err)
	}

	var data Config
	if err = json.Unmarshal(content, &data); err != nil {
		return Config{}, fmt.Errorf("error: %v", err)
	}

	return data, nil
}

func (c Config) SetUser(usr string) (Config, error) {
	c.CurrentUsername = usr
	err := write(c)
	if err != nil {
		return Config{}, fmt.Errorf("error: %v", err)
	}

	return c, nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	cfgPath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	err = os.WriteFile(cfgPath, data, 0644)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	cfgPath := filepath.Join(homeDir, configFileName)

	return cfgPath, nil
}
