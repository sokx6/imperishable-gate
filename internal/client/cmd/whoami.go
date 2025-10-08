package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/response"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Display current authenticated user information",
	Long:  "Display information about the currently authenticated user, including user ID and username",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 使用全局的 accessToken 和 addr（在 PersistentPreRunE 中已验证）
		client := utils.NewAPIClient(addr, accessToken)

		// 发送请求
		var result response.WhoamiResponse
		err := client.DoRequest("GET", "/api/v1/whoami", nil, &result)
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
