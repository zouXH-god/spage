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
	H.Use(middle.Cors())

	apiV1 := H.Group("/api/v1")
	apiV1.Use(middle.Auth.UseAuth())
	apiV1WithoutAuth := H.Group("/api/v1")
	{
		apiV1WithoutAuth.POST("/user/register", TODO()) // Register
		apiV1WithoutAuth.POST("/user/login", handlers.User.Login)

		userGroup := apiV1.Group("/user")
		{
			userGroup.POST("/logout", TODO())            // Logout
			userGroup.PUT("", TODO())                    // Update user info
			userGroup.DELETE("", TODO())                 // Delete user
			userGroup.GET("/*id", handlers.User.GetUser) // Get user info
			userGroup.GET("/:id/projects", TODO())       // Get user projects
			userGroup.GET("/:id/orgs", TODO())           // Get user organizations
		}
		orgGroup := apiV1.Group("/org", TODO())
		{
			orgGroup.POST("", TODO())             // Create organization
			orgGroup.PUT("", TODO())              // Update organization
			orgGroup.DELETE("", TODO())           // Delete organization
			orgGroup.GET("/:id", TODO())          // Get organization info
			orgGroup.GET("/:id/projects", TODO()) // Get organization projects
		}
		projectGroup := apiV1.Group("/project", TODO())
		{
			projectGroup.POST("", TODO())    // Create project
			projectGroup.PUT("", TODO())     // Update project
			projectGroup.DELETE("", TODO())  // Delete project
			projectGroup.GET("/:id", TODO()) // Get project info
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
		web.GET("/:any", handlers.WebHandler)
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
