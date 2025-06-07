package cli

import (
	"fmt"
	"time"
	"teamserver/server"
	"github.com/spf13/cobra"
)

var port int

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starting the protal server...",
	Run: func(cmd *cobra.Command, args []string) {

		asciiArt := `
		RRR   III  FFFFF  TTTTT
		R  R   I   F        T
		RRR    I   FFFF     T
		R  R   I   F        T
		R  R  III  F        T
		`
		// Ausgabe des ASCII-Art
		fmt.Println(asciiArt)


		fmt.Printf("[+] Starting the portal server on port %d (WebSocket)\n", port)

		// Starte Logging-Goroutine für aktive Agents
		go func() {
			for {
				time.Sleep(10 * time.Second)
				clients := server.ClManager.List()
				fmt.Printf("[#] Active clients: %d\n", len(clients))
				for _, c := range clients {
					fmt.Printf("    ID: %s | IP: %s | LastSeen: %s\n", c.ID, c.IP, c.LastSeen.Format(time.RFC3339))
				}
			}
		}()

		server.StartWebSocketServer(port)
	},
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port of the portal server")
	riftCli.AddCommand(serverCmd)
}
