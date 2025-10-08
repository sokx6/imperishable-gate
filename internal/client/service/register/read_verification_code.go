package register

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadVerificationCode 读取邮箱验证码
func ReadVerificationCode(attempt, maxAttempts int) (string, error) {
	reader := bufio.NewScanner(os.Stdin)

	if attempt > 1 {
		fmt.Printf("\n[Attempt %d/%d] ", attempt, maxAttempts)
	}
	fmt.Print("Please enter the verification code (or 'resend' to get a new code): ")

	if !reader.Scan() {
		return "", fmt.Errorf("failed to read verification code")
	}

	code := strings.TrimSpace(reader.Text())
	if code == "" {
		return "", fmt.Errorf("verification code cannot be empty")
	}

	// 允许输入 "resend" 来重新发送验证码
	if strings.ToLower(code) == "resend" {
		return "resend", nil
	}

	// 验证码通常是6位数字
	if len(code) != 6 {
		return "", fmt.Errorf("verification code must be 6 digits")
	}

	return code, nil
}
