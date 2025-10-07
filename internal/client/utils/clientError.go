package utils

import (
	"errors"
)

// 定义常见错误

// 无访问令牌错误
var ErrNoAccessToken = errors.New("no access token found")

// 无刷新令牌错误
var ErrNoRefreshToken = errors.New("no refresh token found")

// 令牌过期错误
var ErrTokenExpired = errors.New("access token expired, please refresh")

// 无效令牌错误
var ErrInvalidToken = errors.New("invalid token, please login again")
