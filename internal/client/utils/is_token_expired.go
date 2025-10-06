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

	token, _, err := jwt.NewParser().ParseUnverified(accessToken, &claims)
	if err != nil {
		return true, err
	}

	if !token.Valid {
		return true, nil
	}

	// 如果即将在30秒内过期，也视为“需要刷新”
	if claims.ExpiresAt != nil {
		return claims.ExpiresAt.Before(time.Now().Add(30 * time.Second)), nil
	}

	return true, nil // 没有 exp？按过期处理更安全
}
