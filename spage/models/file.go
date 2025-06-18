package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey" json:"id"` // 文件ID
	Path string `gorm:"not null" json:"path"` // 文件路径，相较于根目录的相对路径
	MD5  string `gorm:"not null" json:"md5"`  // 文件哈希值
}

// TableName 自定义表名
func (File) TableName() string {
	return "projects"
}
