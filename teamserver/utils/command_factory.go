package utils

import (
	"fmt"
	"net/http"
	"github.com/spf13/cobra"
)

// Define dynamic commands with their entry functions
var Commands = map[string]func(cmd *cobra.Command, args []string) {
	"listener": func(cmd *cobra.Command, args []string) { 
		// TODO: Add flags to set settings
		fmt.Println("[i] Starting listener ...")
		handler := http.NewServeMux()
		handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "[i] You have hit %s\n", r.URL.Path)
		})

		hl := NewHTTPListener(":4444", handler)

		if err := hl.Start(); err != nil {
			fmt.Printf("[!] Failed to start listener: %v\n", err)
		} else {
			fmt.Printf("[+] Listener startet on: %s\n", hl.Addr)
		}

		hl.WaitForShutdown()

		if err := hl.Shutdown(); err != nil {
			fmt.Printf("[!] Failed to gracefully shutdown: %v\n", err)
		}

		fmt.Println("[i] Server gracefully stopped")
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

