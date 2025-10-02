// model/tag.go
package model

type Tag struct {
	ID   uint   `gorm:"primarykey;autoIncrement" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`

	Links []Link `gorm:"many2many:link_tags;"` // 多对多关系
}
