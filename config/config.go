package config

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"time"

	"github.com/spf13/cobra"
)

type Durations struct {
	Break time.Duration `json:"break"`
	Focus time.Duration `json:"focus"`
}

type Emojis struct {
	Break string `json:"break"`
	Focus string `json:"focus"`
}

type Config struct {
	Durations Durations `json:"durations"`
	Emojis    Emojis    `json:"emojis"`
	Sound     string    `json:"sound"`
}

func GetConfigDir() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return path.Join(home, ".config", "pomo")
}

func WriteFile(name string, data []byte) {
	dir := GetConfigDir()

	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		cobra.CheckErr(err)
	}

	err := os.WriteFile(name, data, 0644)
	cobra.CheckErr(err)
}

func GetConfig() Config {
	data, err := os.ReadFile(path.Join(GetConfigDir(), "config.json"))
	cfg := Config{
		Durations: Durations{Break: time.Minute * 5, Focus: time.Minute * 30},
		Emojis:    Emojis{Break: "🥂", Focus: "🍅"},
		Sound:     "default",
	}

	// If the config file doesn't exist, just return the defaults
	if errors.Is(err, os.ErrNotExist) {
		return cfg
	} else {
		cobra.CheckErr(err)
	}

	err = json.Unmarshal(data, &cfg)
	cobra.CheckErr(err)

	return cfg
}
