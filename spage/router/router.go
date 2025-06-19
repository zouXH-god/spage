package router

import (
	"fmt"
	"github.com/LiteyukiStudio/spage/constants"
	"github.com/LiteyukiStudio/spage/spage/handlers"
	"github.com/LiteyukiStudio/spage/spage/middle"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
)

// Run 运行路由服务
func Run(opts ...config.Option) error {
	h := server.New(opts...)
	h.Use(middle.Cors.UseCors(), middle.Trace.UseTrace())

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

	return runWithMode(h, "dev")
}

func runWithMode(h *server.Hertz, mode string) error {
	if mode == constants.ModeDev {
		err := h.Run()
		if err != nil {
			return err
		}
	} else if mode == constants.ModeProd {
		h.Spin()
	} else {
		return fmt.Errorf("unsupported mode: %s", mode)
	}
	return nil
}
