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
			// 优先使用 GATE_SERVER_ADDR，兼容 SERVER_ADDR
			addr = os.Getenv("GATE_SERVER_ADDR")
			if addr == "" {
				addr = os.Getenv("SERVER_ADDR")
			}
		}

		if addr == "" {
			// 如果不能加载地址，使用默认地址
			fmt.Println("Warning: GATE_SERVER_ADDR not set and .env file not found, using default http://localhost:4514")
			addr = "http://localhost:4514"
		}
	}
	return addr
}
