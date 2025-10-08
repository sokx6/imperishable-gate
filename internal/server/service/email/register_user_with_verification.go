package email

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"imperishable-gate/internal/server/service/common"
)

// RegisterUserWithVerification 使用事务处理用户注册和邮件发送
// 如果邮件发送失败，则回滚用户创建，确保数据一致性
func RegisterUserWithVerification(username, email, password string) error {
	// 1. 校验输入
	if username == "" || email == "" || password == "" {
		return common.ErrInvalidRegister
	}

	// 2. 检查用户名或邮箱是否已存在（在事务外检查，提高性能）
	var existingUser model.User

	// 检查用户名是否已存在
	if err := database.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return common.ErrNameAlreadyExists
	}

	// 检查邮箱是否已存在
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return common.ErrEmailAlreadyExists
	}

	// 3. 哈希密码
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	// 4. 先不生成验证码，先创建用户，后续统一用 sendVerificationCodeAndSaveRecord

	// 6. 开始事务：创建用户和验证记录
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
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
		return common.ErrDatabase
	}

	// 8. 创建验证记录并发送邮件（复用 sendVerificationCodeAndSaveRecord，传入事务）
	err = sendVerificationCodeAndSaveRecord(tx, newUser.ID, email)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 9. 提交事务
	if err := tx.Commit().Error; err != nil {
		return common.ErrDatabase
	}

	return nil
}
