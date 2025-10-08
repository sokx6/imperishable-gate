package email

import (
	"errors"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"

	"imperishable-gate/internal/server/service/common"

	"gorm.io/gorm"
)

func VerifyEmailByUsernameAndResetPassword(username, code, newPassword string) error {
	// 根据用户名查找用户邮箱
	var user model.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.ErrUserNotFound
		}
		return common.ErrDatabase
	}
	return VerifyEmailAndResetPassword(user.Email, code, newPassword)
}
