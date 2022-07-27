package cmd

import (
	"github.com/mskelton/pomo/config"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current session",
	Run: func(cmd *cobra.Command, args []string) {
		config.WriteStatus(config.Status{})
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
