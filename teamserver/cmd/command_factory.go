package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Define dynamic commands with their entry functions
var commands = map[string]func(cmd *cobra.Command, args []string) {
	"listener": func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting listener...")
	},
	"test": func(cmd *cobra.Command, args []string) {
		fmt.Println("TEST 1 2 3")
	},
}

// Factory, which creates a new dynamic command
func CommandFactory(command string, runFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	return &cobra.Command{
		Use:	command,
		Short:	"Execute " + command + " command",
		Run:	runFunc,
	}
}

