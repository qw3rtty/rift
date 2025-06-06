package cmd

import (
	"os"
	"teamserver/utils"

	"github.com/spf13/cobra"
)

var riftCli = &cobra.Command{
	Use:   "rift",
	Short: "Like a Portal Gun.",
	Long:  `Like a Portal Gun, But for Red Teams. Tear a Hole in Reality. And Beacon Through It.`,
}

func Execute() {

	// Register dynamic commands
	for cmd, fn := range utils.Commands {
		c := utils.CommandFactory(cmd, fn)
		riftCli.AddCommand(c)
	}

	if err := riftCli.Execute(); err != nil {
		os.Exit(1)
	}
}
