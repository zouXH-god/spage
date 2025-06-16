package handlers

import (
	"context"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/resps"
	"github.com/LiteyukiStudio/spage/spage/store"
	"github.com/LiteyukiStudio/spage/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"net/url"
)

type oidcType struct{}

var Oidc = oidcType{}

func (oidcType) ListOidcConfig(ctx context.Context, c *app.RequestContext) {
	oidcConfigs, err := store.Oidc.ListEnabledOidcConfig()
	if err != nil {
		resps.InternalServerError(c, err.Error())
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"oidc_configs": func() []map[string]any {
			var configsDto []map[string]any
			for _, oidcConfig := range oidcConfigs {
				state := utils.GenerateRandomString(32)
				// TODO 使用utils的键值内存储存和验证state
				configsDto = append(configsDto, map[string]any{
					"id":           oidcConfig.ID,
					"display_name": oidcConfig.DisplayName,
					"icon":         oidcConfig.Icon,
					"login_url": buildURL(oidcConfig.AuthorizationEndpoint, map[string]string{
						"client_id":     oidcConfig.ClientID,
						"redirect_uri":  config.BaseUrl + config.OidcUri + "/" + oidcConfig.Name,
						"response_type": "code",
						"scope":         "openid email profile",
						"state":         state,
					}),
					"name": oidcConfig.Name,
				})
			}
			return configsDto
		}(),
	})
}

func (oidcType) LoginOidcConfig(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	_ = name
	// TODO 实现OIDC登录
}

func buildURL(baseURL string, queryParams map[string]string) string {
	u, err := url.Parse(baseURL)
	if err != nil {
		return baseURL
	}
	q := u.Query()
	for key, value := range queryParams {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
