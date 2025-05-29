package store

import (
	"github.com/LiteyukiStudio/spage/models"
	"gorm.io/gorm"
)

type FileType struct {
	db *gorm.DB
}

var File = FileType{
	db: DB,
}

func (f *FileType) Create(file *models.File) (err error) {
	return f.db.Create(file).Error
}
