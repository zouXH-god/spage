package handlers

import (
	"context"
	"github.com/LiteyukiStudio/spage/resps"
	"github.com/cloudwego/hertz/pkg/app"
)

type AdminApi struct{}

var Admin = AdminApi{}

func (AdminApi) CreateUser(ctx context.Context, c *app.RequestContext) {
	var userDTO *UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
}
