package models

import "gorm.io/gorm"

// User Model
type User struct {
	gorm.Model
	Name          string          `gorm:"not null;unique"`                 // Unique name for the user
	DisplayName   *string         `gorm:"column:display_name"`             // Display name for the user
	Email         *string         `gorm:"unique"`                          // Email address for the user, only user's email is unique (for oidc authentication)
	Description   string          `gorm:"default:'No description.'"`       // Description of the user
	Avatar        *string         `gorm:"column:avatar"`                   // Leave blank to use Gravatar
	Role          string          `gorm:"not null;default:member"`         // Global role of the user
	Organizations []*Organization `gorm:"many2many:organization_members;"` // Belongs to many organizations
	ProjectLimit  int             `gorm:"default:0"`                       // Project limit for the user, 0 means unlimited
	Language      string          `gorm:"default:'zh-cn'"`                 // Language for the user, default is English
	Flag          string          `gorm:"default:'0'"`                     // another flag for the system_admin
	Password      *string         `gorm:"column:password"`                 // Password (hashed) for the user, only used for local authentication
}

func (User) TableName() string {
	return "users"
}

// Organization Model
type Organization struct {
	gorm.Model
	Name         string  `gorm:"not null;unique"`                 // Unique name for the organization
	DisplayName  *string `gorm:"column:display_name"`             // Display name for the organization
	Email        *string `gorm:"column:email"`                    // Email address for the organizations
	Description  string  `gorm:"default:'No description.'"`       // Description of the organization
	Avatar       *string `gorm:"column:avatar"`                   // Leave blank to use Gravatar
	Members      []*User `gorm:"many2many:organization_members;"` // Members of the organization, contains the creator
	Owners       []User  `gorm:"many2many:organization_owners;"`  // Owners of the organization, no reverse relation, contains the creator
	ProjectLimit int     `gorm:"default:0"`                       // Project limit for the organization, 0: follow the policy, -1: unlimited
}

func (Organization) TableName() string {
	return "organizations"
}

// Project Model
type Project struct {
	gorm.Model
	Name        string  `gorm:"not null;unique"`           // Unique name for the project
	DisplayName *string `gorm:"column:display_name"`       // Display name for the project
	Description string  `gorm:"default:'No description.'"` // Description of the project
	OwnerID     uint    `gorm:"not null"`                  // Owner ID (user id or organization id)
	OwnerType   string  `gorm:"not null"`                  // Owner type, can be user or organization
	Owners      []User  `gorm:"many2many:project_owners;"` // Owners of the project, no reverse relation
	SiteLimit   int     `gorm:"default:0"`                 // Site limit for the project, 0: follow the policy, -1: unlimited
}

func (Project) TableName() string {
	return "projects"
}

//
