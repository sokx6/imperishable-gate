package cmd

import (
	"github.com/spf13/cobra"

	"imperishable-gate/internal/client/service/register"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user account",
	Long:  "Register a new user account with username, email, and password",
	RunE: func(cmd *cobra.Command, args []string) error {

		if addr == "" {
			addr = "localhost:4514" // 默认地址
		}

		// 使用模块化的注册处理函数
		if err := register.HandleRegister(addr); err != nil {
			return err
		}

		return nil
	},
}

// 初始化命令行参数
func init() {
	rootCmd.AddCommand(registerCmd)
}
