package service

import (
	"fmt"
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

func Login(addr, username, password string) (accessToken, refreshToken string, err error) {
	// 创建 API 客户端
	client := utils.NewAPIClient(addr, "")

	// 构建登录请求体
	reqBody := request.LoginRequest{
		Username: username,
		Password: password,
	}

	// 使用 APIClient 发送请求
	var loginResp response.LoginResponse
	err = client.DoRequest("POST", "/api/v1/login", reqBody, &loginResp)
	if err != nil {
		return "", "", fmt.Errorf("failed to connect to server: %w", err)
	}

	// 返回访问令牌和刷新令牌
	accessToken = loginResp.AccessToken
	refreshToken = loginResp.RefreshToken
	return accessToken, refreshToken, nil
}
