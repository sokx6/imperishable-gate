package core

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// LoadServerAddr 从命令行参数、环境变量或 .env 文件中加载服务器地址
func LoadServerAddr(cmd *cobra.Command) string {
	addr, _ := cmd.Flags().GetString("addr")
	if addr == "" {
		// 尝试从环境变量或 .env 文件中加载地址
		if err := godotenv.Load(); err == nil {
			addr = os.Getenv("SERVER_ADDR")
		}

		if addr == "" {
			// 如果不能加载地址，使用默认地址
			fmt.Println("Warning: SERVER_ADDR not set and .env file not found, using default localhost:4514")
			addr = "localhost:4514"
		}
	}
	return addr
}
