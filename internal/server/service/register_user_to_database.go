package service

import (
	"imperishable-gate/internal/model" // 替换为你的实际模块路径
	"imperishable-gate/internal/server/database"

	"imperishable-gate/internal/server/utils"
)

// RegisterUser 注册新用户：存储用户名、邮箱和哈希后的密码到数据库
func RegisterUser(username, email, password string) error {
	// 1. 校验输入
	if username == "" || email == "" || password == "" {
		return ErrInvalidRegister
	}

	// 2. 检查用户名或邮箱是否已存在
	var existingUser model.User

	// 检查用户名是否已存在
	if err := database.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return ErrNameAlreadyExists
	}

	// 检查邮箱是否已存在
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return ErrEmailAlreadyExists
	}

	// 3. 哈希密码
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	// 4. 构造用户对象
	newUser := model.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	// 5. 存入数据库
	if err := database.DB.Create(&newUser).Error; err != nil {
		return ErrDatabase
	}

	return nil
}
