package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var riftCmd = &cobra.Command{
	Use:   "rift-cmd",
	Short: "Teamserver commands.",
	Long:  `Teamserver commands.`,
}

func isValidCommand(name string) bool {
	cmd, _, err := riftCmd.Find([]string{name})
	if err != nil || cmd == nil || cmd == riftCmd {
		fmt.Println("[!] Whoops! Invalid command!")
		return false
	}
	return true
}

func ExecuteCommand(name string) {
	if !isValidCommand(name) {
		name = "help"
	}

	cmd, _, _ := riftCmd.Find([]string{name})
	//cmd.SetArgs([]string{name})
	cmd.Run(cmd, []string{})
}

func Execute() {
	if err := riftCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
