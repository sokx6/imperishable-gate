// model/tag.go
package model

type Tag struct {
	ID    uint   `gorm:"primarykey;autoIncrement" json:"id"`
	Name  string `gorm:"uniqueIndex;not null;size:32" json:"name"`
	Links []Link `gorm:"many2many:link_tags;joinForeignKey:TagID;joinReferences:LinkID" json:"links"` // 多对多关系
}
