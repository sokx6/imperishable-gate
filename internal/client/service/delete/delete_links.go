package delete

import (
	"encoding/json"
	"fmt"
	"imperishable-gate/internal/types/response"
	"io"
	"net/http"
	"net/url"
)

func DeleteLinks(links []string, addr string, accessToken string) error {
	// 构建请求 URL 并添加多个 link 查询参数
	apiURL := fmt.Sprintf("%s/api/v1/links", addr)
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

	// 创建 DELETE 请求
	req, err := http.NewRequest(http.MethodDelete, requestURL, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
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
	var result response.Response
	if err := json.Unmarshal(respBody, &result); err != nil {
		return fmt.Errorf("invalid JSON response: %w", err)
	}

	// 输出结果

	fmt.Printf("Message: %s\n", result.Message)
	if result.Links != nil {
		dataBytes, _ := json.MarshalIndent(result.Links, "", "  ")
		fmt.Printf("Data:\n%s\n", dataBytes)
	}

	return nil
}
