package models

import "gorm.io/gorm"

// Migrate 迁移模型，通过依赖注入的方式，使用gorm.DB进行数据库操作
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&User{},
		&Organization{},
		&Project{},
		&OIDCConfig{},
	); err != nil {
		return err
	}
	return nil
}
