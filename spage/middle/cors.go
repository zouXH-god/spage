package middle

import (
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/constants"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
	"github.com/sirupsen/logrus"
)

type corsType struct{}

var Cors = corsType{}

// UseCors 跨域中间件
// Cross-domain middleware
func (corsType) UseCors() app.HandlerFunc {
	var allowedOrigins []string
	if config.Mode == constants.ModeDev {
		allowedOrigins = []string{"*"}
		logrus.Infof("Allowed origins: %v", allowedOrigins)
	} else {
		allowedOrigins = []string{"*"}
		logrus.Infof("Allowed origins: %v", allowedOrigins)
	}
	return cors.New(
		cors.Config{
			AllowOrigins:     allowedOrigins,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
			AllowCredentials: true,
			MaxAge:           3600,
		})
}
