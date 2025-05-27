package router

import (
	"context"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/handlers"
	"github.com/LiteyukiStudio/spage/middle"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

//var H *server.Hertz

func TODO() func(context.Context, *app.RequestContext) {
	return func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{"message": "Hello World" + string(c.Path())})
	}
}

func Run() error {
	// 运行路由
	H := server.New(server.WithHostPorts(":" + config.ServerPort))
	H.Use(middle.Cors.UseCors(), middle.Trace.UseTrace())
	apiV1 := H.Group("/api/v1")
	apiV1.Use(middle.Auth.UseAuth())
	apiV1WithoutAuth := H.Group("/api/v1")
	{
		apiV1WithoutAuth.POST("/user/register", handlers.User.Register).Use(middle.Captcha.UseCaptcha()) // Register
		apiV1WithoutAuth.POST("/user/login", handlers.User.Login).Use(middle.Captcha.UseCaptcha())
		apiV1WithoutAuth.GET("/user/captcha", handlers.User.GetCaptcha) // Get captcha
		apiV1WithoutAuth.POST("/user/logout", handlers.User.Logout)
		userGroup := apiV1.Group("/user")
		{
			userGroup.PUT("", handlers.User.UpdateUser)               // Update user info
			userGroup.GET("", handlers.User.GetUser)                  // Get user info (self)
			userGroup.GET("/:id", handlers.User.GetUser)              // Get user info
			userGroup.GET("/:id/projects", handlers.User.GetProjects) // Get user projects
			userGroup.GET("/:id/orgs", handlers.User.GetOrgs)         // Get user organizations
		}
		orgGroup := apiV1.Group("/org", handlers.Org.UserOrgAuth)
		{
			orgGroup.POST("", handlers.Org.CreateOrganization)                 // Create organization
			orgGroup.PUT("/:id", handlers.Org.UpdateOrganization)              // Update organization
			orgGroup.DELETE("/:id", handlers.Org.DeleteOrganization)           // Delete organization
			orgGroup.GET("/:id", handlers.Org.GetOrganization)                 // Get organization info
			orgGroup.GET("/:id/projects", handlers.Org.GetOrganizationProject) // Get organization projects
			orgGroup.GET("/:id/users", handlers.Org.GetOrganizationUsers)      // Get organization all member and owner
			orgGroup.PUT("/:id/users", handlers.Org.AddOrganizationUser)       // Add organization member or owner
			orgGroup.DELETE("/:id/users", handlers.Org.DeleteOrganizationUser) // Remove organization member or owner
		}
		projectGroup := apiV1.Group("/project", handlers.Project.UserProjectAuth)
		{
			projectGroup.POST("", handlers.Project.Create)       // Create project
			projectGroup.PUT("/:id", handlers.Project.Update)    // Update project
			projectGroup.DELETE("/:id", handlers.Project.Delete) // Delete project
			projectGroup.GET("/:id", handlers.Project.Info)      // Get project info
		}
		siteGroup := apiV1.Group("/site", TODO())
		{
			siteRelease := siteGroup.Group("/release", TODO())
			{
				siteRelease.POST("", TODO())   // Create site release
				siteRelease.DELETE("", TODO()) // Delete site release
			}
			siteGroup.POST("", TODO())    // Create site
			siteGroup.PUT("", TODO())     // Update site
			siteGroup.DELETE("", TODO())  // Delete site
			siteGroup.GET("/:id", TODO()) // Get site info
		}
		adminGroup := apiV1.Group("/admin").Use(middle.Auth.IsAdmin())
		{
			adminGroup.POST("/user", TODO()) // Create user
		}
	}

	// 设置静态文件目录
	web := H.Group("")
	{
		web.GET("/*any", handlers.WebHandler)
	}

	// 运行服务
	if config.Mode == "dev" {
		// Development mode
		err := H.Run()
		if err != nil {
			return err
		}
	} else {
		// Production mode
		H.Spin()
	}
	return nil
}
