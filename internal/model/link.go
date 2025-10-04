package model

type Link struct {
	ID     uint   `gorm:"primarykey" json:"id"`
	Url    string `gorm:"not null;uniqueIndex" json:"url"`
	Tags   []Tag  `gorm:"many2many:link_tags;joinForeignKey:LinkID;joinReferences:TagID" json:"tags"`
	Names  []Name `gorm:"foreignKey:LinkID;constraint:OnDelete:CASCADE;" json:"names"` // 注意字段名用复数更合适！
	Remark string `gorm:"size:128" json:"remark"`
}
