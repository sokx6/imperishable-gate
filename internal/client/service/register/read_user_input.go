package register

import (
	"imperishable-gate/internal/client/utils"
)

// ReadUserInput 从标准输入读取用户注册信息
func ReadUserInput() (*UserInput, error) {
	// 读取用户名
	username, err := utils.ReadUsername("")
	if err != nil {
		return nil, err
	}

	// 读取邮箱
	email, err := utils.ReadEmail("")
	if err != nil {
		return nil, err
	}

	// 读取并确认密码
	password, err := utils.ReadPasswordWithConfirm(
		"Please enter your password (minimum 6 characters): ",
		"Please confirm your password: ",
		utils.MinPasswordLength,
	)
	if err != nil {
		return nil, err
	}

	return &UserInput{
		Username: username,
		Email:    email,
		Password: password,
	}, nil
}
