package service

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"time"

	"gorm.io/gorm"
)

// VerifyEmailByCode 验证邮箱失败时记录尝试次数
func VerifyEmailByCode(email, code string) (userId uint, err error) {
	var verification model.EmailVerification

	// 查找该邮箱的活跃验证记录
	err = database.DB.Where("email = ? AND used = ? AND expires_at > ?", email, false, time.Now()).
		First(&verification).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, ErrInvalidVerificationCode
	} else if err != nil {
		return 0, ErrDatabase
	}

	// 检查尝试次数
	if verification.AttemptCount >= 5 {
		verification.Used = true
		database.DB.Save(&verification)
		return 0, ErrTooManyAttempts
	}

	// 检查验证码是否正确
	if verification.Code != code {
		// 验证码错误，增加尝试次数
		verification.AttemptCount++
		if err := database.DB.Save(&verification).Error; err != nil {
			return 0, ErrDatabase
		}
		return 0, ErrInvalidVerificationCode
	}

	// 验证成功
	verification.Used = true
	if err := database.DB.Save(&verification).Error; err != nil {
		return 0, ErrDatabase
	}

	return verification.UserID, nil
}

func VerifyEmailAndRegister(email, code string) error {
	// 验证验证码并且获取用户ID
	userId, err := VerifyEmailByCode(email, code)
	if err != nil {
		return err
	}
	// 更新用户状态
	now := time.Now()
	if err := database.DB.Model(&model.User{}).
		Where("id = ?", userId).
		Updates(map[string]interface{}{
			"email_verified":    true,
			"email_verified_at": &now,
		}).Error; err != nil {
		return ErrDatabase
	}
	return nil
}

func VerifyEmailAndResetPassword(email, code, newPassword string) error {
	// 验证验证码并且获取用户ID
	userId, err := VerifyEmailByCode(email, code)
	if err != nil {
		return err
	}
	// 更新用户密码
	newPassword, err = utils.HashPassword(newPassword)
	if err != nil {
		return ErrDatabase
	}
	if err := database.DB.Model(&model.User{}).
		Where("id = ?", userId).
		Updates(map[string]interface{}{
			"password": newPassword,
		}).Error; err != nil {
		return ErrDatabase
	}
	return nil
}

func VerifyEmailByUsernameAndResetPassword(username, code, newPassword string) error {
	// 根据用户名查找用户邮箱
	var user model.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return ErrDatabase
	}
	return VerifyEmailAndResetPassword(user.Email, code, newPassword)
}
