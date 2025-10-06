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
	RunE: func(cmd *cobra.Command, args []string) error {

		addr, _ = cmd.Flags().GetString("addr")

		reader := bufio.NewScanner(os.Stdin)

		// 读取用户名
		fmt.Print("Please enter your username: ")
		if !reader.Scan() {
			return fmt.Errorf("failed to read username")
		}
		username := reader.Text()

		// 读取密码（明文显示）
		fmt.Print("Please enter your password: ")
		if !reader.Scan() {
			return fmt.Errorf("failed to read password")
		}
		password := reader.Text()

		// 调用登录服务
		var err error
		var refreshToken string
		accessToken, refreshToken, err = service.Login(addr, username, password)
		if err != nil {
			return fmt.Errorf("login failed: %w", err)
		}
		// 存储 refresh token 到系统密钥链
		utils.SaveRefreshToken(refreshToken)
		refreshToken, err = utils.LoadRefreshToken()
		if err != nil {
			return fmt.Errorf("failed to save refresh token: %w", err)
		}
		fmt.Println("Refresh token saved to system keyring.")
		fmt.Println("Access Token:", accessToken)
		fmt.Println("Refresh Token:", refreshToken)
		fmt.Println("Login successful!")
		return nil
	},
}

// 初始化命令行参数
func init() {
	loginCmd.Flags().StringP("addr", "a", addr, "Server address (host:port)")
	rootCmd.AddCommand(loginCmd)
}
