package auth

import (
	"crypto/rand"
	"encoding/hex"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"time"
)

const RefreshExpiry = 7 * 24 * time.Hour // 7天

// GenerateRefreshToken 生成并存储刷新令牌
func GenerateRefreshToken(userID uint, userName string) (string, error) {
	// 生成随机字符串
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	token := hex.EncodeToString(bytes)

	// 保存到数据库
	refreshToken := model.RefreshToken{
		UserID:    userID,
		Username:  userName,
		Token:     token,
		ExpiresAt: time.Now().Add(RefreshExpiry),
		Revoked:   false,
	}

	if err := database.DB.Create(&refreshToken).Error; err != nil {
		return "", err
	}

	return token, nil
}
