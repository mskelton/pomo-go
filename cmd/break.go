package cmd

import (
	"time"

	"github.com/mskelton/pomo/config"
	"github.com/mskelton/pomo/utils"
	"github.com/spf13/cobra"
)

var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "Start a break",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()

		status := config.Status{
			Type: config.TYPE_BREAK,
			End:  time.Now().Add(utils.GetDuration(args, cfg.Durations.Break)),
		}

		config.WriteStatus(status)
	},
}

func init() {
	rootCmd.AddCommand(breakCmd)
}
