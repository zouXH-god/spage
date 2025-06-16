package router

import (
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/spage/handlers"
	"github.com/LiteyukiStudio/spage/spage/middle"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// Run 运行路由服务
// Run router service
func Run() error {
	// 运行路由 Run router
	H := server.New(server.WithHostPorts(":"+config.ServerPort), server.WithMaxRequestBodySize(config.FileMaxSize))
	H.Use(middle.Cors.UseCors(), middle.Trace.UseTrace())
	apiV1 := H.Group("/api/v1")

	apiV1.Use(middle.Auth.UseAuth(true))
	apiV1WithoutAuth := H.Group("/api/v1")
	apiV1WithoutAuthAndCaptcha := H.Group("/api/v1") // 不需要登录和验证码的路由 Group without auth and captcha
	{
		apiV1WithoutAuthAndCaptcha.GET("/user/captcha", handlers.User.GetCaptcha) // 取验证码 Get captcha
		apiV1WithoutAuth.POST("/user/logout", handlers.User.Logout)
		apiV1WithoutAuth.POST("/user/register", handlers.User.Register).Use(middle.Captcha.UseCaptcha()) // 注册 Register
		apiV1WithoutAuth.POST("/user/login", handlers.User.Login).Use(middle.Captcha.UseCaptcha())
		// OAuth2
		apiV1WithoutAuth.GET("/user/oidc/config", handlers.Oidc.ListOidcConfig)
		//apiV1WithoutAuth.POST("/user/oidc/login/:name")

		userGroup := apiV1.Group("/user")
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
		orgGroup := apiV1.Group("/org", handlers.Org.UserOrgAuth)
		{
			orgGroup.POST("", handlers.Org.CreateOrganization)                 // 创建组织 Create organization
			orgGroup.PUT("/:id", handlers.Org.UpdateOrganization)              // 更新组织 Update organization
			orgGroup.DELETE("/:id", handlers.Org.DeleteOrganization)           // 删除组织 Delete organization
			orgGroup.GET("/:id", handlers.Org.GetOrganization)                 // 获取组织信息 Get organization info
			orgGroup.GET("/:id/projects", handlers.Org.GetOrganizationProject) // 获取组织项目 Get organization projects
			orgGroup.GET("/:id/users", handlers.Org.GetOrganizationUsers)      // 获取组织所有成员和所有者 Get organization users
			orgGroup.PUT("/:id/users", handlers.Org.AddOrganizationUser)       // 添加组织成员或所有者 Add organization user
			orgGroup.DELETE("/:id/users", handlers.Org.DeleteOrganizationUser) // 删除组织成员或所有者 Delete organization user
		}
		projectGroup := apiV1.Group("/project", handlers.Project.UserProjectAuth)
		{
			projectGroup.POST("", handlers.Project.Create)                  // 创建项目 Create project
			projectGroup.PUT("/:id", handlers.Project.Update)               // 更新项目 Update project
			projectGroup.DELETE("/:id", handlers.Project.Delete)            // 删除项目 Delete project
			projectGroup.GET("/:id", handlers.Project.Info)                 // 获取项目信息 Get project info
			projectGroup.GET("/:id/owners", handlers.Project.GetOwners)     // 获取项目所有者 Get project owners
			projectGroup.PUT("/:id/owner", handlers.Project.AddOwner)       // 更新项目所有者 Add project owner
			projectGroup.DELETE("/:id/owner", handlers.Project.DeleteOwner) // 删除项目所有者 Delete project owner
			projectGroup.GET("/:id/sites", handlers.Project.GetSites)       // 获取项目站点 Get project sites
			siteGroup := projectGroup.Group("/:id/site", handlers.Site.SiteAuth)
			{
				siteGroup.POST("", handlers.Site.Create)            // 创建站点 Create site
				siteGroup.PUT("/:site_id", handlers.Site.Update)    // 更新站点 Update site
				siteGroup.DELETE("/:site_id", handlers.Site.Delete) // 删除站点 Delete site
				siteGroup.GET("/:site_id", handlers.Site.Info)      // 获取网站信息 Get site info

				siteGroup.GET("/:site_id/releases", handlers.Release.ReleaseList) // 获取站点 release 列表
				siteRelease := siteGroup.Group("/:site_id/release")
				{
					siteRelease.POST("", handlers.Release.Create)                // 创建站点发布 Create site release
					siteRelease.DELETE("", handlers.Release.Delete)              // 删除站点版本 Delete site release
					siteRelease.POST("/activation", handlers.Release.Activation) // 指定使用该站点版本
				}
			}
		}
		fileGroup := apiV1.Group("/file")
		fileGroupWithoutAuth := apiV1WithoutAuth.Group("/file")
		{
			fileGroup.POST("", handlers.File.UploadFileStream)      // 上传文件 Upload file
			fileGroup.DELETE("")                                    // 删除文件 Delete file
			fileGroupWithoutAuth.GET("/:id", handlers.File.GetFile) // 下载文件 Download file
		}
		adminGroup := apiV1.Group("/admin") // 管理员路由
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

		nodeGroup := apiV1.Group("/node") // 节点路由
		{
			nodeGroup.GET("")  // 节点心跳上报
			nodeGroup.POST("") // 注册节点（由节点自行请求）
		}

	}

	// 设置静态文件目录 Set static file directory
	web := H.Group("")
	{
		web.GET("/*any", handlers.WebHandler)
	}

	// 运行服务 Run service
	if config.Mode == "dev" {
		// 开发模式 Development mode
		err := H.Run()
		if err != nil {
			return err
		}
	} else {
		// 生产模式 Production mode
		H.Spin()
	}
	return nil
}
