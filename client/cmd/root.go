package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var serverURL string

var rootCmd = &cobra.Command{
    Use:   "rift",
    Short: "rift client with tui",
}

func Execute() {
    rootCmd.PersistentFlags().StringVar(&serverURL, "server", "ws://localhost:8080/client", "WebSocket Teamserver URL")
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
