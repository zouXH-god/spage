package router

import "github.com/cloudwego/hertz/pkg/route"

func registerNodeGroup(group *route.RouterGroup) {
	nodeGroup := group.Group("/node") // 节点路由
	{
		nodeGroup.GET("")  // 节点心跳上报
		nodeGroup.POST("") // 注册节点（由节点自行请求）
	}
}
