package store

import (
	"github.com/LiteyukiStudio/spage/constants"
	"github.com/LiteyukiStudio/spage/spage/models"
)

type SiteType struct {
}

var Site = SiteType{}

// Create 创建站点
// Create Site
func (s *SiteType) Create(site *models.Site) (err error) {
	return DB.Create(site).Error
}

// GetByID 根据id获取站点信息
// Get Site Info by ID
func (s *SiteType) GetByID(id uint) (site *models.Site, err error) {
	site = &models.Site{}
	err = DB.Where("id = ?", id).Preload(constants.PreloadFieldProject).First(site).Error
	return
}

func (s *SiteType) Update(site *models.Site) (err error) {
	return DB.Updates(site).Error
}

func (s *SiteType) Delete(site *models.Site) (err error) {
	return DB.Delete(site).Error
}

func (s *SiteType) GetReleaseList(siteID uint) (releases []*models.SiteRelease, err error) {
	err = DB.Where("site_id = ?", siteID).Preload(constants.PreloadFieldFile).Find(&releases).Error
	return
}

func (s *SiteType) GetReleaseById(id uint) (release *models.SiteRelease, err error) {
	release = &models.SiteRelease{}
	err = DB.Where("id = ?", id).Preload(constants.PreloadFieldFile).First(release).Error
	return
}

func (s *SiteType) GetLatestRelease(site *models.Site) (release *models.SiteRelease, err error) {
	release = &models.SiteRelease{}
	err = DB.Where("site_id = ? AND tag = ?", site.ID, "latest").Preload(constants.PreloadFieldFile).First(release).Error
	return
}

func (s *SiteType) CreateRelease(release *models.SiteRelease) (err error) {
	return DB.Create(release).Error
}

func (s *SiteType) DeleteRelease(release *models.SiteRelease) (err error) {
	return DB.Delete(release).Error
}

func (s *SiteType) UpdateRelease(release *models.SiteRelease) (err error) {
	return DB.Updates(release).Error
}
