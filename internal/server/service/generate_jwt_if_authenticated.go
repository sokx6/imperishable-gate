// auth_service.go
package service

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/types/response"
	"time"
)

// 自定义声明

const TokenExpiry = time.Hour * 24
const (
	AccessExpiry  = 15 * time.Minute
	RefreshExpiry = 7 * 24 * time.Hour // 7天
)

var JWTSecret = []byte("locxlfjalkfjelifalngl")

// GenerateJWTIfAuthenticated 处理登录并签发JWT
func GenerateJWTIfAuthenticated(username, password string) response.LoginResponse {
	// 第一步：调用 AuthenticateUser 验证用户
	ok, err := AuthenticateUser(username, password)
	if err != nil {
		switch {
		case errors.Is(err, ErrUsernameNotFound):
			return response.LoginResponse{
				Success: false,

				Message: "User not found",
			}
		case errors.Is(err, ErrDatabase):
			return response.LoginResponse{
				Success: false,

				Message: "Database error, please try again later",
			}
		default:
			return response.LoginResponse{
				Success: false,

				Message: "Internal authentication service error",
			}
		}
	}

	if !ok {
		return response.LoginResponse{
			Success: false,

			Message: "Invalid username or password",
		}
	}
	// 第二步：认证通过，开始签发 JWT
	// 查询用户以获取 userID（AuthenticateUser 只返回 true/false）
	var user model.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return response.LoginResponse{
			Success: false,
			Message: "Failed to load user data",
		}
	}
	// 生成 Access Token
	accessToken, err := GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		return response.LoginResponse{Success: false, Message: "Failed to generate access token"}
	}

	// 生成并存储 Refresh Token
	refreshToken, err := GenerateRefreshToken(user.ID)
	if err != nil {
		return response.LoginResponse{Success: false, Message: "Failed to generate refresh token"}
	}

	return response.LoginResponse{
		Success:      true,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Message:      "Login successful",
	}
}
