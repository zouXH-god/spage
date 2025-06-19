package main

import (
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/spage/router"
	"github.com/LiteyukiStudio/spage/spage/store"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting page server...")

	if err := config.Init(); err != nil {
		logrus.Panicf("failed to load config: %v", err)
		return
	}

	if err := store.Init(); err != nil {
		logrus.Panicf("failed to init data store: %v", err)
		return
	}

	if err := router.Run(server.WithHostPorts(":"+config.ServerPort), server.WithMaxRequestBodySize(config.FileMaxSize)); err != nil {
		logrus.Panicf("failed to run router: %v", err)
		return
	}
}
