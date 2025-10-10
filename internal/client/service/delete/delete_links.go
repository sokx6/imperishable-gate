package delete

import (
	"encoding/json"
	"fmt"
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/response"
	"net/url"
)

func DeleteLinks(links []string, addr string, accessToken string) error {
	// 规范化所有 URL（如果缺少协议则自动添加 https://）
	links = utils.NormalizeURLs(links)

	// 创建 API 客户端
	client := utils.NewAPIClient(addr, accessToken)

	// 构建查询参数
	q := url.Values{}
	for _, link := range links {
		q.Add("link", link) // 每个 link 作为单独参数
	}

	// 构建完整路径
	path := fmt.Sprintf("/api/v1/links?%s", q.Encode())

	// 使用 APIClient 发送请求
	var result response.Response
	err := client.DoRequest("DELETE", path, nil, &result)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	// 输出结果
	fmt.Printf("Message: %s\n", result.Message)
	if result.Links != nil {
		dataBytes, _ := json.MarshalIndent(result.Links, "", "  ")
		fmt.Printf("Data:\n%s\n", dataBytes)
	}

	return nil
}
