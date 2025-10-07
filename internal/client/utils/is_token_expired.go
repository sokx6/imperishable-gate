package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func IsTokenExpired(accessToken string) (bool, error) {
	// JWT 令牌的声明
	var claims jwt.RegisteredClaims

	if accessToken == "" {
		// 处理无效令牌
		return true, nil
	}

	// 解析但不验证签名
	token, _, err := jwt.NewParser().ParseUnverified(accessToken, &claims)
	if err != nil {
		// 处理解析错误
		return true, err
	}
	// 检查令牌是否有效
	if !token.Valid {
		return true, nil
	}

	// 如果即将在30秒内过期，也视为“需要刷新”
	if claims.ExpiresAt != nil {
		return claims.ExpiresAt.Before(time.Now().Add(30 * time.Second)), nil
	}

	return true, nil // 没有 exp 按过期处理更安全
}
