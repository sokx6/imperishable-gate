package cmd

import (
	"fmt"
	"imperishable-gate/internal/server"

	"imperishable-gate/internal/server/service"

	"github.com/spf13/cobra"
)

var Port string
var Dsn string

// StartCmd 是服务器的启动命令
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the gate server",
	Long:  `Starts the web server that listens for client requests.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 定义服务器地址
		address := ":" + Port
		fmt.Printf("Starting Imperishable Gate server on %s...\n", address)

		// 创建新的服务器实例
		srv := server.NewServer(address, Dsn)
		go service.ScheduledMetabaseFetch()

		// 启动服务器
		if err := srv.Start(); err != nil {
			fmt.Println("Server failed to start:", err)
			return err
		}

		return nil
	},
}

func init() {
	StartCmd.Flags().StringVarP(&Port, "port", "p", "1270", "Port to listen on (default: 1270)")
	StartCmd.Flags().StringVarP(&Dsn, "dsn", "d", "", "Data source name for the database")
}
