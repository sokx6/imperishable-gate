package model

type LinkTag struct {
	LinkID uint `gorm:"primaryKey"`
	Link   Link `gorm:"constraint:OnDelete:CASCADE;foreignKey:LinkID"`
	TagID  uint `gorm:"primaryKey"`
	Tag    Tag  `gorm:"constraint:OnDelete:CASCADE;foreignKey:TagID"`
}
