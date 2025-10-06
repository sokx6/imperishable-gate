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

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", "", err
	}

	var resp *http.Response
	resp, err = http.Post(addr+"/api/v1/login", "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusBadRequest:
			err = ErrInvalidRequest
		case http.StatusUnauthorized:
			err = ErrAuthenticationFailed
		case http.StatusNotFound:
			err = ErrUserNotFound
		case http.StatusInternalServerError:
			err = ErrInternalServer
		default:
			err = ErrUnknown
		}

	}

	var loginResp response.LoginResult

	err = json.NewDecoder(resp.Body).Decode(&loginResp)
	if err != nil {
		return "", "", err
	}

	accessToken = loginResp.AccessToken
	refreshToken = loginResp.RefreshToken
	return
}
