package router

import (
	"context"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/middle"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

//var H *server.Hertz

func emptyHandler() func(context.Context, *app.RequestContext) {
	return func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{"message": "Hello World"})
	}
}

func Run() error {
	// 运行路由
	H := server.New(server.WithHostPorts(":" + config.ServerPort))
	H.Use(middle.Cors())
	H.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{"message": "Hello World"})
	})

	apiV1 := H.Group("/api/v1")
	{
		userGroup := apiV1.Group("/user")
		{
			userGroup.POST("/register", emptyHandler())    // Register
			userGroup.POST("/login", emptyHandler())       // Login
			userGroup.POST("/logout", emptyHandler())      // Logout
			userGroup.PUT("", emptyHandler())              // Update user info
			userGroup.DELETE("", emptyHandler())           // Delete user
			userGroup.GET("/:id", emptyHandler())          // Get user info
			userGroup.GET("/:id/projects", emptyHandler()) // Get user projects
			userGroup.GET("/:id/orgs", emptyHandler())     // Get user organizations
		}
		orgGroup := apiV1.Group("/org", emptyHandler())
		{
			orgGroup.POST("", emptyHandler())             // Create organization
			orgGroup.PUT("", emptyHandler())              // Update organization
			orgGroup.DELETE("", emptyHandler())           // Delete organization
			orgGroup.GET("/:id", emptyHandler())          // Get organization info
			orgGroup.GET("/:id/projects", emptyHandler()) // Get organization projects
		}
		projectGroup := apiV1.Group("/project", emptyHandler())
		{
			projectGroup.POST("", emptyHandler())    // Create project
			projectGroup.PUT("", emptyHandler())     // Update project
			projectGroup.DELETE("", emptyHandler())  // Delete project
			projectGroup.GET("/:id", emptyHandler()) // Get project info
		}
		siteGroup := apiV1.Group("/site", emptyHandler())
		{
			siteRelease := siteGroup.Group("/release", emptyHandler())
			{
				siteRelease.POST("", emptyHandler())   // Create site release
				siteRelease.DELETE("", emptyHandler()) // Delete site release
			}
			siteGroup.POST("", emptyHandler())    // Create site
			siteGroup.PUT("", emptyHandler())     // Update site
			siteGroup.DELETE("", emptyHandler())  // Delete site
			siteGroup.GET("/:id", emptyHandler()) // Get site info
		}
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
