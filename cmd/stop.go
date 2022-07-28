package cmd

import (
	"github.com/mskelton/pomo/config"
	"github.com/mskelton/pomo/utils"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current pomodoro session",
	Run: func(cmd *cobra.Command, args []string) {
		config.WriteStatus(config.Status{})
		utils.CmdNotify(cmd, func() {
			cfg := config.GetConfig()
			utils.Alert(cfg.Emojis.Focus, "Your session has stopped!", cfg.Sound)
		})
	},
}

func init() {
	stopCmd.Flags().Bool("notify", false, "display a push notification")
	rootCmd.AddCommand(stopCmd)
}
