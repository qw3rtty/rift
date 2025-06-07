package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)


var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Prints the help infos...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[i] HELP!")
	},
}

func init() {
	riftCmd.AddCommand(helpCmd)
}
