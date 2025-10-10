package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gate-server",
	Short:   "Imperishable Gate Server",
	Long:    `Imperishable Gate Server - A link management service with authentication and tagging support.`,
	Version: "1.0.1",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(StartCmd)
}
