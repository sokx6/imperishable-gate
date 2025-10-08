package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"imperishable-gate/internal/client/service"
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/response"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Display current authenticated user information",
	Long:  "Display information about the currently authenticated user, including user ID and username",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 尝试获取有效的 token（不提示用户登录）
		validToken, err := service.GetTokenAutomatically(addr, accessToken)

		// 如果无法获取有效 token，显示未登录状态
		if err != nil {
			fmt.Println("Not logged in")
			fmt.Println("Please run 'gate login' to authenticate")
			return nil
		}

		// 使用有效的 token 发送请求
		client := utils.NewAPIClient(addr, validToken)

		// 发送请求
		var result response.WhoamiResponse
		err = client.DoRequest("GET", "/api/v1/whoami", nil, &result)
		if err != nil {
			return fmt.Errorf("request failed: %w", err)
		}

		// 输出用户信息
		fmt.Printf("Authenticated as:\n")
		fmt.Printf("  User ID:  %d\n", result.UserInfo.UserID)
		fmt.Printf("  Username: %s\n", result.UserInfo.Username)

		// 可选：输出 JSON 格式（用于调试）
		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			respBody, _ := json.MarshalIndent(result, "", "  ")
			fmt.Printf("\nDetailed response:\n%s\n", respBody)
		}

		return nil
	},
}

func init() {
	// 添加 verbose 标志用于详细输出
	whoamiCmd.Flags().BoolP("verbose", "v", false, "Show detailed JSON response")
	rootCmd.AddCommand(whoamiCmd)
}
