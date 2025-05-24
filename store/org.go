package store

import (
	"github.com/LiteyukiStudio/spage/models"
	"gorm.io/gorm"
)

type orgType struct {
	db *gorm.DB
}

var Org = orgType{
	db: DB,
}

// ListByUserID 通过UserID获取用户组织，支持分页和预加载关系
func (o *orgType) ListByUserID(userID string, page, limit int) (orgs []models.Organization, err error) {
	// 使用连接查询
	query := o.db.Joins("JOIN organization_members ON organizations.id = organization_members.organization_id").
		Where("organization_members.user_id = ?", userID)
	// 预加载关系
	query = WithPreloads(query, "Members", "Owners")
	// 使用通用分页方法
	orgs, _, err = Paginate[models.Organization](
		query,
		page,
		limit,
	)

	return
}
