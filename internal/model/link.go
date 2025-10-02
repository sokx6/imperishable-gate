// model/link.go 修改版
package model

type Link struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Url  string `gorm:"not null;uniqueIndex" json:"url"`
	Tags []Tag  `gorm:"many2many:link_tags;" json:"tags"` // 改为关联 Tags
	Note string `json:"note"`                             // 备注字段新增
}
