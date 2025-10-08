package email

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
