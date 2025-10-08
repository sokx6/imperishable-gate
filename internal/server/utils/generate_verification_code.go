package utils

import (
	"crypto/rand"
	"math/big"
)

// GenerateVerificationCode 生成6位数字验证码
func GenerateVerificationCode() (string, error) {
	code := ""
	for i := 0; i < 6; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		code += num.String()
	}
	return code, nil
}
