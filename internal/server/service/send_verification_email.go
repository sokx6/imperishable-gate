package service

import (
	"fmt"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"
	"time"

	"gorm.io/gorm"
)

// sendVerificationCodeAndSaveRecord 生成验证码、保存验证记录并发送邮件，返回验证码和错误
// 支持传入事务的 DB

func sendVerificationCodeAndSaveRecord(tx *gorm.DB, userID uint, email string) error {
	code, err := utils.GenerateVerificationCode()
	if err != nil {
		return fmt.Errorf("failed to generate verification code: %w", err)
	}
	verification := model.EmailVerification{
		UserID:       userID,
		Email:        email,
		Code:         code,
		ExpiresAt:    time.Now().Add(15 * time.Minute),
		Used:         false,
		AttemptCount: 0,
	}
	if err := tx.Create(&verification).Error; err != nil {
		return fmt.Errorf("failed to save verification record: %w", err)
	}
	subject := "邮箱验证码 - Imperishable Gate"
	body := utils.GetVerificationEmailTemplate(code)
	if err := utils.SendEmail(email, subject, body); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}

// SendVerificationEmail 发送验证邮件
func SendVerificationEmail(userID uint, email string) error {
	// 非事务场景下直接用 database.DB
	return sendVerificationCodeAndSaveRecord(database.DB, userID, email)
}

// ResendVerificationEmail 重新发送验证邮件（优化版：2分钟冷却时间）
func ResendVerificationEmail(email string) error {
	// 查找用户
	var user model.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return ErrUserNotFound
	}

	// 检查是否已验证
	if user.EmailVerified {
		return ErrEmailAlreadyVerified
	}

	// 检查是否存在未过期的验证码
	var activeVerification model.EmailVerification
	err := database.DB.Where("user_id = ? AND used = ? AND expires_at > ?", user.ID, false, time.Now()).
		Order("created_at DESC").
		First(&activeVerification).Error

	if err == nil {
		// 存在未过期的验证码，检查是否在2分钟冷却期内
		timeSinceCreation := time.Since(activeVerification.CreatedAt)
		if timeSinceCreation < 2*time.Minute {
			// 还在冷却期内，不允许重发
			return ErrResendTooSoon
		}
		// 已超过2分钟，允许重发，标记旧验证码为已使用
		activeVerification.Used = true
		database.DB.Save(&activeVerification)
	}

	// 清理该用户所有已过期或已使用的验证码
	_ = database.DB.Model(&model.EmailVerification{}).
		Where("user_id = ? AND (used = ? OR expires_at <= ?)", user.ID, true, time.Now()).
		Delete(&model.EmailVerification{}).Error

	// 发送新的验证码
	err = SendVerificationEmail(user.ID, user.Email)
	return err
}
