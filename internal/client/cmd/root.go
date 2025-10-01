package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gate",
	Short: "gate is a CLI link management tool",
	Long:  "gate is a CLI link management tool . . .",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to gate CLI tool. Use -h for help.")
	},
	Version: "1.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
