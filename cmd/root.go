package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/mskelton/pomo/config"
	"github.com/mskelton/pomo/utils"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "pomo",
	Short: "Simple pomodoro timer",
	Run: func(cmd *cobra.Command, args []string) {
		status := config.ReadStatus()

		// Don't print anything if there is no active session
		if status.End.IsZero() {
			return
		}

		// Print the remaining time
		remaining := status.End.Sub(time.Now()).Round(time.Second)
		fmt.Printf("%s %s\n", getEmoji(status, remaining), remaining)

		// Notify the user when the remaining time has elapsed
		if !status.Notified && remaining.Seconds() <= 0 {
			if status.Type == config.TYPE_FOCUS {
				utils.Alert("🥂", "Focus completed, let's take a break!", "Glass")
			} else {
				utils.Alert("🍅", "Break is over, back to work!", "Glass")
			}

			// Update the status to indicate the notification has been queued to
			// prevent duplicate notifications.
			status.Notified = true
			config.WriteStatus(status)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_CONFIG_HOME/pomo/config.json)")
}

func getEmoji(status config.Status, remaining time.Duration) string {
	// Blink the emoji when the pomodoro has finished
	if remaining.Seconds() <= 0 {
		if int(remaining.Seconds())%2 == 0 {
			return "️🔴"
		} else {
			return "⭕"
		}
	}

	if status.Type == config.TYPE_FOCUS {
		return "🍅"
	} else {
		return "🥂"
	}
}
