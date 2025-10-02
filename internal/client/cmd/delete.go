// cmd/delete.go 或内联在现有文件中

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"

	types "imperishable-gate/internal"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete links from the server database using query parameters",
	RunE: func(cmd *cobra.Command, args []string) error {
		links, _ := cmd.Flags().GetStringSlice("links")

		if len(links) == 0 {
			return fmt.Errorf("no links provided for deletion. Use -l or --links to specify one or more URLs")
		}

		// 构建请求 URL 并添加多个 link 查询参数
		apiURL := fmt.Sprintf("http://%s/api/v1/links/delete", Config.Addr)
		u, err := url.Parse(apiURL)
		if err != nil {
			return fmt.Errorf("invalid base URL: %w", err)
		}

		q := u.Query()
		for _, link := range links {
			q.Add("link", link) // 每个 link 作为单独参数
		}
		u.RawQuery = q.Encode() // 设置编码后的查询字符串

		requestURL := u.String()
		fmt.Printf("-- Sending DELETE request to: %s\n", requestURL)

		// 创建 DELETE 请求，NO BODY
		req, err := http.NewRequest(http.MethodDelete, requestURL, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		// 发起请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("request failed: %w", err)
		}
		defer resp.Body.Close()

		// 读取响应体
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response: %w", err)
		}

		// 解析 JSON 响应
		var result types.DeleteResponse
		if err := json.Unmarshal(respBody, &result); err != nil {
			return fmt.Errorf("invalid JSON response: %w", err)
		}

		// 输出结果
		fmt.Printf("Response Code: %d\n", result.Code)
		fmt.Printf("Message: %s\n", result.Message)
		if result.Data != nil {
			dataBytes, _ := json.MarshalIndent(result.Data, "", "  ")
			fmt.Printf("Data:\n%s\n", dataBytes)
		}

		// 根据返回 code 决定是否报错（可选）
		if result.Code != 0 {
			return fmt.Errorf("deletion failed with code %d", result.Code)
		}

		return nil
	},
}

// 初始化命令行参数
func init() {
	deleteCmd.Flags().StringSliceP("links", "l", []string{}, "Links to delete (repeat -l or use comma-separated values)")
	rootCmd.AddCommand(deleteCmd)
}
