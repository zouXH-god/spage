package middle

import (
	"context"
	"fmt"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/spage/constants"
	models "github.com/LiteyukiStudio/spage/spage/models"
	store "github.com/LiteyukiStudio/spage/spage/store"
	"github.com/LiteyukiStudio/spage/utils"
	"strings"
	"time"

	"github.com/LiteyukiStudio/spage/resps"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
)

type authType struct{}

var Auth = authType{}

// PersistentHandler 持久化处理函数，使用依赖注入到 utils 中防止循环引用
// Persistent Handler Function, using dependency injection to prevent circular references
func PersistentHandler(userID uint) (*models.Token, error) {
	token, err := store.JWT.CreateToken(userID)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// RevokeChecker 令牌撤销检查器，使用依赖注入到 utils 中防止循环引用
// Token Revocation Checker, using dependency injection to prevent circular references
func RevokeChecker(tokenID uint) bool {
	return store.JWT.IsTokenRevoked(tokenID)
}

// UseAuth 中间件函数
// Middleware function for authentication
func (authType) UseAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 1. 检查认证方式
		// 1. Check authentication method
		authHeader := string(c.GetHeader("Authorization"))

		// 认证方式1：使用 Authorization Header
		// Authentication method 1: Use Authorization Header
		if authHeader != "" {
			// 检查 token 是否以 "Bearer " 开头,如果是，则去掉前缀
			// Check if the token starts with "Bearer ", if so, remove the prefix
			token := authHeader
			if strings.HasPrefix(token, "Bearer ") {
				token = strings.TrimPrefix(token, "Bearer ")
			}

			// 验证令牌
			// Verify token
			claims, err := utils.Token.ParseToken(token, RevokeChecker)
			if err != nil {
				resps.Unauthorized(c, "Invalid token")
				c.Abort()
				return
			}

			// 将用户信息存储到上下文中
			// Store user information in the context
			c.Set("user", claims.UserID)
			c.Next(ctx)
			return
		}

		// 认证方式2：使用 Cookie（启用无感刷新）
		// Authentication method 2: Use Cookie (Enable silent refresh)
		token := string(c.Cookie("token"))
		if token == "" {
			// 尝试通过 refresh_token 刷新
			// Try to refresh by refresh_token
			refreshToken := string(c.Cookie("refresh_token"))
			if refreshToken == "" {
				resps.BadRequest(c, "Refresh token not found 1")
				c.Abort()
				return
			}

			// 验证刷新令牌
			// Verify refresh token
			refreshClaims, err := utils.Token.ParseToken(refreshToken, RevokeChecker)
			if err != nil {
				resps.Unauthorized(c, "Refresh token expired or invalid 2")
				c.Abort()
				return
			}

			// 生成新的访问令牌
			// Generate new access token
			newToken, err := utils.Token.CreateToken(refreshClaims.UserID, time.Duration(config.TokenExpireTime)*time.Second, false, PersistentHandler)
			if err != nil {
				resps.InternalServerError(c, "Create access token failed 3")
				c.Abort()
				return
			}

			// 设置新的访问令牌
			// Set new access token
			c.SetCookie("token", newToken, config.TokenExpireTime, "/", "", protocol.CookieSameSiteLaxMode, true, true)

			// 保存用户信息并继续请求
			// Save user information and continue request
			c.Set("user", refreshClaims.UserID)
			c.Next(ctx)
			return
		}

		// Cookie 中存在 token，验证其有效性
		// Cookie contains token, verify its validity
		claims, err := utils.Token.ParseToken(token, RevokeChecker)
		if err != nil {
			// token 无效，尝试刷新
			// Token is invalid, try to refresh
			refreshToken := string(c.Cookie("refresh_token"))
			if refreshToken == "" {
				resps.Unauthorized(c, "Refresh token not found 4")
				c.Abort()
				return
			}

			// 验证刷新令牌
			// Verify refresh token
			refreshClaims, err := utils.Token.ParseToken(refreshToken, RevokeChecker)
			if err != nil {
				fmt.Println(err)
				resps.Unauthorized(c, "Refresh token expired or invalid 5")
				c.Abort()
				return
			}

			// 生成新的访问令牌
			// Generate new access token
			newToken, err := utils.Token.CreateToken(refreshClaims.UserID, time.Duration(config.TokenExpireTime)*time.Second, false, PersistentHandler)
			if err != nil {
				resps.InternalServerError(c, "Create access token failed 6")
				c.Abort()
				return
			}

			// 设置新的访问令牌
			// Set new access token
			c.SetCookie("token", newToken, config.TokenExpireTime, "/", "", protocol.CookieSameSiteLaxMode, true, true)

			// 保存用户信息并继续请求
			// Save user information and continue request
			c.Set("user", refreshClaims.UserID)
			c.Next(ctx)
			return
		}

		// token 有效，继续请求
		// Token is valid, continue request
		ctx = context.WithValue(ctx, "user", claims.UserID)
		c.Next(ctx)
	}
}

// IsAdmin 是一个中间件，用于检查用户是否为管理员
// IsAdmin is a middleware that checks if the user is an admin
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
// GetUser retrieves user information from the authenticated context, if the user does not exist it terminates the request and returns
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
