package cmd

import (
	"fmt"
	"imperishable-gate/internal/server"

	"github.com/spf13/cobra"
)

var Port string

// StartCmd 是服务器的启动命令
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the gate server",
	Long:  `Starts the web server that listens for client requests.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		address := ":" + Port
		fmt.Printf("Starting Imperishable Gate server on %s...\n", address)

		srv := server.NewServer(address)
		if err := srv.Start(); err != nil {
			fmt.Println("Server failed to start:", err)
			return err
		}

		return nil
	},
}

func init() {
	StartCmd.Flags().StringVarP(&Port, "port", "p", "1270", "Port to listen on (default: 1270)")
}
