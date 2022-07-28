package config

import (
	"errors"
	"os"
	"path"

	"github.com/mskelton/pomo/utils"
	"github.com/spf13/cobra"
)

type durations struct {
	Break string `json:"break"`
	Focus string `json:"focus"`
}

type emojis struct {
	Break string   `json:"break"`
	Focus string   `json:"focus"`
	Warn  []string `json:"warn"`
}

type Config struct {
	Durations durations `json:"durations"`
	Emojis    emojis    `json:"emojis"`
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
	cfg := Config{
		Durations: durations{Break: "5m", Focus: "30m"},
		Emojis:    emojis{Break: "🥂", Focus: "🍅", Warn: []string{"🔴", "⭕"}},
		Sound:     "default",
	}

	utils.ReadJson(path.Join(GetConfigDir(), "config.json"), &cfg)
	return cfg
}
