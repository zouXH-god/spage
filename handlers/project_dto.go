package handlers

type ProjectDTO struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	DisplayName *string   `json:"display_name"`
	Description string    `json:"description"`
	OwnerType   string    `json:"owner_type"`
	OwnerID     uint      `json:"owner_id"`
	Owners      []UserDTO `json:"owners"`
	SiteLimit   int       `json:"site_limit"`
}

type CreateProjectReq struct {
	Name        string  `json:"name" binding:"required"`
	DisplayName *string `json:"display_name"`
	Description string  `json:"description"`
	OwnerType   string  `json:"owner_type" binding:"required"  vd:"in($,'user','organization')"`
	OwnerID     uint    `json:"owner_id" binding:"required"`
}

type UpdateProjectReq struct {
	Name        *string `json:"name"`
	DisplayName *string `json:"display_name"`
	Description *string `json:"description"`
}

type ProjectUserReq struct {
	UserID uint `json:"user_id" binding:"required"`
}
