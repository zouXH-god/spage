package middle

import (
	"context"
	"fmt"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/constants"
	"github.com/LiteyukiStudio/spage/spage/models"
	"github.com/LiteyukiStudio/spage/spage/store"
	"github.com/LiteyukiStudio/spage/utils"
	"strings"
	"time"

	"github.com/LiteyukiStudio/spage/resps"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
)

type authType struct{}

var Auth = authType{}

// JwtPersistentHandler 持久化处理函数，使用依赖注入到 utils 中防止循环引用
func JwtPersistentHandler(userID uint) (*models.JsonWebToken, error) {
	token, err := store.Token.CreateJsonWebToken(userID)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ApiTokenPersistentHandler(token *models.ApiToken) error {
	return store.Token.CreateApiToken(token)
}

// RevokeChecker 令牌撤销检查器，使用依赖注入到 utils 中防止循环引用
func RevokeChecker(tokenID uint) bool {
	return store.Token.IsJsonWebTokenRevoked(tokenID)
}

// IsApiTokenValid 检查API令牌是否有效
func IsApiTokenValid(token string) (*models.ApiToken, error) {
	return store.Token.GetApiTokenByToken(token)
}

// UseAuth Middleware function for authentication
func (authType) UseAuth(block bool) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 尝试 Cookie 认证
		cookiePass, cookieClaims, cookieErr := handleCookieTokenAuth(c)
		if cookiePass {
			ctx = context.WithValue(ctx, "user", cookieClaims.UserID)
			ctx = context.WithValue(ctx, "auth_type", "cookie")
			c.Next(ctx)
			return
		}

		// 尝试 Bearer 认证
		apiPass, apiToken, apiErr := handleBearerTokenAuth(c)
		if apiPass {
			ctx = context.WithValue(ctx, "user", apiToken.UserID)
			ctx = context.WithValue(ctx, "auth_type", "api")
			c.Next(ctx)
			return
		}

		// 认证失败处理
		if block {
			// 可以记录具体的错误原因便于调试
			if cookieErr != nil {
				c.Set("auth_error_cookie", cookieErr.Error())
			}
			if apiErr != nil {
				c.Set("auth_error_api", apiErr.Error())
			}
			resps.Unauthorized(c, resps.UnauthorizedText)
			c.Abort()
			return
		}

		// 非阻塞模式，继续但标记未认证
		ctx = context.WithValue(ctx, "authenticated", false)
		c.Next(ctx)
	}
}

// GetUser 改进版
func (authType) GetUser(ctx context.Context, c *app.RequestContext) *models.User {
	userIDValue := ctx.Value("user")
	if userIDValue == nil {
		resps.Unauthorized(c, resps.UnauthorizedText)
		c.Abort()
		return nil
	}

	userID := userIDValue.(uint)
	user, err := store.User.GetByID(userID)
	if err != nil || user == nil || user.ID == 0 {
		resps.Unauthorized(c, resps.UnauthorizedText)
		c.Abort()
		return nil
	}

	return user
}

// handleBearerTokenAuth
func handleBearerTokenAuth(c *app.RequestContext) (bool, *models.ApiToken, error) {
	token := strings.TrimPrefix(string(c.GetHeader("Authorization")), "Bearer ")
	apiToken, err := utils.Token.ParseApiToken(token, IsApiTokenValid)
	if err != nil {
		return false, nil, err
	}
	if apiToken == nil || apiToken.UserID == 0 {
		return false, nil, fmt.Errorf("invalid API token")
	}
	return true, apiToken, nil
}

// handleCookieTokenAuth
func handleCookieTokenAuth(c *app.RequestContext) (bool, *utils.Claims, error) {
	if claims, err := utils.Token.ParseJsonWebToken(string(c.Cookie("token")), RevokeChecker); err == nil {
		return true, claims, err
	}
	if refreshClaims, err := utils.Token.ParseJsonWebToken(string(c.Cookie("refresh_token")), RevokeChecker); err == nil {
		newToken, err := utils.Token.CreateJsonWebToken(refreshClaims.UserID, time.Duration(config.TokenExpireTime)*time.Second, false, JwtPersistentHandler)
		if err != nil {
			return false, nil, err
		}
		c.SetCookie("token", newToken, config.TokenExpireTime, "/", "", protocol.CookieSameSiteLaxMode, true, true)
		return true, refreshClaims, nil
	}
	return false, nil, fmt.Errorf("authentication failed")
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

// SetTokenForCookie 设置令牌到 Cookie 中 自动响应错误和返回，无需处理错误
func (authType) SetTokenForCookie(c *app.RequestContext, user *models.User, response bool, remember bool) {
	token, err := utils.Token.CreateJsonWebToken(user.ID, time.Duration(config.TokenExpireTime)*time.Second, false, JwtPersistentHandler)
	if err != nil {
		resps.InternalServerError(c, "Failed to create token")
		return
	}

	// 根据 remember 参数决定 refresh token 的过期时间
	refreshExpire := config.RefreshTokenExpireTime
	if remember {
		refreshExpire = config.RefreshTokenExpireTime * 7 // 延长7倍
	}

	refreshToken, err := utils.Token.CreateJsonWebToken(user.ID, time.Duration(refreshExpire)*time.Second, false, JwtPersistentHandler)
	if err != nil {
		resps.InternalServerError(c, "Failed to create refresh token")
		return
	}
	c.SetCookie("token", token, config.TokenExpireTime, "/", "", protocol.CookieSameSiteLaxMode, true, true)
	c.SetCookie("refresh_token", refreshToken, refreshExpire, "/", "", protocol.CookieSameSiteLaxMode, true, true)

	if !response {
		return
	}
	resps.Ok(c, "Login successful", map[string]any{
		"token":         token,
		"refresh_token": refreshToken,
		"user_id":       user.ID,
	})
	return
}
