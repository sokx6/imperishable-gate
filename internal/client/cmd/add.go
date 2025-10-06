package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"

	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new link to the server database",
	RunE: func(cmd *cobra.Command, args []string) error {
		link, _ := cmd.Flags().GetString("link")

		// 构造请求体
		reqBody := request.AddRequest{
			Link: link,
		}

		// 将请求体编码为 JSON
		body, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}

		// 构造请求 URL
		url := fmt.Sprintf("http://%s/api/v1/links", addr)
		fmt.Printf("-- Requesting POST method to %s with payload\n", url)
		fmt.Printf("%s\n", body)

		request, _ := http.NewRequest("POST", url, strings.NewReader(string(body)))
		request.Header.Set("Authorization", "Bearer "+accessToken)
		request.Header.Set("Content-Type", "application/json")

		// 发起 POST 请求
		resp, err := http.DefaultClient.Do(request)
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
		var result response.Response
		if err := json.Unmarshal(respBody, &result); err != nil {
			return fmt.Errorf("invalid JSON response: %w", err)
		}

		fmt.Printf("-- Receiving response\n%s\n", respBody)

		return nil
	},
}

// 初始化命令行参数
func init() {
	// 为 add 命令添加参数link，用来指定要添加的链接
	addCmd.Flags().StringP("link", "l", "", "link to add")
	rootCmd.AddCommand(addCmd)
}
