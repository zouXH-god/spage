package router

import (
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/spage/handlers"
	"github.com/LiteyukiStudio/spage/spage/middle"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// Run 运行路由服务
func Run() error {
	// 运行路由
	H := server.New(server.WithHostPorts(":"+config.ServerPort), server.WithMaxRequestBodySize(config.FileMaxSize))
	H.Use(middle.Cors.UseCors(), middle.Trace.UseTrace())
	apiV1 := H.Group("/api/v1")

	apiV1.Use(middle.Auth.UseAuth())
	apiV1WithoutAuth := H.Group("/api/v1")
	apiV1WithoutAuthAndCaptcha := H.Group("/api/v1") // 不需要登录和验证码的路由
	{
		apiV1WithoutAuthAndCaptcha.GET("/user/captcha", handlers.User.GetCaptcha) // 取验证码
		apiV1WithoutAuth.POST("/user/logout", handlers.User.Logout)
		apiV1WithoutAuth.POST("/user/register", handlers.User.Register).Use(middle.Captcha.UseCaptcha()) // 注册
		apiV1WithoutAuth.POST("/user/login", handlers.User.Login).Use(middle.Captcha.UseCaptcha())
		userGroup := apiV1.Group("/user")
		{
			userGroup.PUT("", handlers.User.UpdateUser)               // 更新用户信息
			userGroup.GET("", handlers.User.GetUser)                  // 获取用户信息
			userGroup.GET("/:id", handlers.User.GetUser)              // 获取用户信息
			userGroup.GET("/:id/projects", handlers.User.GetProjects) // 获取用户项目
			userGroup.GET("/:id/orgs", handlers.User.GetOrgs)         // 获取用户组织
		}
		orgGroup := apiV1.Group("/org", handlers.Org.UserOrgAuth)
		{
			orgGroup.POST("", handlers.Org.CreateOrganization)                 // 创建组织
			orgGroup.PUT("/:id", handlers.Org.UpdateOrganization)              // 更新组织
			orgGroup.DELETE("/:id", handlers.Org.DeleteOrganization)           // 删除组织
			orgGroup.GET("/:id", handlers.Org.GetOrganization)                 // 获取组织信息
			orgGroup.GET("/:id/projects", handlers.Org.GetOrganizationProject) // 获取组织项目
			orgGroup.GET("/:id/users", handlers.Org.GetOrganizationUsers)      // 获取组织所有成员和所有者
			orgGroup.PUT("/:id/users", handlers.Org.AddOrganizationUser)       // 添加组织成员或所有者
			orgGroup.DELETE("/:id/users", handlers.Org.DeleteOrganizationUser) // 删除组织成员或所有者
		}
		projectGroup := apiV1.Group("/project", handlers.Project.UserProjectAuth)
		{
			projectGroup.POST("", handlers.Project.Create)                  // 创建项目
			projectGroup.PUT("/:id", handlers.Project.Update)               // 更新项目
			projectGroup.DELETE("/:id", handlers.Project.Delete)            // 删除项目
			projectGroup.GET("/:id", handlers.Project.Info)                 // 获取项目信息
			projectGroup.GET("/:id/owners", handlers.Project.GetOwners)     // 获取项目所有者
			projectGroup.PUT("/:id/owner", handlers.Project.AddOwner)       // 更新项目所有者
			projectGroup.DELETE("/:id/owner", handlers.Project.DeleteOwner) // 删除项目所有者
			projectGroup.GET("/:id/sites", handlers.Project.GetSites)       // 获取项目站点
			siteGroup := projectGroup.Group("/:id/site", handlers.Site.SiteAuth)
			{
				siteGroup.POST("", handlers.Site.Create)            // 创建站点
				siteGroup.PUT("/:site_id", handlers.Site.Update)    // 更新站点
				siteGroup.DELETE("/:site_id", handlers.Site.Delete) // 删除站点
				siteGroup.GET("/:site_id", handlers.Site.Info)      // 获取网站信息

				siteGroup.GET("/:site_id/releases", handlers.Release.ReleaseList) // 获取站点 release 列表
				siteRelease := siteGroup.Group("/:site_id/release")
				{
					siteRelease.POST("", handlers.Release.Create)                // 创建站点发布
					siteRelease.DELETE("", handlers.Release.Delete)              // 删除站点版本
					siteRelease.POST("/activation", handlers.Release.Activation) // 指定使用该站点版本
				}
			}
		}
		fileGroup := apiV1.Group("/file")
		{
			fileGroup.POST("", handlers.File.UploadFileStream) // 上传文件
			fileGroup.GET("")                                  // 下载文件
			fileGroup.DELETE("")                               // 删除文件
		}
		adminGroup := apiV1.Group("/admin") // 管理员路由
		adminGroup.Use(middle.Auth.IsAdmin())
		{
			adminUser := adminGroup.Group("/user")
			{
				adminUser.POST("", handlers.Admin.CreateUser) // 创建用户
			}
			adminNode := adminGroup.Group("/node")
			{
				adminNode.DELETE("")    // 删除节点
				adminNode.POST("")      // 创建节点（上传ssh密码自动化创建）
				adminNode.GET("/token") // 获取节点令牌
			}
		}
		nodeGroup := apiV1.Group("/node") // 节点路由
		{
			nodeGroup.GET("")  // 节点心跳上报
			nodeGroup.POST("") // 注册节点（由节点自行请求）
		}

	}

	// 设置静态文件目录
	web := H.Group("")
	{
		web.GET("/*any", handlers.WebHandler)
	}

	// 运行服务
	if config.Mode == "dev" {
		// 开发模式
		err := H.Run()
		if err != nil {
			return err
		}
	} else {
		// 生产模式
		H.Spin()
	}
	return nil
}
