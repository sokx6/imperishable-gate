package register

import (
	"bufio"
	"fmt"
	"strings"
)

// readEmail 读取并验证邮箱
func readEmail(reader *bufio.Scanner) (string, error) {
	fmt.Print("Please enter your email: ")
	if !reader.Scan() {
		return "", fmt.Errorf("failed to read email")
	}
	email := strings.TrimSpace(reader.Text())
	if email == "" {
		return "", fmt.Errorf("email cannot be empty")
	}
	// 简单的邮箱格式验证
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return "", fmt.Errorf("invalid email format")
	}
	return email, nil
}
