package utils

import (
	"time"

	"github.com/spf13/cobra"
)

func GetDuration(args []string, cfg string) time.Duration {
	if len(args) == 1 {
		duration, err := time.ParseDuration(args[0])
		cobra.CheckErr(err)
		return duration
	}

	duration, err := time.ParseDuration(cfg)
	cobra.CheckErr(err)
	return duration
}
