package utils

import (
	"fmt"
	"strings"
)

// ValidateEmail 简单邮箱格式校验
func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

// ValidatePassword 密码长度校验
func ValidatePassword(password string, minLen int) error {
	if len(password) < minLen {
		return fmt.Errorf("password must be at least %d characters long", minLen)
	}
	return nil
}
