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
