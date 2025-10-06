package cmd

import (
	"fmt"
	"os"

	"imperishable-gate/internal/client/service"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var accessToken string = ""
var addr string

var rootCmd = &cobra.Command{
	Use:   "gate",
	Short: "gate is a CLI link management tool",
	Long:  "gate is a CLI link management tool . . .",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		addr, _ = cmd.Flags().GetString("addr")
		if addr == "" {
			if err := godotenv.Load(); err == nil {
				addr = os.Getenv("SERVER_ADDR")
			} else {
				fmt.Println("Warning: SERVER_ADDR not set and .env file not found, using default 127.0.0.1:8080")
				addr = "127.0.0.1:8080"
			}
		}
		var err error
		accessToken, err = service.EnsureValidTokenWithPrompt(addr, accessToken)
		if err != nil {
			return err
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to gate CLI tool. Use -h for help.")
	},
	Version: "1.0.0",
}

func init() {
	rootCmd.PersistentFlags().StringP("addr", "a", "", "Server address (host:port)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
