package router

import (
	"context"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

//var H *server.Hertz

func Run() error {
	// 运行路由
	H := server.New(server.WithHostPorts(":" + config.ServerPort))
	H.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{"message": "Hello World"})
	})

	apiV1 := H.Group("/api/v1")
	{
		userGroup := apiV1.Group("/user")
		{
			userGroup.POST("/register")    // Register
			userGroup.POST("/login")       // Login
			userGroup.POST("/logout")      // Logout
			userGroup.PUT("")              // Update user info
			userGroup.DELETE("")           // Delete user
			userGroup.GET("/:id")          // Get user info
			userGroup.GET("/:id/projects") // Get user projects
			userGroup.GET("/:id/orgs")     // Get user organizations
		}
		orgGroup := apiV1.Group("/org")
		{
			orgGroup.POST("")             // Create organization
			orgGroup.PUT("")              // Update organization
			orgGroup.DELETE("")           // Delete organization
			orgGroup.GET("/:id")          // Get organization info
			orgGroup.GET("/:id/projects") // Get organization projects
		}
		projectGroup := apiV1.Group("/project")
		{
			projectGroup.POST("")    // Create project
			projectGroup.PUT("")     // Update project
			projectGroup.DELETE("")  // Delete project
			projectGroup.GET("/:id") // Get project info
		}
		siteGroup := apiV1.Group("/site")
		{
			siteRelease := siteGroup.Group("/release")
			{
				siteRelease.POST("")   // Create site release
				siteRelease.DELETE("") // Delete site release
			}
			siteGroup.POST("")    // Create site
			siteGroup.PUT("")     // Update site
			siteGroup.DELETE("")  // Delete site
			siteGroup.GET("/:id") // Get site info
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
