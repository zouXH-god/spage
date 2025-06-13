package store

import (
	"github.com/LiteyukiStudio/spage/spage/config"
	"gorm.io/gorm"
)

// Paginate 封装通用的分页查询逻辑
// T 是任意数据模型类型
// Paginate generic pagination query logic
// T is any data model type
func Paginate[T any](db *gorm.DB, page, limit int, conditions ...any) (items []T, total int64, err error) {
	// 查询总记录数
	countDB := db
	if len(conditions) > 0 {
		countDB = countDB.Where(conditions[0], conditions[1:]...)
	}
	err = countDB.Model(new(T)).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 确保分页参数有效
	// Ensure pagination parameters are valid
	if limit <= 0 {
		limit = config.PageLimit
	}

	// 执行分页查询
	// Execute pagination query
	queryDB := db
	if len(conditions) > 0 {
		queryDB = queryDB.Where(conditions[0], conditions[1:]...)
	}

	if page > 0 {
		offset := (page - 1) * limit
		queryDB = queryDB.Offset(offset)
	}

	err = queryDB.Limit(limit).Order("id DESC").Find(&items).Error
	return
}

// WithPreloads 添加预加载关系的辅助函数
// Add a helper function to add preloaded relationships
func WithPreloads(db *gorm.DB, preloads ...string) *gorm.DB {
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	return db
}
