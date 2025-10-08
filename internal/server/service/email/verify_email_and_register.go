package email

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/service/common"
	"time"
)

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
		return common.ErrDatabase
	}
	return nil
}
