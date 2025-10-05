// authenticate_user.go
package service

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AuthenticateUser(username, password string) (bool, error) {
	var user model.User

	// 查询数据库中是否存在该用户名的用户
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, ErrUsernameNotFound
		}
		return false, ErrDatabase // 数据库错误
	}

	// 使用 bcrypt 验证密码是否匹配
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// 密码不匹配
		return false, nil
	}

	// 用户存在且密码正确
	return true, nil
}
