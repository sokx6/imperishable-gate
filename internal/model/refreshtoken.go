package model

import "time"

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;index"`
	Username  string    `gorm:"not null;index"`
	Token     string    `gorm:"not null;uniqueIndex"`
	ExpiresAt time.Time `gorm:"not null"`
	Revoked   bool      `gorm:"default:false"`
	CreatedAt time.Time
}
