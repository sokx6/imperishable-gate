package service

import (
	"bytes"
	"encoding/json"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"
)

func Login(addr, username, password string) (accessToken, refreshToken string, err error) {
	// 构建登录请求体
	reqBody := request.LoginRequest{
		Username: username,
		Password: password,
	}

	// 将请求体编码为 JSON
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", "", err
	}

	// 构建 HTTP 请求
	var resp *http.Response
	resp, err = http.Post(addr+"/api/v1/login", "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return "", "", err
	}
	// 确保响应体在函数退出时关闭
	defer resp.Body.Close()

	// 处理非 200 OK 的响应
	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusBadRequest:
			// 处理错误请求
			err = ErrInvalidRequest
		case http.StatusUnauthorized:
			// 处理认证失败
			err = ErrAuthenticationFailed
		case http.StatusNotFound:
			// 处理用户不存在
			err = ErrUserNotFound
		case http.StatusInternalServerError:
			// 处理服务器内部错误
			err = ErrInternalServer
		default:
			// 处理其他未知错误
			err = ErrUnknown
		}

	}

	// 解析响应体
	var loginResp response.LoginResult
	err = json.NewDecoder(resp.Body).Decode(&loginResp)
	if err != nil {
		// 解析失败
		return "", "", err
	}
	// 返回访问令牌和刷新令牌
	accessToken = loginResp.AccessToken
	refreshToken = loginResp.RefreshToken
	return
}
