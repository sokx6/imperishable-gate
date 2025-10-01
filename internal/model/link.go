package model

// Link 是数据库中的链接记录
type Link struct {
	ID  uint   `gorm:"primaryKey"`
	URL string `gorm:"not null;uniqueIndex;size:2048"` // 唯一索引并且限制长度
}

// TableName 设置表名为 'links'
func (Link) TableName() string {
	return "links"
}
