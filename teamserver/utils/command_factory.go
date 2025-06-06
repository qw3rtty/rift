package utils

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Define dynamic commands with their entry functions
var Commands = map[string]func(cmd *cobra.Command, args []string) {
	"listener": func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting listener...")
	},
	"test": func(cmd *cobra.Command, args []string) {
		fmt.Println("TEST 1 2 3")
	},
	"help": func(cmd *cobra.Command, args []string) {
		fmt.Println("Help!")
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

// Check if given command is valid and exists
func IsValidCommand(name string) bool {
    for c, _ := range Commands {
        if c == name {
            return true
        }
    }
	return false
}

