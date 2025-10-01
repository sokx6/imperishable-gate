package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gate-server",
	Short: "Imperishable Gate Server",
	Long:  `A simple server that responds to ping requests.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(StartCmd)
}
