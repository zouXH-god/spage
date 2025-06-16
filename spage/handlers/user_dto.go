package handlers

// RegisterReq 注册请求结构体
type RegisterReq struct {
	Username string `json:"username" binding:"required"` // 用户名 Username
	Password string `json:"password" binding:"required"` // 密码 Password
	Email    string `json:"email" binding:"required"`    // 邮箱 Email

}

// LoginReq 登录请求结构体
type LoginReq struct {
	Username     string `json:"username" binding:"required"`      // 用户名 Username
	Password     string `json:"password" binding:"required"`      // 密码 Password
	CaptchaToken string `json:"captcha_token" binding:"required"` // 验证码 Token
}

// CreateTokenReq 创建Token请求
type CreateTokenReq struct {
	Duration uint `json:"duration" binding:"required"`
}

// UserDTO 组织信息数据传输对象
type UserDTO struct {
	ID            uint              `json:"id"`            // 用户ID User ID
	Name          string            `json:"name"`          // 用户名 Username
	DisplayName   *string           `json:"display_name"`  // 显示名称 DisplayName
	Email         *string           `json:"email"`         // 邮箱 Email
	Description   string            `json:"description"`   // 描述 Description
	Avatar        *string           `json:"avatar_url"`    // 头像 Avatar URL
	Role          string            `json:"role"`          // 角色 Role
	Organizations []OrganizationDTO `json:"organizations"` // 组织 Organizations
	Language      string            `json:"language"`      // 语言 Language
	//Password      string            `json:"password"` // 密码 Password
}
