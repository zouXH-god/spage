package models

import "gorm.io/gorm"

// User 用户模型
type User struct {
	gorm.Model
	Name          string          `gorm:"not null;unique"`                 // 用户的唯一名称
	DisplayName   *string         `gorm:"column:display_name"`             // 用户的显示名称
	Email         *string         `gorm:"unique"`                          // 用户的电子邮件地址，只有用户的电子邮件地址是唯一的（用于 oidc 身份验证）
	Description   string          `gorm:"default:'No description.'"`       // 用户描述
	AvatarURL     *string         `gorm:"column:avatar_url"`               // 留空以使用
	Role          string          `gorm:"not null;default:member"`         // 用户的全局角色
	Organizations []*Organization `gorm:"many2many:organization_members;"` // 隶属于许多组织
	ProjectLimit  int             `gorm:"default:-1"`                      // 用户的项目限制，0 表示无限制
	Language      string          `gorm:"default:'zh-cn'"`                 // 用户的语言，默认为英语
	Flag          string          `gorm:"default:'0'"`                     // system_admin 的另一面旗帜
	Password      *string         `gorm:"column:password"`                 // 用户的密码（经过哈希处理），仅用于本地身份验证
}

// TableName 用户
func (User) TableName() string {
	return "users"
}

// Organization 组织模型
type Organization struct {
	gorm.Model
	Name         string  `gorm:"not null;unique"`                 // 组织的唯一名称
	DisplayName  *string `gorm:"column:display_name"`             // 组织的显示名称
	Email        *string `gorm:"column:email"`                    // 组织的电子邮件地址
	Description  string  `gorm:"default:'No description.'"`       // 组织描述
	AvatarURL    *string `gorm:"column:avatar_url"`               // 留空以使用
	Members      []*User `gorm:"many2many:organization_members;"` // 组织的成员包含创建者
	Owners       []User  `gorm:"many2many:organization_owners;"`  // 组织的所有者（无反向关系）包含创建者
	ProjectLimit int     `gorm:"default:0"`                       // 组织的项目限制，0：遵循策略，-1：无限制
}

// TableName 组织
func (Organization) TableName() string {
	return "organizations"
}

// Project 项目模型
type Project struct {
	gorm.Model
	Name        string  `gorm:"not null;unique"`            // 项目的唯一名称
	DisplayName *string `gorm:"column:display_name"`        // 项目的显示名称
	Description string  `gorm:"default:'No description.'"`  // 项目描述
	OwnerID     uint    `gorm:"not null"`                   // 所有者 ID（用户 ID 或组织 ID）
	OwnerType   string  `gorm:"not null"`                   // 所有者类型，可以是用户或组织
	Owners      []User  `gorm:"many2many:project_owners;"`  // 项目的所有者，无反向关系
	Members     []*User `gorm:"many2many:project_members;"` // 项目的成员
	SiteLimit   int     `gorm:"default:0"`                  // 项目的站点限制，0：遵循策略，-1：无限制
}

// TableName 项目
func (Project) TableName() string {
	return "projects"
}
