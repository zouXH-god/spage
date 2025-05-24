package store

import (
	"fmt"
	"github.com/LiteyukiStudio/spage/constants"
	"github.com/LiteyukiStudio/spage/models"
	"gorm.io/gorm"
)

type projectType struct {
	db *gorm.DB
}

var Project = projectType{
	db: DB,
}

// Create 创建项目
func (p *projectType) Create(project *models.Project) (err error) {
	return p.db.Create(project).Error
}

// ListByOwner 通过用户ID获取项目列表，支持分页和从新到旧排序
func (p *projectType) ListByOwner(ownerType, ownerID string, page, limit int) (projects []models.Project, err error) {
	tableName := ""
	switch ownerType {
	case constants.OwnerTypeUser:
		tableName = models.User{}.TableName()
	case constants.OwnerTypeOrg:
		tableName = models.Organization{}.TableName()
	default:
		err = fmt.Errorf("invalid owner type")
		return
	}

	projects, _, err = Paginate[models.Project](
		p.db,
		page,
		limit,
		"owner_type = ? AND owner_id = ?",
		tableName,
		ownerID,
	)
	return
}
