package models

import (
	"gorm.io/gorm"
)

type Site struct {
	gorm.Model
	Name        string   `gorm:"unique"`                                                            // 站点名称 Site name
	Description string   `gorm:"size:255"`                                                          // 站点描述 Site description
	ProjectID   uint     `gorm:"not null"`                                                          // 项目ID Project ID
	Project     Project  `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // 项目 Project
	SubDomain   string   `gorm:"unique;size:255"`                                                   // 子域前缀 Subdomain prefix
	Domains     []string `gorm:"type:json;default:'[]'"`                                            // 允许的域名，json格式 Allowed domains, json format
}

func (Site) TableName() string {
	return "sites"
}

type SiteRelease struct {
	gorm.Model
	SiteID uint   `gorm:"not null"` // 站点ID Site ID
	Site   Site   `gorm:"foreignKey:SiteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tag    string `gorm:"not null"`                                                        // 版本标签 Version tag
	FileID uint   `gorm:"not null"`                                                        // 版本文件ID Version file ID
	File   File   `gorm:"foreignKey:FileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // 版本文件 Version file
}

func (SiteRelease) TableName() string {
	return "site_releases"
}
