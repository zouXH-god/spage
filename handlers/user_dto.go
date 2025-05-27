package handlers

import "github.com/LiteyukiStudio/spage/models"

type RegisterReq struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
	Email    string `json:"email" binding:"required"`
}

type LoginReq struct {
	Username     string `json:"username" binding:"required"` // 用户名
	Password     string `json:"password" binding:"required"` // 密码
	CaptchaToken string `json:"captcha_token" binding:"required"`
}

type UserDTO struct {
	ID            uint              `json:"id"` // 用户ID
	Name          string            `json:"name"`
	DisplayName   *string           `json:"display_name"` // 显示名称
	Email         *string           `json:"email"`
	Description   string            `json:"description"`
	Avatar        *string           `json:"avatar_url"`
	Role          string            `json:"role"`
	Organizations []OrganizationDTO `json:"organizations"`
	Language      string            `json:"language"`
	//Password      string            `json:"password"`
}

func (u *UserDTO) FromUser(user *models.User) {
	u.ID = user.ID
	u.Name = user.Name
	u.DisplayName = user.DisplayName
	u.Email = user.Email
	u.Description = user.Description
	u.Avatar = user.AvatarURL
	u.Role = user.Role
	u.Language = user.Language
}
