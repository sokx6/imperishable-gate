package email

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/service/common"
	"time"
)

// ResendVerificationEmail 重新发送验证邮件（优化版：2分钟冷却时间）
func ResendVerificationEmail(email string) error {
	// 查找用户
	var user model.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return common.ErrUserNotFound
	}

	// 检查是否已验证
	if user.EmailVerified {
		return common.ErrEmailAlreadyVerified
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
			return common.ErrResendTooSoon
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
