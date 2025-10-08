package register

import (
	"bufio"
	"os"
)

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
