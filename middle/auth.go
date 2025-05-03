package middle

import (
	"context"
	"github.com/LiteyukiStudio/spage/resps"
	"github.com/cloudwego/hertz/pkg/app"
)

type Auth struct{}

func (Auth) AuthToken() func(context.Context, *app.RequestContext) {
	return func(ctx context.Context, c *app.RequestContext) {
		token := string(c.GetHeader("Authorization"))
		if token == "" {
			resps.BadRequest(c, "Token is required")
			c.Abort()
			return
		}
	}
}
