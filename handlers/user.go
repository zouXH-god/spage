package handlers

import (
	"context"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/middle"
	"github.com/LiteyukiStudio/spage/models"
	"github.com/LiteyukiStudio/spage/resps"
	"github.com/LiteyukiStudio/spage/store"
	"github.com/LiteyukiStudio/spage/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"strconv"
	"time"
)

type UserApi struct{}

var User = UserApi{}

func (UserApi) Login(ctx context.Context, c *app.RequestContext) {
	loginReq := &LoginReq{}
	// TODO: 这里需要验证验证码
	err := c.BindForm(loginReq)
	if err != nil {
		resps.BadRequest(c, "Parameter error")
		return
	}
	if loginReq.Username == "" || loginReq.Password == "" {
		resps.BadRequest(c, "Username or password cannot be empty")
		return
	}
	user, err := store.User.GetUserByName(loginReq.Username)
	if err != nil {
		user, err = store.User.GetUserByEmail(loginReq.Username)
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

func (UserApi) Logout(ctx context.Context, c *app.RequestContext) {
	// 删除cookie
	c.SetCookie("token", "", -1, "/", "", protocol.CookieSameSiteLaxMode, true, true)
	c.SetCookie("refresh_token", "", -1, "/", "", protocol.CookieSameSiteLaxMode, true, true)
	resps.Ok(c, "Logout successful")
}

func (UserApi) GetCaptcha(ctx context.Context, c *app.RequestContext) {
	resps.Ok(c, "ok", map[string]any{
		"provider": config.CaptchaType,
		"site_key": config.CaptchaSiteKey,
		"url":      config.CaptchaUrl,
	})
}

func (UserApi) GetUser(ctx context.Context, c *app.RequestContext) {
	userID := c.Param("id")
	crtUser := middle.Auth.GetUser(ctx, c)
	if userID == "" {
		userID = strconv.Itoa(int(crtUser.ID))
	}
	if userID == strconv.Itoa(int(crtUser.ID)) {
		// 本人
		resps.Ok(c, "ok", map[string]any{
			"user": UserDTO{
				ID:          crtUser.ID,
				Name:        crtUser.Name,
				DisplayName: crtUser.Name,
				Email:       crtUser.Email,
				Description: crtUser.Description,
				Avatar:      crtUser.Avatar,
				Role:        crtUser.Role,
				Language:    crtUser.Language,
			},
		})
	} else {
		// 其他人
		resps.Ok(c, "ok", map[string]any{
			"user": UserDTO{
				ID:          crtUser.ID,
				Name:        crtUser.Name,
				DisplayName: crtUser.Name,
				Email:       crtUser.Email,
				Description: crtUser.Description,
				Avatar:      crtUser.Avatar,
			},
		})
	}
}

func (UserApi) Register(ctx context.Context, c *app.RequestContext) {
	// 接收参数
	request := &RegisterReq{}
	err := c.BindForm(request)
	if err != nil {
		resps.BadRequest(c, "Parameter error")
		return
	}
	// TODO 校验邮箱验证码
	// 校验密码复杂度
	passwordLevel := config.GetInt("password_complexity", 3)
	if !utils.CheckPasswordComplexity(request.Password, passwordLevel) {
		resps.BadRequest(c, "Password complexity is too low")
		return
	}
	// 判断用户名是否存在
	if store.User.IsUserNameExist(request.Username) {
		resps.BadRequest(c, "Username already exists")
		return
	}
	// 创建用户
	hashPassword, err := utils.Password.HashPassword(request.Password, config.JwtSecret)
	if err != nil {
		resps.InternalServerError(c, "Failed to hash password")
		return
	}
	err = store.User.CreateUser(&models.User{
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

func (UserApi) UpdateUser(ctx context.Context, c *app.RequestContext) {

}
