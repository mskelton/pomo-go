package cmd

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/mskelton/pomo/config"
	"github.com/mskelton/pomo/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pomo",
	Short: "Simple pomodoro timer",
	Run: func(cmd *cobra.Command, args []string) {
		status := config.ReadStatus()

		// Don't print anything if there is no active session
		if status.End.IsZero() {
			return
		}

		// Save IO ops by reading the config only if there is a running session
		cfg := config.GetConfig()

		// Print the remaining time
		remaining := status.End.Sub(time.Now()).Round(time.Second)
		fmt.Printf("%s %s\n", getEmoji(cfg, status, remaining), remaining)

		// Notify the user when the remaining time has elapsed
		if !status.Notified && remaining.Seconds() <= 0 {
			if status.Type == config.TYPE_FOCUS {
				utils.Alert(cfg.Emojis.Break, "Focus completed, let's take a break!", cfg.Sound)
			} else {
				utils.Alert(cfg.Emojis.Focus, "Break is over, back to work!", cfg.Sound)
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

func getEmoji(cfg config.Config, status config.Status, remaining time.Duration) string {
	// Cycle through the warning emojis to make the timer "blink" when used in a
	// statusline. This can be disabled by overriding the configuration to only
	// provide a single emoji.
	if remaining.Seconds() <= 0 {
		index := int(math.Abs(remaining.Seconds())) % len(cfg.Emojis.Warn)
		return cfg.Emojis.Warn[index]
	}

	if status.Type == config.TYPE_FOCUS {
		return cfg.Emojis.Focus
	} else {
		return cfg.Emojis.Break
	}
}
