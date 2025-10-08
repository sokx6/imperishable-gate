package model

import "time"

type User struct {
	ID              uint       `gorm:"primarykey" json:"id"`
	Username        string     `gorm:"not null;uniqueIndex;size:40" json:"username"`
	Password        string     `gorm:"not null;" json:"password"`
	Email           string     `gorm:"not null;uniqueIndex" json:"email"`
	EmailVerified   bool       `gorm:"default:false" json:"email_verified"`
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty"`
	Links           []Link     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"links"`
	Tags            []Tag      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"tags"`
	Names           []Name     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"names"`
}
