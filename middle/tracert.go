package middle

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type traceType struct{}

var Trace = traceType{}

func (traceType) UseTrace() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		path := string(c.Request.URI().Path())
		method := string(c.Request.Header.Method())

		c.Next(ctx)

		latency := time.Since(start)
		statusCode := c.Response.StatusCode()

		// 只记录必要信息，使用简洁格式
		message := method + " " + path + " " + strconv.Itoa(statusCode) + " " + latency.String()

		if statusCode >= 500 {
			logrus.Error(message)
		} else if statusCode >= 400 {
			logrus.Warn(message)
		} else {
			logrus.Info(message)
		}
	}
}
