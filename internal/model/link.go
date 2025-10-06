package model

type Link struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	UserID      uint   `gorm:"not null;uniqueIndex:idx_url_user" json:"user_id"`
	Url         string `gorm:"not null;uniqueIndex:idx_url_user" json:"url"`
	Tags        []Tag  `gorm:"many2many:link_tags;joinForeignKey:LinkID;joinReferences:TagID" json:"tags"`
	Names       []Name `gorm:"foreignKey:LinkID;constraint:OnDelete:CASCADE;" json:"names"` // 注意字段名用复数更合适！
	Remark      string `gorm:"size:128" json:"remark"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Watching    bool   `gorm:"default:false" json:"watching"`
	StatusCode  int    `gorm:"default:200" json:"status_code"`
}
