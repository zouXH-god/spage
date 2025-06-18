package router

import (
	"github.com/LiteyukiStudio/spage/spage/handlers"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerMetaGroup(group *route.RouterGroup, groupWithoutAuth *route.RouterGroup) {
	metaGroup := groupWithoutAuth.Group("/meta")
	{
		metaGroup.GET("/info", handlers.Meta.GetMetaInfo) // 获取版本信息 Get version info
	}
}
