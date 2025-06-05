package main

import (
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/router"
	"github.com/LiteyukiStudio/spage/store"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting page server...")

	// 第一步初始化配置文件
	if err := config.Init(); err != nil {
		logrus.Panicf("failed to load config: %v", err)
		return
	}

	// 初始化数据相关
	if err := store.Init(); err != nil {
		logrus.Panicf("failed to init data store: %v", err)
		return
	}

	// TODO 创建节点检查任务 task/node_check.go

	if err := router.Run(); err != nil {
		logrus.Panicf("failed to run router: %v", err)
		return
	}
}
