package router

import (
	middle3 "github.com/LiteyukiStudio/spage/pkg/middle"
	"github.com/LiteyukiStudio/spage/pkg/utils"
	"github.com/LiteyukiStudio/spage/spage/handlers"
	"github.com/LiteyukiStudio/spage/spage/middle"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
)

// Run 运行路由服务
func Run(opts ...config.Option) error {
	h := server.New(opts...)
	h.Use(middle3.Cors.UseCors(), middle3.Trace.UseTrace())

	apiV1 := h.Group("/api/v1")
	apiV1.Use(middle.Auth.UseAuth(true))
	apiV1WithoutAuth := h.Group("/api/v1")
	registerUserGroup(apiV1, apiV1WithoutAuth)
	registerOrgGroup(apiV1)
	registerOwnerGroup(apiV1)
	registerProjectGroup(apiV1)
	registerFileGroup(apiV1, apiV1WithoutAuth)
	registerAdminGroup(apiV1)
	registerNodeGroup(apiV1)
	registerMetaGroup(apiV1, apiV1WithoutAuth)
	h.GET("/*any", handlers.WebHandler)
	return utils.RunWithMode(h, "dev")
}
