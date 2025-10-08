package cmd

import (
	"imperishable-gate/internal/client/service"

	"github.com/spf13/cobra"
)

var resetPasswordCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Reset your account password via email verification",
	RunE: func(cmd *cobra.Command, args []string) error {
		serverAddr, _ := cmd.Flags().GetString("addr")
		if serverAddr == "" {
			serverAddr = "127.0.0.1:8080"
		}
		return service.HandleResetPassword(serverAddr)
	},
}

func init() {
	rootCmd.AddCommand(resetPasswordCmd)
}
