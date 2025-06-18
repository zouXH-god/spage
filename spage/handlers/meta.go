package handlers

import (
	"context"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/resps"
	"github.com/cloudwego/hertz/pkg/app"
)

type metaType struct{}

var Meta = &metaType{}

func (metaType) GetMetaInfo(ctx context.Context, c *app.RequestContext) {
	resps.Ok(c, resps.OK, map[string]any{
		"name":        config.Name,
		"icon":        config.Icon,
		"version":     config.Version,
		"build_time":  config.BuildTime,
		"commit_hash": config.CommitHash,
	})
}
