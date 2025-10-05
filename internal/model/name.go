package model

type Name struct {
	ID     uint   `gorm:"primarykey" json:"id"`
	UserID uint   `gorm:"not null;uniqueIndex:idx_user_name" json:"user_id"`
	Name   string `gorm:"not null;uniqueIndex:idx_user_name;size:32" json:"name"`
	LinkID uint   `gorm:"not null" json:"link_id"`
	Link   Link   `gorm:"foreignKey:LinkID;" json:"link"`
}
