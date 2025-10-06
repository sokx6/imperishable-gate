package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"imperishable-gate/internal/client/service"
	"imperishable-gate/internal/client/utils"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the server and store the refresh token securely",
	Run: func(cmd *cobra.Command, args []string) {

		reader := bufio.NewScanner(os.Stdin)

		// 读取用户名
		fmt.Print("Please enter your username: ")
		if !reader.Scan() {
			fmt.Fprintln(os.Stderr, "Failed to read username")
		}
		username := reader.Text()

		// 读取密码（明文显示）
		fmt.Print("Please enter your password: ")
		if !reader.Scan() {
			fmt.Fprintln(os.Stderr, "Failed to read password")
		}
		password := reader.Text()

		// 调用登录服务
		var err error
		var refreshToken string
		accessToken, refreshToken, err = service.Login(addr, username, password)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Login failed:", err)
		}
		// 存储 refresh token 到系统密钥链
		utils.SaveRefreshToken(refreshToken)
		fmt.Println("Login successful!")

	},
}

// 初始化命令行参数
func init() {
	rootCmd.AddCommand(loginCmd)
}
