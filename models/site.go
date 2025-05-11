package models

import (
	"gorm.io/gorm"
)

type Site struct {
	gorm.Model
	Name        string   `gorm:"unique"`
	Description string   `gorm:"size:255"`
	ProjectID   uint     `gorm:"not null"`
	Project     Project  `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	SubDomain   string   `gorm:"unique;size:255"`        // 子域前缀
	Domains     []string `gorm:"type:json;default:'[]'"` // 允许的域名，json格式
}

func (Site) TableName() string {
	return "sites"
}

type SiteRelease struct {
	gorm.Model
	SiteID uint   `gorm:"not null"`
	Site   Site   `gorm:"foreignKey:SiteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tag    string `gorm:"not null"`                                                        // 版本标签
	FileID uint   `gorm:"not null"`                                                        // 版本文件ID
	File   File   `gorm:"foreignKey:FileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // 版本文件
}

func (SiteRelease) TableName() string {
	return "site_releases"
}
