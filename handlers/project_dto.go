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
