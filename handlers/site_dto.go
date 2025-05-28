package handlers

type SiteDTO struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ProjectID   uint       `json:"project_id"`
	Project     ProjectDTO `json:"project"`
	SubDomain   *string    `json:"sub_domain"`
	Domains     []string   `json:"domains"`
}

type CreateSiteReq struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	ProjectID   uint     `json:"project_id" binding:"required"`
	SubDomain   *string  `json:"sub_domain"`
	Domains     []string `json:"domains"`
}
