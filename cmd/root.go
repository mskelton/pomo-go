package cmd

import (
	"errors"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/mskelton/pomo/config"
	"github.com/mskelton/pomo/utils"
	"github.com/spf13/cobra"
)

var noEmoji = false
var format string

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
		formatted, err := formatTime(format, remaining)
		cobra.CheckErr(err)

		if noEmoji {
			fmt.Println(formatted)
		} else {
			fmt.Printf("%s %s\n", getEmoji(cfg, status, remaining), formatted)
		}

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

func init() {
	rootCmd.Flags().BoolVar(&noEmoji, "no-emoji", false, "disable emojis in the status")
	rootCmd.Flags().StringVar(&format, "format", "duration", "format style for the remaining time")
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

func fmtDuration(duration time.Duration) string {
	absolute := duration.Abs()
	sign := ""

	// Indicate if the time has fully expired
	if math.Signbit(duration.Seconds()) {
		sign = "-"
	}

	// Get the components of the duration
	h := absolute / time.Hour
	absolute -= h * time.Hour
	m := absolute / time.Minute
	absolute -= m * time.Minute
	s := absolute / time.Second

	// Don't display seconds for durations greater than an hour
	if h >= 1 {
		return fmt.Sprintf("%s%dh%02dm", sign, h, m)
	} else if m >= 1 {
		return fmt.Sprintf("%s%dm%02ds", sign, m, s)
	} else {
		return fmt.Sprintf("%s%ds", sign, s)
	}
}

func formatTime(format string, duration time.Duration) (string, error) {
	if format == "duration" {
		return fmtDuration(duration), nil
	}

	if format == "time" {
		// Display the format either as hh:mm:ss or mm:ss
		layout := "15:04:05"
		if duration.Hours() < 1 {
			layout = "04:05"
		}

		// Format the time taking into account negative durations
		base := time.Unix(0, 0).UTC()
		if duration.Seconds() >= 0 {
			return base.Add(duration).Format(layout), nil
		} else {
			return "-" + base.Add(duration.Abs()).Format(layout), nil
		}
	}

	return "", errors.New("invalid format option")
}
