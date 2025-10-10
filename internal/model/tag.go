package model

type Tag struct {
	ID     uint   `gorm:"primarykey" json:"id"`
	UserID uint   `gorm:"not null;uniqueIndex:tag_userid_idx" json:"user_id"`
	Name   string `gorm:"not null;uniqueIndex:tag_userid_idx;size:32" json:"name"`
	User   User   `gorm:"foreignKey:UserID;" json:"user"`
	Links  []Link `gorm:"many2many:link_tags;joinForeignKey:TagID;joinReferences:LinkID" json:"links"`
}
