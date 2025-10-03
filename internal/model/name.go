package model

type Name struct {
	ID     uint   `gorm:"primarykey" json:"id"`
	Name   string `gorm:"not null;uniqueIndex" json:"name"`
	LinkID uint   `gorm:"not null" json:"link_id"`
	Link   Link   `gorm:"foreignKey:LinkID;" json:"link"`
}
