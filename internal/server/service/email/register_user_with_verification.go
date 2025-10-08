package email

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/service/common"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/utils/logger"
)

// RegisterUserWithVerification 使用事务处理用户注册和邮件发送
// 如果邮件发送失败，则回滚用户创建，确保数据一致性
func RegisterUserWithVerification(username, email, password string) error {
	// 1. 校验输入
	if username == "" || email == "" || password == "" {
		logger.Warning("Invalid registration request: missing username, email or password")
		return common.ErrInvalidRegister
	}

	logger.Info("Registering new user: %s, email: %s", username, email)

	// 2. 检查用户名或邮箱是否已存在（在事务外检查，提高性能）
	var existingUser model.User

	// 检查用户名是否已存在
	if err := database.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		logger.Warning("Username already exists: %s", username)
		return common.ErrNameAlreadyExists
	}

	// 检查邮箱是否已存在
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		logger.Warning("Email already exists: %s", email)
		return common.ErrEmailAlreadyExists
	}

	// 3. 哈希密码
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		logger.Error("Failed to hash password for user %s: %v", username, err)
		return err
	}

	// 6. 开始事务：创建用户和验证记录
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		logger.Error("Failed to start transaction for user registration: %v", tx.Error)
		return common.ErrDatabase
	}

	// 7. 创建用户
	newUser := model.User{
		Username:      username,
		Email:         email,
		Password:      hashedPassword,
		EmailVerified: false,
	}

	// 创建用户失败
	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		logger.Error("Failed to create user %s: %v", username, err)
		return common.ErrDatabase
	}

	logger.Debug("User created in database: %s (ID: %d)", username, newUser.ID)

	// 8. 创建验证记录并发送邮件（复用 sendVerificationCodeAndSaveRecord，传入事务）
	err = sendVerificationCodeAndSaveRecord(tx, newUser.ID, email)
	if err != nil {
		tx.Rollback()
		logger.Error("Failed to send verification email for user %s: %v", username, err)
		return err
	}

	// 9. 提交事务
	if err := tx.Commit().Error; err != nil {
		logger.Error("Failed to commit transaction for user %s: %v", username, err)
		return common.ErrDatabase
	}

	logger.Success("User %s registered successfully, verification email sent to %s", username, email)
	return nil
}
