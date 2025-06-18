package handlers

// SiteDTO 网站详情
// Site Detail
type SiteDTO struct {
	ID          uint       `json:"id"`          // 网站ID WebSiteID
	Name        string     `json:"name"`        // 网站名称 WebSiteName
	Description string     `json:"description"` // 网站描述 WebSiteDescription
	ProjectID   uint       `json:"project_id"`  // 项目ID ProjectID
	Project     ProjectDTO `json:"project"`     // 项目详情 ProjectDetail
	SubDomain   *string    `json:"sub_domain"`  // 子域名 SubDomain
	Domains     []string   `json:"domains"`     // 域名 Domains
}

// CreateSiteReq 创建网站请求参数
// Create Site Request Parameters
type CreateSiteReq struct {
	Name        string   `json:"name" binding:"required"`       // 网站名称 WebSiteName
	Description string   `json:"description"`                   // 网站描述 WebSiteDescription
	ProjectID   uint     `json:"project_id" binding:"required"` // 项目ID ProjectID
	SubDomain   *string  `json:"sub_domain"`                    // 子域名 SubDomain
	Domains     []string `json:"domains"`                       // 域名 Domains
}

type UpdateSiteReq struct {
	Name        *string  `json:"name"`        // 网站名称 WebSiteName
	Description *string  `json:"description"` // 网站描述 WebSiteDescription
	SubDomain   *string  `json:"sub_domain"`  // 子域名 SubDomain
	Domains     []string `json:"domains"`     // 域名 Domains
}
