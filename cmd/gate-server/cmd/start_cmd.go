package cmd

import (
	"fmt"
	"imperishable-gate/internal/server"
	"os"

	"imperishable-gate/internal/server/service"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// StartCmd 是服务器的启动命令
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the gate server",
	Long:  `Starts the web server that listens for client requests.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 定义服务器地址
		dsn, _ := cmd.Flags().GetString("dsn")
		addr, _ := cmd.Flags().GetString("addr")
		if err := godotenv.Load(); err != nil {
			fmt.Println("No .env file found, now using default values.")
		}
		if addr == "" {
			addr = os.Getenv("SERVER_ADDR")
		}
		if dsn == "" {
			dsn = os.Getenv("DSN")
		}
		fmt.Printf("Starting Imperishable Gate server on %s...\n", addr)

		// 创建新的服务器实例
		srv := server.NewServer(addr, dsn)
		go service.ScheduledNotWatchingMetabaseFetch()
		go service.ScheduledWatchingMetabaseFetch()

		// 启动服务器
		if err := srv.Start(); err != nil {
			fmt.Println("Server failed to start:", err)
			return err
		}

		return nil
	},
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, now using default values.")
	}
	StartCmd.Flags().StringP("addr", "a", "", "Address to listen on (default: localhost:4514)")
	StartCmd.Flags().StringP("dsn", "d", "", "Data source name for the database")

}
