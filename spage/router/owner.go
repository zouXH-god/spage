package router

import (
	"github.com/LiteyukiStudio/spage/spage/handlers"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerOwnerGroup(group *route.RouterGroup) {
	ownerGroup := group.Group("/owner")
	{
		ownerGroup.GET("/:name", handlers.Owner.GetByName) // 获取所有者列表 Get owners list
	}
}
