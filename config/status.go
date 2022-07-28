package config

import (
	"encoding/json"
	"path"
	"time"

	"github.com/mskelton/pomo/utils"
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
	status := Status{}
	utils.ReadJson(getStatusPath(), &status)

	return status
}
