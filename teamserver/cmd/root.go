package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var riftCli = &cobra.Command{
	Use:   "rift",
	Short: "Like a Portal Gun.",
	Long:  `Like a Portal Gun, But for Red Teams. Tear a Hole in Reality. And Beacon Through It.`,
}

func Execute() {
	if err := riftCli.Execute(); err != nil {
		os.Exit(1)
	}
}
