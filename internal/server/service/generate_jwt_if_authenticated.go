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
func GenerateJWTIfAuthenticated(username, password string) response.LoginResult {
	// 第一步：调用 AuthenticateUser 验证用户
	err := AuthenticateUser(username, password)
	if err != nil {
		// 处理不同类型的错误返回
		switch {
		case errors.Is(err, ErrUsernameNotFound):
			// 用户不存在
			return response.LoginResult{
				Success: false,

				Message: "User not found",
			}
		case errors.Is(err, ErrDatabase):
			// 数据库错误
			return response.LoginResult{
				Success: false,

				Message: "Database error, please try again later",
			}
		case errors.Is(err, ErrInvalidPassword):
			// 密码错误
			return response.LoginResult{
				Success: false,

				Message: "Invalid username or password",
			}
		default:
			return response.LoginResult{
				// 其他未知错误
				Success: false,

				Message: "Internal authentication service error",
			}
		}
	}

	// 第二步：认证通过，开始签发 JWT
	// 查询用户以获取 userID（AuthenticateUser 只返回 true/false）
	var user model.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return response.LoginResult{
			Success: false,
			Message: "Failed to load user data",
		}
	}
	// 生成 Access Token
	accessToken, err := GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		return response.LoginResult{Success: false, Message: "Failed to generate access token"}
	}

	// 生成并存储 Refresh Token
	refreshToken, err := GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		return response.LoginResult{Success: false, Message: "Failed to generate refresh token"}
	}
	// 返回成功响应
	return response.LoginResult{
		Success:      true,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Message:      "Login successful",
	}
}
