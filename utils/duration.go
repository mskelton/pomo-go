package utils

import (
	"time"

	"github.com/spf13/cobra"
)

func GetDuration(args []string, fallback time.Duration) time.Duration {
	if len(args) == 1 {
		duration, err := time.ParseDuration(args[0])
		cobra.CheckErr(err)
		return duration
	}

	return fallback
}
