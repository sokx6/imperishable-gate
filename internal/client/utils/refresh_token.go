package utils

import (
	"bytes"
	"encoding/json"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"
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
	// 构建刷新令牌请求体
	reqBody := request.RefreshRequest{RefreshToken: refreshToken}
	// 将请求体序列化为 JSON 字节数组
	reqBytes, _ := json.Marshal(reqBody)

	// 向服务器发送 POST 请求以刷新令牌
	resp, err := http.Post(addr+"/api/v1/refresh", "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 解析响应体中的访问令牌
	var accessToken response.RefreshResponse
	json.NewDecoder(resp.Body).Decode(&accessToken)

	// 检查是否成功获取到访问令牌
	if accessToken.AccessToken == "" {
		return "", ErrNoAccessToken
	}

	// 返回新的访问令牌
	return accessToken.AccessToken, nil
}
