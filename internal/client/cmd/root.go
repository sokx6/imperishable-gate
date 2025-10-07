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
		// 对于 login, register, logout, help, version 命令，不需要预先验证 token
		if cmd.Use == "login" || cmd.Name() == "login" {
			return nil
		}
		if cmd.Use == "register" || cmd.Name() == "register" {
			return nil
		}
		if cmd.Use == "logout" || cmd.Name() == "logout" {
			return nil
		}
		if cmd.Use == "help" || cmd.Name() == "help" {
			return nil
		}
		if cmd.Use == "version" || cmd.Name() == "version" {
			return nil
		}
		// 获取服务器地址
		addr, _ = cmd.Flags().GetString("addr")
		if addr == "" {
			// 尝试从环境变量或 .env 文件中加载地址
			if err := godotenv.Load(); err == nil {
				addr = os.Getenv("SERVER_ADDR")
			} else {
				// 如果 .env 文件不存在，尝试直接从环境变量获取
				fmt.Println("Warning: SERVER_ADDR not set and .env file not found, using default 127.0.0.1:8080")
				addr = "127.0.0.1:8080"
			}
		}
		// 获取并验证访问令牌
		var err error
		accessToken, err = service.EnsureValidTokenWithPrompt(addr, accessToken)
		if err != nil {
			return err
		}
		return nil
	},
	// Run 执行根命令
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to gate CLI tool. Use -h for help.")
	},
	Version: "1.0.0",
}

func init() {
	// 定义全局持久标志addr（服务器地址）
	rootCmd.PersistentFlags().StringP("addr", "a", "", "Server address (host:port)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
