package handlers

// ProjectDTO 项目信息数据传输对象
// Project Information Data Transfer Object (DTO)
type ProjectDTO struct {
	ID          uint      `json:"id"`           // 项目ID Project ID
	Name        string    `json:"name"`         // 项目名称 Project Name
	DisplayName *string   `json:"display_name"` // 项目显示名称 Project Display Name
	Description string    `json:"description"`  // 项目描述 Project Description
	OwnerType   string    `json:"owner_type"`   // 项目拥有者类型 Project Owner Type
	OwnerID     uint      `json:"owner_id"`     // 项目拥有者ID Project Owner ID
	Owners      []UserDTO `json:"owners"`       // 项目拥有者列表 Project Owner List
	SiteLimit   int       `json:"site_limit"`   // 项目站点数量限制 Project Site Limit
}

// CreateProjectReq 创建项目请求参数
// Create Project Request Parameters
type CreateProjectReq struct {
	Name        string  `json:"name" binding:"required"`                                         // 项目名称 Project Name
	DisplayName *string `json:"display_name"`                                                    // 项目显示名称 Project Display Name
	Description string  `json:"description"`                                                     // 项目描述 Project Description
	OwnerType   string  `json:"owner_type" binding:"required"  vd:"in($,'user','organization')"` // 项目拥有者类型 Project Owner Type
	OwnerID     uint    `json:"owner_id" binding:"required"`                                     // 项目拥有者ID Project Owner ID

}

// UpdateProjectReq 更新项目请求参数
// Update Project Request Parameters
type UpdateProjectReq struct {
	Name        *string `json:"name"`         // 项目名称 Project Name
	DisplayName *string `json:"display_name"` // 项目显示名称 Project Display Name
	Description *string `json:"description"`  // 项目描述 Project Description
}

// ProjectUserReq 项目用户请求参数
// Project User Request Parameters
type ProjectUserReq struct {
	UserID uint `json:"user_id" binding:"required"` // 用户ID User ID
}

// GetProjectListReq 获取项目列表请求参数
// Get Project List Request Parameters
type GetSiteListReq struct {
	Page     int    `form:"page" binding:"required"`    // 页码 Page
	Limit    int    `form:"limit" binding:"required"`   // 每页数量 Page Limit
	Project  string `form:"project" binding:"required"` // 项目名称 Project Name
	SiteName string `form:"site_name"`                  // 站点名称 Site Name
}
