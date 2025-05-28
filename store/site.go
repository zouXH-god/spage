package store

import (
	"github.com/LiteyukiStudio/spage/models"
	"gorm.io/gorm"
)

type SiteType struct {
	db *gorm.DB
}

var Site = SiteType{
	db: DB,
}

func (s *SiteType) Create(site *models.Site) (err error) {
	return s.db.Create(site).Error
}

func (s *SiteType) GetByID(id uint) (site *models.Site, err error) {
	site = &models.Site{}
	err = s.db.Where("id = ?", id).First(site).Error
	return
}
