package cmd

import (
	"fmt"
	"net/http"
	"teamserver/utils"
	"github.com/spf13/cobra"
)

var port int

var listenerCmd = &cobra.Command{
	Use:   "listener",
	Short: "Starting listener...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[i] Starting listener ...")
		handler := http.NewServeMux()
		handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "[i] You have hit %s\n", r.URL.Path)
		})

		hl := utils.NewHTTPListener(":4444", handler)
		utils.HLManager.Register(hl)

		if err := hl.Start(); err != nil {
			fmt.Printf("[!] Failed to start listener: %v\n", err)
		} else {
			fmt.Printf("[+] Listener startet on: %s\n", hl.Addr)
		}

		hl.WaitForShutdown()

		if err := hl.Shutdown(); err != nil {
			fmt.Printf("[!] Failed to gracefully shutdown: %v\n", err)
		}

		utils.HLManager.Unregister(0)
		fmt.Println("[i] Server gracefully stopped")
	},
}

func init() {
	listenerCmd.Flags().IntVarP(&port, "port", "p", 4444, "Port for the listener")
	riftCmd.AddCommand(listenerCmd)
}
