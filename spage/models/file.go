package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"` // 文件ID File ID
	Hash       string `gorm:"not null"`   // 文件哈希值 File hash
	UploaderID uint   `gorm:"not null"`   // 上传者ID Uploader ID
}

// TableName 自定义表名 Custom table name
func (File) TableName() string {
	return "files" // 表名为 files
}
