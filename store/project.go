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

// GetByID 通过项目ID获取项目
func (p *projectType) GetByID(id uint) (project *models.Project, err error) {
	err = p.db.First(&project, id).Preload("Owners").Error
	return
}

// UserIsOwner 判断用户是否是项目的所有者
func (p *projectType) UserIsOwner(project *models.Project, userID uint) bool {
	if project.OwnerType == constants.OwnerTypeUser && project.OwnerID == userID {
		return true
	}
	for _, owner := range project.Owners {
		if owner.ID == userID {
			return true
		}
	}
	return false
}

// ListByOwner 通过用户ID获取项目列表，支持分页和从新到旧排序
func (p *projectType) ListByOwner(ownerType, ownerID string, page, limit int) (projects []models.Project, total int64, err error) {
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

	projects, total, err = Paginate[models.Project](
		p.db,
		page,
		limit,
		"owner_type = ? AND owner_id = ?",
		tableName,
		ownerID,
	)
	return
}

// Update 更新项目
func (p *projectType) Update(project *models.Project) (err error) {
	return p.db.Updates(project).Error
}

// Delete 删除项目
func (p *projectType) Delete(project *models.Project) (err error) {
	return p.db.Delete(project).Error
}
