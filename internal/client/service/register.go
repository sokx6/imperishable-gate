package service

import (
	"fmt"
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

// Register 向服务器发送用户注册请求
func Register(addr, username, email, password string) error {
	// 创建 API 客户端
	client := utils.NewAPIClient(addr, "")

	// 构建注册请求体
	reqBody := request.UserRegisterRequest{
		Username: username,
		Email:    email,
		Password: password,
	}

	// 使用 APIClient 发送请求
	var registerResp response.Response
	err := client.DoRequest("POST", "/api/v1/register", reqBody, &registerResp)
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}

	return nil
}
