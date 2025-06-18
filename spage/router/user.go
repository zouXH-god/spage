package router

import (
	"github.com/LiteyukiStudio/spage/spage/handlers"
	"github.com/LiteyukiStudio/spage/spage/middle"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerUserGroup(group *route.RouterGroup, groupWithoutAuth *route.RouterGroup) {
	userGroupWithoutAuth := groupWithoutAuth.Group("/user")
	{
		userGroupWithoutAuth.GET("/captcha", handlers.User.GetCaptchaConfig)
		userGroupWithoutAuth.POST("/logout", handlers.User.Logout)
		userGroupWithoutAuth.GET("/oidc/config", handlers.Oidc.ListOidcConfig)
		userGroupWithoutAuth.GET("/oidc/login/:name", handlers.Oidc.LoginOidcConfig)

		userGroupWithoutAuthNeedCaptcha := userGroupWithoutAuth.Group("").Use(middle.Captcha.UseCaptcha())
		{
			userGroupWithoutAuthNeedCaptcha.POST("/register", handlers.User.Register)
			userGroupWithoutAuthNeedCaptcha.POST("/login", handlers.User.Login)
		}
	}

	userGroup := group.Group("/user")
	{
		userGroup.PUT("", handlers.User.UpdateUser) // 更新用户信息 Update user info
		userGroup.GET("", handlers.User.GetUser)    // 获取用户信息 Get user info
		userGroup.POST("/token", handlers.User.CreateApiToken)
		userGroup.GET("/token/list", handlers.User.ListApiToken)
		userGroup.DELETE("/token/:id", handlers.User.RevokeApiToken)
		userGroup.GET("/:id", handlers.User.GetUser)              // 获取用户信息 Get user info
		userGroup.GET("/:id/projects", handlers.User.GetProjects) // 获取用户项目 Get user projects
		userGroup.GET("/:id/orgs", handlers.User.GetOrgs)         // 获取用户组织 Get user orgs
	}
}
