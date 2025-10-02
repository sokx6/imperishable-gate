package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Config struct {
	Addr string
}

var rootCmd = &cobra.Command{
	Use:   "gate",
	Short: "gate is a CLI link management tool",
	Long:  "gate is a CLI link management tool . . .",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		Config.Addr, _ = cmd.Flags().GetString("addr")
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to gate CLI tool. Use -h for help.")
	},
	Version: "1.0.0",
}

func init() {
	rootCmd.PersistentFlags().StringP("addr", "a", "127.0.0.1:8080", "Server address (host:port)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
