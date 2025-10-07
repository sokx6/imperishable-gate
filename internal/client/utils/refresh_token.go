package utils

import (
	"fmt"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

// RefreshToken 使用刷新令牌获取新的访问令牌
// 参数:
//   - refreshToken: 刷新令牌字符串
//   - addr: 服务器地址
//
// 返回:
//   - string: 新的访问令牌
//   - error: 错误信息
func RefreshToken(refreshToken, addr string) (string, error) {
	// 创建 API 客户端
	client := NewAPIClient(addr, "")

	// 构建刷新令牌请求体
	reqBody := request.RefreshRequest{RefreshToken: refreshToken}

	// 使用 APIClient 发送请求
	var refreshResp response.RefreshResponse
	err := client.DoRequest("POST", "/api/v1/refresh", reqBody, &refreshResp)
	if err != nil {
		return "", fmt.Errorf("failed to refresh token: %w", err)
	}

	// 检查是否成功获取到访问令牌
	if refreshResp.AccessToken == "" {
		return "", ErrNoAccessToken
	}

	// 返回新的访问令牌
	return refreshResp.AccessToken, nil
}
