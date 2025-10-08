package model

import "time"

// EmailVerification 邮箱验证记录
type EmailVerification struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	UserID       uint      `gorm:"not null;index" json:"user_id"`
	Email        string    `gorm:"not null" json:"email"`
	Code         string    `gorm:"not null;size:6" json:"code"`
	ExpiresAt    time.Time `gorm:"not null" json:"expires_at"`
	Used         bool      `gorm:"default:false" json:"used"`
	AttemptCount int       `gorm:"default:0" json:"attempt_count"` // 验证尝试次数（防止暴力破解）
	CreatedAt    time.Time `json:"created_at"`
	User         User      `gorm:"foreignKey:UserID" json:"user"`
}
