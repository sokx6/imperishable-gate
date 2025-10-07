package register

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// UserInput 存储用户注册输入
type UserInput struct {
	Username string
	Email    string
	Password string
}

// ReadUserInput 从标准输入读取用户注册信息
func ReadUserInput() (*UserInput, error) {
	reader := bufio.NewScanner(os.Stdin)

	// 读取用户名
	username, err := readUsername(reader)
	if err != nil {
		return nil, err
	}

	// 读取邮箱
	email, err := readEmail(reader)
	if err != nil {
		return nil, err
	}

	// 读取密码
	password, err := readPassword(reader)
	if err != nil {
		return nil, err
	}

	// 确认密码
	if err := confirmPassword(reader, password); err != nil {
		return nil, err
	}

	return &UserInput{
		Username: username,
		Email:    email,
		Password: password,
	}, nil
}

// readUsername 读取并验证用户名
func readUsername(reader *bufio.Scanner) (string, error) {
	fmt.Print("Please enter your username (3-32 characters): ")
	if !reader.Scan() {
		return "", fmt.Errorf("failed to read username")
	}
	username := strings.TrimSpace(reader.Text())
	if username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}
	if len(username) < 3 || len(username) > 32 {
		return "", fmt.Errorf("username must be between 3 and 32 characters")
	}
	return username, nil
}

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

// readPassword 读取并验证密码
func readPassword(reader *bufio.Scanner) (string, error) {
	fmt.Print("Please enter your password (minimum 6 characters): ")
	if !reader.Scan() {
		return "", fmt.Errorf("failed to read password")
	}
	password := reader.Text()
	if len(password) < 6 {
		return "", fmt.Errorf("password must be at least 6 characters long")
	}
	return password, nil
}

// confirmPassword 确认密码
func confirmPassword(reader *bufio.Scanner, password string) error {
	fmt.Print("Please confirm your password: ")
	if !reader.Scan() {
		return fmt.Errorf("failed to read password confirmation")
	}
	confirmPassword := reader.Text()
	if password != confirmPassword {
		return fmt.Errorf("passwords do not match")
	}
	return nil
}
