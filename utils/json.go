package utils

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/spf13/cobra"
)

func ReadJson(name string, v any) {
	data, err := os.ReadFile(name)

	// If the file doesn't exist, don't error as we just use the defaults
	if errors.Is(err, os.ErrNotExist) {
		return
	} else {
		cobra.CheckErr(err)
	}

	err = json.Unmarshal(data, &v)
	cobra.CheckErr(err)
}
