package cmd

import (
	"errors"
	"time"

	"github.com/mskelton/pomo/config"
	"github.com/spf13/cobra"
)

var durationCmd = &cobra.Command{
	Use:   "duration",
	Short: "Change the duration of the current session",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		status := config.ReadStatus()

		// Don't allow editing the duration if no session is in progress
		if time.Now().After(status.End) {
			cobra.CheckErr(errors.New("no session in progress"))
		}

		duration, err := time.ParseDuration(args[0])
		cobra.CheckErr(err)

		status.End = time.Now().Add(duration)
		config.WriteStatus(status)
	},
}

func init() {
	rootCmd.AddCommand(durationCmd)
}
