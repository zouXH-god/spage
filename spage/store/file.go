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

func (f *FileType) GetByHash(hash string) (file models.File, err error) {
	err = DB.Where("hash = ?", hash).First(&file).Error
	return
}

func (f *FileType) GetByID(id uint) (file models.File, err error) {
	err = DB.Where("id = ?", id).First(&file).Error
	return
}
