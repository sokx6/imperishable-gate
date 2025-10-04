package model

type Name struct {
	ID     uint   `gorm:"primarykey" json:"id"`
	Name   string `gorm:"not null;uniqueIndex;size:32" json:"name"`
	LinkID uint   `gorm:"not null" json:"link_id"`
	Link   Link   `gorm:"foreignKey:LinkID;" json:"link"`
}
