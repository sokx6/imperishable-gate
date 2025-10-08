package email

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/utils"

	"imperishable-gate/internal/server/service/common"
)

func VerifyEmailAndResetPassword(email, code, newPassword string) error {
	// 验证验证码并且获取用户ID
	userId, err := VerifyEmailByCode(email, code)
	if err != nil {
		return err
	}
	// 更新用户密码
	newPassword, err = utils.HashPassword(newPassword)
	if err != nil {
		return common.ErrDatabase
	}
	if err := database.DB.Model(&model.User{}).
		Where("id = ?", userId).
		Updates(map[string]interface{}{
			"password": newPassword,
		}).Error; err != nil {
		return common.ErrDatabase
	}
	return nil
}
