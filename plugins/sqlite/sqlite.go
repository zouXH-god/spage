package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

// InitSQLite 初始化 SQLite 连接
func InitSQLite(path string, gormConfig *gorm.Config) (*gorm.DB, error) {
	if path == "" {
		path = "./data/data.db"
	}
	// 创建 SQLite 数据库文件的目录
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory for SQLite database: %w", err)
	}

	db, err := gorm.Open(sqlite.Open(path), gormConfig)
	return db, err
}
