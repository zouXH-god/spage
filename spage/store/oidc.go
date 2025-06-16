package store

import "github.com/LiteyukiStudio/spage/spage/models"

type oidcType struct{}

var Oidc = oidcType{}

func (oidcType) ListEnabledOidcConfig() (configs []models.OIDCConfig, err error) {
	err = DB.Where("enabled = true").Find(&configs).Error
	return
}

func (oidcType) CreateOidcConfig(oidcConfig *models.OIDCConfig) error {
	return DB.Create(oidcConfig).Error
}

func (oidcType) UpdateOidcConfig(oidcConfig *models.OIDCConfig) error {
	return DB.Updates(oidcConfig).Error
}

func (oidcType) DeleteOidcConfig(oidcConfig *models.OIDCConfig) error {
	return DB.Delete(oidcConfig).Error
}
