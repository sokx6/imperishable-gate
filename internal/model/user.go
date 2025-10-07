package model

type User struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Username string `gorm:"not null;uniqueIndex;size:40" json:"username"`
	Password string `gorm:"not null;size:64" json:"password"`
	Email    string `gorm:"not null;uniqueIndex" json:"email"`
	Links    []Link `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"links"`
	Tags     []Tag  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"tags"`
}
