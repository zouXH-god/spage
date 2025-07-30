package router

import (
	"context"
	"errors"
	middle2 "github.com/LiteyukiStudio/spage/pkg/middle"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(ctx context.Context, opts ...config.Option) error {
	h := server.New(opts...)
	h.Use(middle2.Cors.UseCors(), middle2.Trace.UseTrace())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		if err := h.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Errorf("Agent Server startup failed: %v", err)
		}
	}()
	select {
	case <-ctx.Done():
		logrus.Info("Received context cancellation, shutting down Agent server...")
	case sig := <-quit:
		logrus.Infof("Received signal %s, shutting down Agent server...", sig)
	}
	logrus.Info("Shutting down Agent server gracefully...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("Failed to shut down Agent server gracefully: %v", err)
		return err
	}
	logrus.Info("Agent server shut down gracefully")
	return nil
}
