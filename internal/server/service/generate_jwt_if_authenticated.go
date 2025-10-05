// auth_service.go
package service

import (
	"errors"
	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
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
func GenerateJWTIfAuthenticated(username, password string) types.LoginResult {
	// 第一步：调用 AuthenticateUser 验证用户
	ok, err := AuthenticateUser(username, password)
	if err != nil {
		switch {
		case errors.Is(err, ErrUsernameNotFound):
			return types.LoginResult{
				Success: false,

				Message: "用户未找到",
			}
		case errors.Is(err, ErrDatabase):
			return types.LoginResult{
				Success: false,

				Message: "数据库错误，请稍后重试",
			}
		default:
			return types.LoginResult{
				Success: false,

				Message: "认证服务内部错误",
			}
		}
	}

	if !ok {
		return types.LoginResult{
			Success: false,

			Message: "用户名或密码不正确",
		}
	}
	// 第二步：认证通过，开始签发 JWT
	// 查询用户以获取 userID（AuthenticateUser 只返回 true/false）
	var user model.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return types.LoginResult{
			Success: false,
			Message: "用户数据加载失败",
		}
	}
	// 生成 Access Token
	accessToken, err := GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		return types.LoginResult{Success: false, Message: "访问令牌生成失败"}
	}

	// 生成并存储 Refresh Token
	refreshToken, err := GenerateRefreshToken(user.ID)
	if err != nil {
		return types.LoginResult{Success: false, Message: "刷新令牌生成失败"}
	}

	return types.LoginResult{
		Success:      true,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Message:      "登录成功",
	}
}
