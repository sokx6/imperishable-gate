package core

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// LoadServerAddr 从命令行参数、环境变量或 .env 文件中加载服务器地址
// 优先级: 命令行参数 > 系统环境变量 > .env 文件 > 默认值
func LoadServerAddr(cmd *cobra.Command) string {
	addr, _ := cmd.Flags().GetString("addr")
	if addr == "" {
		// 先尝试从系统环境变量中获取
		addr = os.Getenv("GATE_SERVER_ADDR")
		if addr == "" {
			addr = os.Getenv("SERVER_ADDR")
		}

		// 如果系统环境变量中没有,再尝试从 .env 文件加载
		if addr == "" {
			_ = godotenv.Load()
			addr = os.Getenv("GATE_SERVER_ADDR")
			if addr == "" {
				addr = os.Getenv("SERVER_ADDR")
			}
		}

		if addr == "" {
			// 如果都没有,使用默认地址
			fmt.Println("Warning: GATE_SERVER_ADDR not set and .env file not found, using default http://localhost:4514")
			addr = "http://localhost:4514"
		}
	}
	return addr
}
