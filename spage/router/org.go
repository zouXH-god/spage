package router

import (
	"github.com/LiteyukiStudio/spage/spage/handlers"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerOrgGroup(group *route.RouterGroup) {
	orgGroup := group.Group("/org")
	orgGroupWithMember := group.Group("/org").Use(handlers.Org.IsOrgMemberMiddleware)
	orgGroupWithOwner := group.Group("/org").Use(handlers.Org.IsOrgOwnerMiddleware)
	{
		orgGroup.POST("", handlers.Org.CreateOrganization)                           // 创建组织 Create organization
		orgGroupWithMember.GET("/:id", handlers.Org.GetOrganization)                 // 获取组织信息 Get organization info
		orgGroupWithMember.GET("/:id/projects", handlers.Org.GetOrganizationProject) // 获取组织项目 Get organization projects
		orgGroupWithMember.GET("/:id/users", handlers.Org.GetOrganizationUsers)      // 获取组织所有成员和所有者 Get organization users
		orgGroupWithOwner.PUT("/:id", handlers.Org.UpdateOrganization)               // 更新组织 Update organization
		orgGroupWithOwner.DELETE("/:id", handlers.Org.DeleteOrganization)            // 删除组织 Delete organization
		orgGroupWithOwner.PUT("/:id/user", handlers.Org.AddOrganizationUser)         // 添加组织成员或所有者 Add organization user
		orgGroupWithOwner.DELETE("/:id/user", handlers.Org.DeleteOrganizationUser)   // 删除组织成员或所有者 Delete organization user
	}
}
