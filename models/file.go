package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"` // 文件ID
	Path string `gorm:"not null"`   // 文件路径，相较于根目录的相对路径
}

func (File) TableName() string {
	return "projects"
}
