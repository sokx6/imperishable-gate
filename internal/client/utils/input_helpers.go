package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	// MinPasswordLength 最小密码长度
	MinPasswordLength = 6
	// MinUsernameLength 最小用户名长度
	MinUsernameLength = 3
	// MaxUsernameLength 最大用户名长度
	MaxUsernameLength = 32
)

// ReadPasswordWithConfirm 读取密码并确认
func ReadPasswordWithConfirm(prompt, confirmPrompt string, minLen int) (string, error) {
	if confirmPrompt == "" {
		confirmPrompt = "Please confirm your password: "
	}

	password, err := ReadPassword(prompt, minLen)
	if err != nil {
		return "", err
	}

	if err := ConfirmInput(confirmPrompt, password); err != nil {
		return "", err
	}

	return password, nil
}

// ReadEmail 读取并验证邮箱
func ReadEmail(prompt string) (string, error) {
	if prompt == "" {
		prompt = "Please enter your email: "
	}

	email, err := ReadLine(prompt)
	if err != nil {
		return "", err
	}

	if err := ValidateEmail(email); err != nil {
		return "", err
	}

	return email, nil
}

// ReadPassword 读取并验证密码
func ReadPassword(prompt string, minLen int) (string, error) {
	if prompt == "" {
		prompt = fmt.Sprintf("Please enter your password (minimum %d characters): ", minLen)
	}
	if minLen <= 0 {
		minLen = MinPasswordLength
	}

	password, err := ReadLine(prompt)
	if err != nil {
		return "", err
	}

	if err := ValidatePassword(password, minLen); err != nil {
		return "", err
	}

	return password, nil
}

// ReadUsername 读取并验证用户名
func ReadUsername(prompt string) (string, error) {
	if prompt == "" {
		prompt = fmt.Sprintf("Please enter your username (%d-%d characters): ", MinUsernameLength, MaxUsernameLength)
	}

	username, err := ReadLine(prompt)
	if err != nil {
		return "", err
	}

	if err := ValidateUsername(username); err != nil {
		return "", err
	}

	return username, nil
}

// ReadLine 读取一行输入
func ReadLine(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewScanner(os.Stdin)
	if !reader.Scan() {
		return "", fmt.Errorf("failed to read input")
	}
	return strings.TrimSpace(reader.Text()), nil
}

// ConfirmInput 确认两次输入一致
func ConfirmInput(prompt string, origin string) error {
	input, err := ReadLine(prompt)
	if err != nil {
		return err
	}
	if input != origin {
		return fmt.Errorf("inputs do not match")
	}
	return nil
}
