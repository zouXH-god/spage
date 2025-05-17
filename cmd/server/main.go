package main

import (
	"github.com/LiteyukiStudio/spage/router"
	"github.com/LiteyukiStudio/spage/store"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting page server...")

	if err := store.Init(); err != nil {
		logrus.Panicf("failed to init data store: %v", err)
		return
	}

	if err := router.Run(); err != nil {
		logrus.Panicf("failed to run router: %v", err)
		return
	}
}
