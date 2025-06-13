package store

import (
	"github.com/LiteyukiStudio/spage/spage/models"
)

type FileType struct {
}

var File = FileType{}

func (f *FileType) Create(file *models.File) (err error) {
	return DB.Create(file).Error
}
