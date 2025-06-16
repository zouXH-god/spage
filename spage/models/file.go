package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey" json:"id"` // 文件ID File ID
	Path string `gorm:"not null" json:"path"` // 文件路径，相较于根目录的相对路径 File path, relative to the root directory
	MD5  string `gorm:"not null" json:"md5"`  // 文件哈希值 File hash
}

// TableName 自定义表名 Custom table name
func (File) TableName() string {
	return "projects"
}
