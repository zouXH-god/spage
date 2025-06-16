package middle

import (
	"context"
	"fmt"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/constants"
	models "github.com/LiteyukiStudio/spage/spage/models"
	store "github.com/LiteyukiStudio/spage/spage/store"
	"github.com/LiteyukiStudio/spage/utils"
	"time"

	"github.com/LiteyukiStudio/spage/resps"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
)

type authType struct{}

var Auth = authType{}

// PersistentHandler 持久化处理函数，使用依赖注入到 utils 中防止循环引用
func PersistentHandler(userID uint) (*models.Token, error) {
	token, err := store.JWT.CreateToken(userID)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// RevokeChecker 令牌撤销检查器，使用依赖注入到 utils 中防止循环引用
func RevokeChecker(tokenID uint) bool {
	return store.JWT.IsTokenRevoked(tokenID)
}

// UseAuth Middleware function for authentication
func (authType) UseAuth(block bool) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// Authentication method 2: Use Cookie (Enable silent refresh)
		token := string(c.Cookie("token"))
		if token == "" {
			// Try to refresh by refresh_token
			refreshToken := string(c.Cookie("refresh_token"))
			if refreshToken == "" {
				resps.BadRequest(c, "Refresh token not found 1")
				c.Abort()
				return
			}
			// Verify refresh token
			refreshClaims, err := utils.Token.ParseToken(refreshToken, RevokeChecker)
			if err != nil {
				resps.Unauthorized(c, "Refresh token expired or invalid 2")
				c.Abort()
				return
			}
			// Generate new access token
			newToken, err := utils.Token.CreateToken(refreshClaims.UserID, time.Duration(config.TokenExpireTime)*time.Second, false, PersistentHandler)
			if err != nil {
				resps.InternalServerError(c, "Create access token failed 3")
				c.Abort()
				return
			}
			// Set new access token
			c.SetCookie("token", newToken, config.TokenExpireTime, "/", "", protocol.CookieSameSiteLaxMode, true, true)
			// Save user information and continue request
			ctx = context.WithValue(ctx, "user", refreshClaims.UserID)
			c.Next(ctx)
			return
		}

		// Cookie contains token, verify its validity
		claims, err := utils.Token.ParseToken(token, RevokeChecker)
		if err != nil {
			// Token is invalid, try to refresh
			refreshToken := string(c.Cookie("refresh_token"))
			if refreshToken == "" {
				resps.Unauthorized(c, "Refresh token not found 4")
				c.Abort()
				return
			}
			// Verify refresh token
			refreshClaims, err := utils.Token.ParseToken(refreshToken, RevokeChecker)
			if err != nil {
				fmt.Println(err)
				resps.Unauthorized(c, "Refresh token expired or invalid 5")
				c.Abort()
				return
			}
			// Generate new access token
			newToken, err := utils.Token.CreateToken(refreshClaims.UserID, time.Duration(config.TokenExpireTime)*time.Second, false, PersistentHandler)
			if err != nil {
				resps.InternalServerError(c, "Create access token failed 6")
				c.Abort()
				return
			}
			// Set new access token
			c.SetCookie("token", newToken, config.TokenExpireTime, "/", "", protocol.CookieSameSiteLaxMode, true, true)
			// Save user information and continue request
			ctx = context.WithValue(ctx, "user", refreshClaims.UserID)
			c.Next(ctx)
			return
		}
		// Token is valid, continue request
		ctx = context.WithValue(ctx, "user", claims.UserID)
		c.Next(ctx)
	}
}

// IsAdmin 是一个中间件，用于检查用户是否为管理员
func (authType) IsAdmin() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		user := Auth.GetUser(ctx, c)
		if user.Role != constants.RoleAdmin {
			resps.Forbidden(c, "Permission denied")
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}

// GetUser 从已认证的上下文中获取用户信息,如果用户不存在则终止请求并返回
func (authType) GetUser(ctx context.Context, c *app.RequestContext) *models.User {
	userID := ctx.Value("user").(uint)
	if userID == 0 {
		resps.Unauthorized(c, resps.TargetNotFound)
		c.Abort()
		return nil
	}

	user, err := store.User.GetByID(uint(userID))
	if err != nil {
		resps.Unauthorized(c, resps.TargetNotFound)
		c.Abort()
		return nil
	}
	if user == nil {
		resps.Unauthorized(c, resps.TargetNotFound)
		c.Abort()
		return nil
	}
	return user
}
