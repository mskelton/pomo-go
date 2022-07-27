package config

import (
	"errors"
	"os"
	"path"

	"github.com/spf13/cobra"
)

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
