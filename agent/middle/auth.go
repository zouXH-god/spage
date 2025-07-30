package middle

import (
	"context"
	"crypto/hmac"
	"github.com/LiteyukiStudio/spage/agent/env"
	"github.com/cloudwego/hertz/pkg/app"
)

func UseAuth() app.HandlerFunc {
	expectedToken := "Bearer " + env.AgentApiToken
	return func(ctx context.Context, c *app.RequestContext) {
		authToken := string(c.GetHeader("Authorization"))
		if authToken == "" {
			c.AbortWithStatusJSON(401, map[string]string{"error": "Authorization header is required"})
			return
		}
		if !hmac.Equal([]byte(authToken), []byte(expectedToken)) {
			c.AbortWithStatusJSON(403, map[string]string{"error": "Forbidden"})
			return
		}
		c.Next(ctx)
	}
}
