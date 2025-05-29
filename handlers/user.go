package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/constants"
	"github.com/LiteyukiStudio/spage/middle"
	"github.com/LiteyukiStudio/spage/models"
	"github.com/LiteyukiStudio/spage/resps"
	"github.com/LiteyukiStudio/spage/store"
	"github.com/LiteyukiStudio/spage/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
)

type UserApi struct{}

var User = UserApi{}

// UserDTO 用户信息数据传输对象
// User Information Data Transfer Object (DTO)
func (UserApi) ToDTO(user *models.User, self bool) UserDTO {
	userDTO := UserDTO{
		ID:          user.ID,
		Name:        user.Name,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		Description: user.Description,
		Avatar:      user.AvatarURL,
	}
	if self {
		userDTO.Role = user.Role
		userDTO.Language = user.Language
	}
	return userDTO
}

// Login 用户登录
// User login
func (UserApi) Login(ctx context.Context, c *app.RequestContext) {
	loginReq := &LoginReq{}
	// TODO: 这里需要验证验证码
	err := c.BindJSON(loginReq)
	if err != nil {
		resps.BadRequest(c, "Parameter error")
		return
	}
	if loginReq.Username == "" || loginReq.Password == "" {
		resps.BadRequest(c, "Username or password cannot be empty")
		return
	}
	user, err := store.User.GetByName(loginReq.Username)
	if err != nil {
		user, err = store.User.GetByEmail(loginReq.Username)
		if err != nil {
			resps.BadRequest(c, "User does not exist")
			return
		}
	}

	if user.Password == nil {
		resps.Forbidden(c, "Password not set, please use another login method")
		return
	} else {
		if utils.Password.VerifyPassword(loginReq.Password, *user.Password, config.JwtSecret) {
			token, err := utils.Token.CreateToken(user.ID, time.Duration(config.TokenExpireTime)*time.Second, false, middle.PersistentHandler)
			if err != nil {
				resps.InternalServerError(c, "Failed to create token")
				return
			}
			refreshToken, err := utils.Token.CreateToken(user.ID, time.Duration(config.RefreshTokenExpireTime)*time.Second, true, middle.PersistentHandler)
			if err != nil {
				resps.InternalServerError(c, "Failed to create refresh token")
				return
			}
			c.SetCookie("token", token, config.TokenExpireTime, "/", "", protocol.CookieSameSiteLaxMode, true, true)
			c.SetCookie("refresh_token", refreshToken, config.RefreshTokenExpireTime, "/", "", protocol.CookieSameSiteLaxMode, true, true)
			resps.Ok(c, "Login successful", map[string]any{
				"token":         token,
				"refresh_token": refreshToken,
			})
			return
		} else {
			resps.Forbidden(c, "Incorrect password")
			return
		}
	}
}

// Logout 用户登出
// User logout
func (UserApi) Logout(ctx context.Context, c *app.RequestContext) {
	// 删除cookie
	c.SetCookie("token", "", -1, "/", "", protocol.CookieSameSiteLaxMode, true, true)
	c.SetCookie("refresh_token", "", -1, "/", "", protocol.CookieSameSiteLaxMode, true, true)
	resps.Ok(c, "Logout successful")
}

// GetCaptcha 获取验证码
// Get captcha
func (UserApi) GetCaptcha(ctx context.Context, c *app.RequestContext) {
	resps.Ok(c, "ok", map[string]any{
		"provider": config.CaptchaType,
		"site_key": config.CaptchaSiteKey,
		"url":      config.CaptchaUrl,
	})
}

// GetUserOrgs 获取用户的组织
// Get user organizations
func (UserApi) GetOrgs(ctx context.Context, c *app.RequestContext) {
	userID := c.Param("id")
	crtUser := middle.Auth.GetUser(ctx, c)
	if userID != strconv.Itoa(int(crtUser.ID)) {
		resps.Forbidden(c, resps.PermissionDenied)
		return
	}
	page, limit := utils.Ctx.GetPageLimit(c)

	orgs, err := store.Org.ListByUserID(userID, page, limit)
	if err != nil {
		resps.InternalServerError(c, "Failed to get organizations")
		return
	}

	resps.Ok(c, resps.OK, map[string]any{
		"organizations": func() (orgDTOs []OrganizationDTO) {
			for _, org := range orgs {
				orgDTOs = append(orgDTOs, Org.ToDTO(&org))
			}
			return
		}(),
	})
}

