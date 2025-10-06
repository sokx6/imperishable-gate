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

type APIClient struct {
	BaseURL      string
	AccessToken  string
	Client       *http.Client
	OutputFormat string // "table", "json"
}

func NewAPIClient(baseURL, accessToken string) *APIClient {
	if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
		baseURL = "http://" + baseURL
	}
	return &APIClient{
		BaseURL:      baseURL,
		AccessToken:  accessToken,
		Client:       &http.Client{},
		OutputFormat: "table",
	}
}

// 发起请求并解码通用 Response
func (c *APIClient) DoRequest(method, path string, body interface{}, result interface{}) error {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return err
	}

	// 设置头信息
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response failed: %w", err)
	}

	// 先解析通用错误结构
	var commonResp response.Response
	if err := json.Unmarshal(respBody, &commonResp); err != nil {
		return fmt.Errorf("invalid JSON response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("request failed [%d]: %s", resp.StatusCode, commonResp.Message)
	}

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
	}

	return nil
}
