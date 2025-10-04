package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"

	types "imperishable-gate/internal"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the links stored in the server",
	RunE: func(cmd *cobra.Command, args []string) error {

		// 构造请求 URL
		url := fmt.Sprintf("http://%s/api/v1/links", Config.Addr)
		fmt.Printf("-- Requesting GET method to %s\n", url)

		// 发起 GET 请求
		resp, err := http.Get(url)
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
		var result types.ListResponse
		if err := json.Unmarshal(respBody, &result); err != nil {
			return fmt.Errorf("invalid JSON response: %w", err)
		}

		fmt.Println("-- Receiving response")

		for _, link := range result.Data {
			fmt.Printf("ID: %d\nURL: %s\nTags: %v\nNames: %v\nRemark: %s\n\n",
				link.ID, link.Url, link.Tags, link.Names, link.Remark)
		}

		return nil
	},
}

// 初始化命令行参数
func init() {
	// 为 list 命令添加参数host，用来指定服务器地址
	listCmd.Flags().StringP("host", "H", "127.0.0.1:8080", "Server host:port to send list")
	rootCmd.AddCommand(listCmd)
}
