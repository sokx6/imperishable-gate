package service

import (
	types "imperishable-gate/internal"

	"github.com/golang-jwt/jwt/v5"
)

// ParseJWT 解析并验证 JWT token
func ParseJWT(tokenString, secret string) (types.UserInfo, error) {
	var userInfo types.UserInfo

	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &types.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		// 1. 验证签名算法是否为 HS256
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrAuthenticationFailed
		}
		// 2. 返回用于验证签名的密钥
		return []byte(secret), nil
	})

	// 3. 检查解析过程是否有错误
	if err != nil {
		return types.UserInfo{}, ErrAuthenticationFailed
	}

	// 4. 检查 token 是否有效（包括签名有效性和时间范围 exp/iat/nbf）
	if !token.Valid {
		return types.UserInfo{}, ErrAuthenticationFailed
	}

	// 5. 提取 claims
	claims, ok := token.Claims.(*types.CustomClaims)
	if !ok {
		return types.UserInfo{}, ErrAuthenticationFailed
	}

	// 7. 构造返回的 types.UserInfo
	userInfo = types.UserInfo{
		UserID:   claims.UserID,
		Username: claims.Username,
	}

	return userInfo, nil
}
