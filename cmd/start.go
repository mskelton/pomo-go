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
		cfg := config.GetConfig()
		status := config.Status{
			Type: config.TYPE_FOCUS,
			End:  time.Now().Add(utils.GetDuration(args, cfg.Durations.Focus)),
		}

		config.WriteStatus(status)
		utils.CmdNotify(cmd, func() {
			cfg := config.GetConfig()
			utils.Alert(cfg.Emojis.Focus, "Your focus session has started!", cfg.Sound)
		})
	},
}

func init() {
	startCmd.Flags().Bool("notify", false, "display a push notification")
	rootCmd.AddCommand(startCmd)
}
