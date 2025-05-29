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

// Create 创建站点
// Create Site
func (s *SiteType) Create(site *models.Site) (err error) {
	return s.db.Create(site).Error
}

// GetByID 根据id获取站点信息
// Get Site Info by ID
func (s *SiteType) GetByID(id uint) (site *models.Site, err error) {
	site = &models.Site{}
	err = s.db.Where("id = ?", id).Preload("Project").First(site).Error
	return
}

func (s *SiteType) Update(site *models.Site) (err error) {
	return s.db.Updates(site).Error
}

func (s *SiteType) Delete(site *models.Site) (err error) {
	return s.db.Delete(site).Error
}

func (s *SiteType) GetReleaseList(siteID uint) (releases []*models.SiteRelease, err error) {
	err = s.db.Where("site_id = ?", siteID).Find(&releases).Error
	return
}

func (s *SiteType) GetReleaseById(id uint) (release *models.SiteRelease, err error) {
	release = &models.SiteRelease{}
	err = s.db.Where("id = ?", id).Preload("File").First(release).Error
	return
}

func (s *SiteType) CreateRelease(release *models.SiteRelease) (err error) {
	return s.db.Create(release).Error
}

func (s *SiteType) DeleteRelease(release *models.SiteRelease) (err error) {
	return s.db.Delete(release).Error
}