// GetUserProjects 获取用户的项目
// Get user projects
func (UserApi) GetProjects(ctx context.Context, c *app.RequestContext) {
	userID := c.Param("id")
	crtUser := middle.Auth.GetUser(ctx, c)
	if userID != strconv.Itoa(int(crtUser.ID)) {
		resps.Forbidden(c, resps.PermissionDenied)
	}
	page, limit := utils.Ctx.GetPageLimit(c)

	projects, total, err := store.Project.ListByOwner(constants.OwnerTypeUser, userID, page, limit)
	if err != nil {
		resps.InternalServerError(c, "Failed to get projects")
	}
	resps.Ok(c, resps.OK, map[string]any{
		"projects": func() (projectDTOs []ProjectDTO) {
			for _, project := range projects {
				projectDTOs = append(projectDTOs, Project.toDTO(&project, false)) // 这里naloveyuki尝试修复了下gitd7b49ff的问题
			}
			return
		}(),
		"total": total,
	})
}

// GetUser 获取用户信息
// Get user information
func (UserApi) GetUser(ctx context.Context, c *app.RequestContext) {
	userID := c.Param("id")
	crtUser := middle.Auth.GetUser(ctx, c)
	if userID == "" {
		userID = strconv.Itoa(int(crtUser.ID))
	}
	if userID == strconv.Itoa(int(crtUser.ID)) {
		// 本人
		resps.Ok(c, "ok", map[string]any{
			"user": User.ToDTO(crtUser, true),
		})
	} else {
		// 其他人
		resps.Ok(c, "ok", map[string]any{
			"user": User.ToDTO(crtUser, false),
		})
	}
}

// Register 用户注册
// User registration
func (UserApi) Register(ctx context.Context, c *app.RequestContext) {
	// 接收参数
	request := &RegisterReq{}
	err := c.BindJSON(request)
	if err != nil {
		resps.BadRequest(c, "Parameter error")
		return
	}
	// TODO 校验邮箱验证码
	// 校验密码复杂度
	passwordLevel := config.GetInt("password_complexity", 3)
	if !utils.Password.CheckPasswordComplexity(request.Password, passwordLevel) {
		resps.BadRequest(c, "Password complexity is too low")
		return
	}
	// 判断用户名是否存在
	if store.User.IsNameExist(request.Username) {
		resps.BadRequest(c, "Username already exists")
		return
	}
	// 创建用户
	hashPassword, err := utils.Password.HashPassword(request.Password, config.JwtSecret)
	if err != nil {
		resps.InternalServerError(c, "Failed to hash password")
		return
	}
	err = store.User.Create(&models.User{
		Name:     request.Username,
		Email:    &request.Email,
		Password: &hashPassword,
	})
	if err != nil {
		resps.InternalServerError(c, "Failed to create user")
		return
	}
	resps.Ok(c, "Register successful", map[string]any{
		"user": UserDTO{
			Name:  request.Username,
			Email: &request.Email,
		},
	})
}

// UpdateUser 更新用户信息
// Update user information
func (UserApi) UpdateUser(ctx context.Context, c *app.RequestContext) {
	userDTO := &UserDTO{}
	if err := c.BindJSON(userDTO); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}

	crtUser := middle.Auth.GetUser(ctx, c)
	crtUser.Name = userDTO.Name
	crtUser.DisplayName = userDTO.DisplayName
	crtUser.Email = userDTO.Email
	crtUser.Description = userDTO.Description
	crtUser.AvatarURL = userDTO.Avatar
	crtUser.Language = userDTO.Language

	if err := store.User.Update(crtUser); err != nil {
		resps.InternalServerError(c, "Failed to update user")
		return
	}

	resps.Ok(c, resps.OK, map[string]any{})
}
