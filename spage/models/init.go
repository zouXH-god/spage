package models

import "gorm.io/gorm"

// Migrate 迁移模型，通过依赖注入的方式，使用gorm.DB进行数据库操作
// Migrate models, using gorm.DB for database operations through dependency injection
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		// entity.go
		&User{},
		&Organization{},
		&Project{},
		// file.go
		&File{},
		// token
		&JsonWebToken{},
		&ApiToken{},
		// oidc_config.go
		&OIDCConfig{},
		// site.go
		&Site{},
		&SiteRelease{},
		// node.go
		&Node{},
	); err != nil {
		return err
	}
	return nil
}
