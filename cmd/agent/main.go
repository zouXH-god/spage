package main

import (
	"context"
	"github.com/LiteyukiStudio/spage/agent/caddy"
	"github.com/LiteyukiStudio/spage/agent/router"
	"github.com/LiteyukiStudio/spage/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	caddy.StartDaemon(ctx)

	if err := router.Run(ctx,
		server.WithHostPorts(":"+utils.GetEnv("PORT", "8888")),
		server.WithMaxRequestBodySize(utils.GetEnvInt("MAX_REQUEST_BODY_SIZE", 1073741824))); err != nil {
		logrus.Panicf("failed to run router: %v", err)
		return
	}
}
