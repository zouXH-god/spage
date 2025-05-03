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
		_ = userGroup
	}

	// 运行服务
	if config.Mode == "dev" {
		err := H.Run()
		if err != nil {
			return err
		}
	} else {
		H.Spin()
	}
	return nil
}
