package utils

import (
	"fmt"
	"net/http"
	"github.com/spf13/cobra"
)

// Define dynamic commands with their entry functions
var Commands = map[string]func(cmd *cobra.Command, args []string) {
	"listener": func(cmd *cobra.Command, args []string) { // TODO: Add flags to set settings
		fmt.Println("[+] Starting listener...")
		
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "[i] You have hit %s\n", r.URL.Path)
		})

		addr := "localhost:4444" 
		fmt.Printf("[+] Starting HTTP server on %s...\n", addr)
		if err := http.ListenAndServe(addr, nil); err != nil {
			fmt.Printf("[!] Failed to start server: %v\n", err)
		}
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

