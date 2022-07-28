package utils

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func Alert(emoji string, message string, sound string) {
	osa, err := exec.LookPath("osascript")
	cobra.CheckErr(err)

	cmd := exec.Command(osa, "-e", fmt.Sprintf(
		`tell application "System Events" to display notification "%s" with title "Pomo %s" sound name "%s"`,
		message,
		emoji,
		sound,
	))

	err = cmd.Run()
	cobra.CheckErr(err)
}

// Notify the user if they requested it. This is useful in context such as
// tmux where you want to know if you accidentally start/stop a session.
func CmdNotify(cmd *cobra.Command, cb func()) {
	notify, err := cmd.Flags().GetBool("notify")
	cobra.CheckErr(err)

	if notify {
		cb()
	}
}
