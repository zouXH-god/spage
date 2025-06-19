package router

import (
	"github.com/LiteyukiStudio/spage/spage/handlers"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerProjectGroup(group *route.RouterGroup) {
	// TODO 整理project路由，使用读写检查中间件
	projectGroup := group.Group("/project", handlers.Project.UserProjectAuth)
	{
		projectGroup.POST("", handlers.Project.Create)                  // 创建项目 Create project
		projectGroup.PUT("/:id", handlers.Project.Update)               // 更新项目 Update project
		projectGroup.DELETE("/:id", handlers.Project.Delete)            // 删除项目 Delete project
		projectGroup.GET("/:id", handlers.Project.Info)                 // 获取项目信息 Get project info
		projectGroup.GET("/:id/owners", handlers.Project.GetOwners)     // 获取项目所有者 Get project owners
		projectGroup.PUT("/:id/owner", handlers.Project.AddOwner)       // 更新项目所有者 Add project owner
		projectGroup.DELETE("/:id/owner", handlers.Project.DeleteOwner) // 删除项目所有者 Delete project owner
		projectGroup.GET("/:id/sites", handlers.Project.GetSites)       // 获取项目站点 Get project sites

	}
	siteGroup := group.Group("/site", handlers.Site.SiteAuth)
	{
		siteGroup.POST("", handlers.Site.Create)                     // 创建站点 Create site
		siteGroup.PUT("/:id", handlers.Site.Update)                  // 更新站点 Update site
		siteGroup.DELETE("/:id", handlers.Site.Delete)               // 删除站点 Delete site
		siteGroup.GET("/:id", handlers.Site.Info)                    // 获取网站信息 Get site info
		siteGroup.GET("/:id/releases", handlers.Release.ReleaseList) // 获取站点 release 列表
		siteRelease := siteGroup.Group("/:id/release")
		{
			siteRelease.POST("", handlers.Release.Create)                // 创建站点发布 Create site release
			siteRelease.DELETE("", handlers.Release.Delete)              // 删除站点版本 Delete site release
			siteRelease.POST("/activation", handlers.Release.Activation) // 指定使用该站点版本
		}
	}
}
