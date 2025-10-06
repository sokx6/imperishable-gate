package cmd

import (
	"fmt"
	"imperishable-gate/internal/server"

	"imperishable-gate/internal/server/service"

	"os"

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
		Host := cmd.Flag("host").Value.String()
		Port := cmd.Flag("port").Value.String()
		Dsn := cmd.Flag("dsn").Value.String()
		address := Host + ":" + Port
		fmt.Printf("Starting Imperishable Gate server on %s...\n", address)

		// 创建新的服务器实例
		srv := server.NewServer(address, Dsn)
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
	Host := os.Getenv("SERVER_HOST")
	Port := os.Getenv("SERVER_PORT")
	Dsn := os.Getenv("DATABASE_URL")
	StartCmd.Flags().StringVarP(&Port, "port", "p", "1270", "Port to listen on (default: 1270)")
	StartCmd.Flags().StringVarP(&Dsn, "dsn", "d", "", "Data source name for the database")
	StartCmd.Flags().StringVarP(&Host, "host", "H", "localhost", "Host to listen on (default: localhost)")
}
