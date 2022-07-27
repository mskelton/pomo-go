package config

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"time"

	"github.com/spf13/cobra"
)

const (
	TYPE_BREAK = iota
	TYPE_FOCUS
)

type Status struct {
	Type     int       `json:"type"`
	End      time.Time `json:"end"`
	Notified bool      `json:"notified"`
}

func getStatusPath() string {
	return path.Join(GetConfigDir(), "status.json")
}

func WriteStatus(status Status) {
	data, err := json.Marshal(status)
	cobra.CheckErr(err)

	WriteFile(getStatusPath(), data)
}

func ReadStatus() Status {
	data, err := os.ReadFile(getStatusPath())

	// If the status file doesn't exist, don't throw an error since it just means
	// the user hasn't started a session yet which is not an error.
	if errors.Is(err, os.ErrNotExist) {
		return Status{}
	} else {
		cobra.CheckErr(err)
	}

	status := Status{}
	err = json.Unmarshal(data, &status)
	cobra.CheckErr(err)

	return status
}
