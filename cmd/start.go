package cmd

import (
	"time"

	"github.com/mskelton/pomo/config"
	"github.com/mskelton/pomo/utils"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start focusing",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		status := config.Status{
			End:     time.Now().Add(utils.GetDuration(args, time.Minute*25)),
			IsFocus: true,
		}

		config.WriteStatus(status)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
