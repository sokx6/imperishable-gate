// client/utils/http_client.go
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"imperishable-gate/internal/types/response"
)

// APIClient 是一个简单的 HTTP 客户端，用于与 Imperishable Gate API 交互
type APIClient struct {
	BaseURL     string
	AccessToken string
	Client      *http.Client
}

// NewAPIClient 创建并返回一个新的 APIClient 实例
func NewAPIClient(baseURL, accessToken string) *APIClient {
	// 确保 BaseURL 以 http:// 或 https:// 开头
	if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
		baseURL = "http://" + baseURL
	}
	// 返回一个新的 APIClient 实例
	return &APIClient{
		BaseURL:     baseURL,
		AccessToken: accessToken,
		Client:      &http.Client{},
	}
}

// 发起请求并解码通用 Response
// 参数:
// - method: HTTP 方法 (GET, POST, etc.)
// - path: API 路径
// - body: 请求体 (可以为 nil)
// - result: 用于存储解码后的响应数据的指针 (可以为 nil)
// 返回值:
// - error: 如果请求失败或响应解析失败，则返回错误
// 其他情况下返回 nil
func (c *APIClient) DoRequest(method, path string, body interface{}, result interface{}) error {
	// 准备请求体
	var reqBody io.Reader
	if body != nil {
		// 将 body 编码为 JSON
		jsonData, err := json.Marshal(body)
		if err != nil {
			// 处理请求错误
			return fmt.Errorf("failed to marshal request: %w", err)
		}
		// 创建请求体
		reqBody = bytes.NewBuffer(jsonData)
	}
	// 创建 url = c.BaseURL + path
	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	// 创建 HTTP 请求
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return err
	}

	// 设置头信息
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := c.Client.Do(req)
	if err != nil {
		// 处理请求错误
		return fmt.Errorf("request failed: %w", err)
	}
	// 确保响应体在函数退出时关闭
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		// 处理响应错误
		return fmt.Errorf("read response failed: %w", err)
	}

	// 先解析通用错误结构
	var commonResp response.Response
	if err := json.Unmarshal(respBody, &commonResp); err != nil {
		// 处理解析错误
		return fmt.Errorf("invalid JSON response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// 处理错误响应
		return fmt.Errorf("request failed [%d]: %s", resp.StatusCode, commonResp.Message)
	}

	if result != nil {
		// 解析具体的响应数据
		if err := json.Unmarshal(respBody, result); err != nil {
			// 处理解析错误
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
	}

	return nil
}
