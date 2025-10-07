package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"
)

// Register 向服务器发送用户注册请求
func Register(addr, username, email, password string) error {
	// 构建注册请求体
	reqBody := request.UserRegisterRequest{
		Username: username,
		Email:    email,
		Password: password,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	// 发送 HTTP POST 请求
	resp, err := http.Post(addr+"/api/v1/register", "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}
	defer resp.Body.Close()

	// 处理响应状态码
	if resp.StatusCode != http.StatusOK {
		var errResp response.Response
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err == nil {
			return fmt.Errorf("registration failed: %s", errResp.Message)
		}

		switch resp.StatusCode {
		case http.StatusBadRequest:
			return ErrInvalidRequest
		case http.StatusConflict:
			// 用户名或邮箱已存在
			return ErrUserNameExists
		case http.StatusInternalServerError:
			return ErrInternalServer
		default:
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
	}

	// 解析成功响应
	var registerResp response.Response
	if err := json.NewDecoder(resp.Body).Decode(&registerResp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
