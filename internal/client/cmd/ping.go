package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"

	types "imperishable-gate/internal"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Send a ping request to the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		msg, _ := cmd.Flags().GetString("message")

		// 构造请求体
		reqBody := types.PingRequest{
			Action:  "ping",
			Message: msg,
		}

		// 将请求体编码为 JSON
		body, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}

		// 构造请求 URL
		url := fmt.Sprintf("http://%s/api/v1/ping", Config.Addr)
		fmt.Printf("-- Requesting POST method to %s with payload\n", url)
		fmt.Printf("%s\n", body)

		// 发起 POST 请求
		resp, err := http.Post(url, "application/json", strings.NewReader(string(body)))
		if err != nil {
			return fmt.Errorf("request failed: %w", err)
		}

		// 确保响应体最终被关闭
		defer resp.Body.Close()

		// 读取响应体
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response: %w", err)
		}

		// 讲respBody响应体解析为JSON
		// 并存储到result中
		var result types.PingResponse
		if err := json.Unmarshal(respBody, &result); err != nil {
			return fmt.Errorf("invalid JSON response: %w", err)
		}

		fmt.Printf("-- Receiving response\n%s\n", respBody)

		return nil
	},
}

// 初始化命令行参数
func init() {
	// 为 ping 命令添加参数message，用来指定要发送的消息
	pingCmd.Flags().StringP("message", "m", "default message", "Message to send")
	rootCmd.AddCommand(pingCmd)
}
