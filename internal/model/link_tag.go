package model

type LinkTag struct {
	LinkID uint `gorm:"primaryKey"`
	TagID  uint `gorm:"primaryKey"`
}
