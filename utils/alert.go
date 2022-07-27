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
