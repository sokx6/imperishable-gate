package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Send a ping request to the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		msg, _ := cmd.Flags().GetString("message")

		// 创建 API 客户端
		client := utils.NewAPIClient(addr, "")

		// 构造请求体
		reqBody := request.PingRequest{
			Action:  "ping",
			Message: msg,
		}

		// 将请求体编码为 JSON (用于日志输出)
		body, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}

		fmt.Printf("-- Requesting POST method to %s/api/v1/ping with payload\n", addr)
		fmt.Printf("%s\n", body)

		// 使用 APIClient 发送请求
		var result response.Response
		err = client.DoRequest("POST", "/api/v1/ping", reqBody, &result)
		if err != nil {
			return fmt.Errorf("request failed: %w", err)
		}

		// 输出响应
		respBody, _ := json.Marshal(result)
		fmt.Printf("-- Receiving response\n%s\n", respBody)

		return nil
	},
}

// 初始化命令行参数
func init() {
	// 为 ping 命令添加参数message,用来指定要发送的消息
	pingCmd.Flags().StringP("message", "m", "default message", "Message to send")
	rootCmd.AddCommand(pingCmd)
}
