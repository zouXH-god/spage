package api

import (
	"context"
	"github.com/LiteyukiStudio/spage/resps"
	"github.com/cloudwego/hertz/pkg/app"
)

type UserApi struct{}

var User = UserApi{}

func (UserApi) Login(ctx context.Context, c *app.RequestContext) {
	loginReq := &LoginReq{}
	err := c.BindJSON(loginReq)
	if err != nil {
		resps.BadRequest(c, "Parameter error")
	}
}
