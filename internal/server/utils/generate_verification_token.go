package utils

import (
	"crypto/rand"
	"fmt"
)

// GenerateVerificationToken 生成唯一验证令牌
func GenerateVerificationToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", bytes), nil
}
