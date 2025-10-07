package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func IsTokenExpired(accessToken string) (bool, error) {
	var claims jwt.RegisteredClaims

	if accessToken == "" {
		return true, nil
	}

	_, _, err := jwt.NewParser().ParseUnverified(accessToken, &claims)
	if err != nil {
		return true, err // 解析失败视为过期 + 返回错误原因
	}

	// 检查 exp 字段
	if claims.ExpiresAt != nil {
		// 如果将在 30 秒内过期，认为需要刷新（即“已过期”）
		return claims.ExpiresAt.Before(time.Now().Add(30 * time.Second)), nil
	}

	return true, nil // 无 exp 字段，默认视为过期（最安全行为）
}
