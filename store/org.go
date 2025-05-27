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

// GetOrgById 通过ID获取组织
func (o *orgType) GetOrgById(id uint) (org *models.Organization, err error) {
	err = o.db.Model(&models.Organization{}).Where("id = ?", id).First(&org).Error
	return
}

// LoadOrgUsers 加载组织成员
func (o *orgType) LoadOrgUsers(org *models.Organization) error {
	return o.db.Model(org).Preload("Members").Preload("Owners").Find(org).Error
}

// OrgNameIsExist 判断组织名称是否存在
func (o *orgType) OrgNameIsExist(name string) bool {
	var count int64
	o.db.Model(&models.Organization{}).Where("name = ?", name).Count(&count)
	return count > 0
}

// CreateOrg 创建组织
func (o *orgType) CreateOrg(org *models.Organization) error {
	return o.db.Create(org).Error
}

// UpdateOrg 更新组织
func (o *orgType) UpdateOrg(org *models.Organization) error {
	return o.db.Updates(org).Error
}

// DeleteOrg 删除组织
func (o *orgType) DeleteOrg(org *models.Organization) error {
	return o.db.Model(org).Delete(org).Error
}
