package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from the server and clear stored tokens",
	Long:  "Logout from the server, invalidate the refresh token, and clear tokens from system keyring",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载存储的刷新令牌
		refreshToken, err := utils.LoadRefreshToken()
		if err != nil {
			fmt.Println("No active session found or failed to load refresh token")
			// 即使没有找到token，也尝试清理本地存储
			utils.ClearTokens()
			return nil
		}

		// 创建 API 客户端（logout不需要access token）
		client := utils.NewAPIClient(addr, accessToken)

		// 构造请求体
		reqBody := request.LogoutRequest{
			RefreshToken: refreshToken,
		}

		// 发送登出请求到服务器
		var result response.Response
		err = client.DoRequest("POST", "/api/v1/logout", reqBody, &result)
		if err != nil {
			// 即使服务器请求失败，也清理本地tokens
			fmt.Printf("Warning: Failed to notify server: %v\n", err)
		} else {
			fmt.Println(result.Message)
		}

		// 清理本地存储的tokens
		utils.ClearTokens()
		fmt.Println("Tokens cleared from system keyring.")
		fmt.Println("Logout successful!")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
