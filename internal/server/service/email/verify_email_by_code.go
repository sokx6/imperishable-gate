package email

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"time"

	"imperishable-gate/internal/server/service/common"

	"gorm.io/gorm"
)

// VerifyEmailByCode 验证邮箱失败时记录尝试次数
func VerifyEmailByCode(email, code string) (userId uint, err error) {
	var verification model.EmailVerification

	// 查找该邮箱的活跃验证记录
	err = database.DB.Where("email = ? AND used = ? AND expires_at > ?", email, false, time.Now()).
		First(&verification).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, common.ErrInvalidVerificationCode
	} else if err != nil {
		return 0, common.ErrDatabase
	}

	// 检查尝试次数
	if verification.AttemptCount >= 5 {
		verification.Used = true
		database.DB.Save(&verification)
		return 0, common.ErrTooManyAttempts
	}

	// 检查验证码是否正确
	if verification.Code != code {
		// 验证码错误，增加尝试次数
		verification.AttemptCount++
		if err := database.DB.Save(&verification).Error; err != nil {
			return 0, common.ErrDatabase
		}
		return 0, common.ErrInvalidVerificationCode
	}

	// 验证成功
	verification.Used = true
	if err := database.DB.Save(&verification).Error; err != nil {
		return 0, common.ErrDatabase
	}

	return verification.UserID, nil
}
