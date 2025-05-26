package handlers

type OrganizationDTO struct {
	ID           uint      `json:"id"` // 组织ID
	Name         string    `json:"name"`
	DisplayName  *string   `json:"display_name"`
	Email        *string   `json:"email"`
	Description  string    `json:"description"`
	AvatarURL    *string   `json:"avatar_url"`
	ProjectLimit int       `json:"project_limit"`
	Members      []UserDTO `json:"members"` // 组织成员
	Owners       []UserDTO `json:"owners"`  // 组织所有者
}

type CreateOrgReq struct {
	Name        string  `json:"name" binding:"required"`
	DisplayName string  `json:"display_name" binding:"required"`
	Email       *string `json:"email"`
	Description string  `json:"description" binding:"required"`
	AvatarURL   *string `json:"avatar_url"`
}

type UpdateOrgReq struct {
	DisplayName *string `json:"display_name"`
	Email       *string `json:"email"`
	Description *string `json:"description"`
	AvatarURL   *string `json:"avatar_url"`
}

type GetOrgProjectReq struct {
	Page    int `json:"page" binding:"required"`
	Limit   int `json:"limit" binding:"required"`
	OrderBy string
}

type OrgUserReq struct {
	UserID uint   `json:"user_id" binding:"required"`
	Role   string `json:"role" binding:"required"`
}
