// auth_service.go
package service

import (
	"errors"
	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 自定义声明

const TokenExpiry = time.Hour * 24

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
				Token:   "",
				Message: "用户未找到",
			}
		case errors.Is(err, ErrDatabase):
			return types.LoginResult{
				Success: false,
				Token:   "",
				Message: "数据库错误，请稍后重试",
			}
		default:
			return types.LoginResult{
				Success: false,
				Token:   "",
				Message: "认证服务内部错误",
			}
		}
	}

	if !ok {
		return types.LoginResult{
			Success: false,
			Token:   "",
			Message: "用户名或密码不正确",
		}
	}
	// 第二步：认证通过，开始签发 JWT
	// 查询用户以获取 userID（AuthenticateUser 只返回 true/false）
	var user model.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return types.LoginResult{
			Success: false,
			Token:   "",
			Message: "用户数据加载失败",
		}
	}

	// 构建自定义声明
	claims := types.CustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "imperishable-gate",
			Subject:   "auth",
		},
	}

	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWTSecret)
	if err != nil {
		return types.LoginResult{
			Success: false,
			Token:   "",
			Message: "令牌生成失败",
		}
	}

	return types.LoginResult{
		Success: true,
		Token:   signedToken,
		Message: "登录成功",
	}
}
