package model

import (
	"github.com/lib/pq"
)

// Link 是数据库中的链接记录
type Link struct {
	Id   uint           `gorm:"primaryKey"`
	Url  string         `gorm:"not null;uniqueIndex;size:2048"` // 唯一索引并且限制长度
	Tags pq.StringArray `gorm:"type:text[]"`                    // PostgreSQL 的文本数组类型                    // PostgreSQL 的文本数组类型
}

// TableName 设置表名为 'links'
func (Link) TableName() string {
	return "links"
}
