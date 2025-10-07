// auth_service.go
package service

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/types/response"
	"os"
)

// 自定义声明

var JWTSecret = getJWTSecret()

// getJWTSecret 从环境变量读取 JWT 密钥，如果未设置则使用默认值
func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// 默认值（仅用于开发环境，生产环境必须设置环境变量）
		secret = "kF2*oP4/tP3`jY0`nF7*tH4.xW6(jT7&nD1`rU4;oO8:cR2[eF0[oZ6&lN6_mW2.vA4=mJ1=rB5^jZ1,bO7<cO4!mZ9;dU3]oW9$bK2*wA8.pK0{zR9=wL7^rL3{qA5^"
	}
	return []byte(secret)
}

// GenerateJWTIfAuthenticated 处理登录并签发JWT
func GenerateJWTIfAuthenticated(username, password string) (response.LoginResponse, error) {
	// 第一步：调用 AuthenticateUser 验证用户
	err := AuthenticateUser(username, password)
	if err != nil {
		return response.LoginResponse{}, err
	}

	// 第二步：认证通过，开始签发 JWT
	// 查询用户以获取 userID（AuthenticateUser 只返回 true/false）
	var user model.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return response.LoginResponse{}, err
	}
	// 生成 Access Token
	accessToken, err := GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		return response.LoginResponse{}, err
	}

	// 生成并存储 Refresh Token
	refreshToken, err := GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		return response.LoginResponse{}, err
	}
	// 返回成功响应
	return response.LoginResponse{
		Success:      true,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Message:      "Login successful",
	}, nil
}
