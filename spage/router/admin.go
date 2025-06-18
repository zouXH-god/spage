package router

import (
	"github.com/LiteyukiStudio/spage/spage/handlers"
	"github.com/LiteyukiStudio/spage/spage/middle"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerAdminGroup(group *route.RouterGroup) {
	adminGroup := group.Group("/admin") // 管理员路由
	adminGroup.Use(middle.Auth.IsAdmin())
	{
		adminUser := adminGroup.Group("/user")
		{
			adminUser.POST("", handlers.Admin.CreateUser) // 创建用户 Create user
		}
		adminNode := adminGroup.Group("/node")
		{
			adminNode.DELETE("")    // 删除节点
			adminNode.POST("")      // 创建节点（上传ssh密码自动化创建）
			adminNode.GET("/token") // 获取节点令牌
		}
		adminOidc := adminGroup.Group("/oidc")
		{
			adminOidc.POST("", handlers.Admin.CreateOidcConfig)
			adminOidc.DELETE("/:id")
			adminOidc.PUT("/:id", handlers.Admin.UpdateOidcConfig)
		}
	}
}
