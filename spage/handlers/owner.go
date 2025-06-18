package handlers

import (
	"context"
	"github.com/LiteyukiStudio/spage/constants"
	"github.com/LiteyukiStudio/spage/resps"
	"github.com/LiteyukiStudio/spage/spage/middle"
	"github.com/LiteyukiStudio/spage/spage/store"
	"github.com/cloudwego/hertz/pkg/app"
)

type ownerType struct{}

var Owner = &ownerType{}

func (ownerType) GetByName(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	crtUser := middle.Auth.GetUserWithBlock(ctx, c)

	// try user
	user, err := store.User.GetByName(name)
	if err == nil && user != nil {
		if user.IsPrivate && (crtUser == nil || user.ID != crtUser.ID) {
			resps.NotFound(c, resps.TargetNotFound)
			return
		}
		resps.Ok(c, resps.OK, map[string]any{
			"type": constants.OwnerTypeUser,
			"id":   user.ID,
		})
		return
	}

	// try org
	org, err := store.Org.GetOrgByName(name)
	if err == nil && org != nil {

		if org.IsPrivate {
			isMemberOfOrg, err := store.User.IsOwnerOfOrg(crtUser.ID, org.ID)
			if err != nil || !isMemberOfOrg {
				resps.NotFound(c, resps.TargetNotFound)
				return
			}
		}

		resps.Ok(c, resps.OK, map[string]any{
			"type": constants.OwnerTypeOrg,
			"id":   org.ID,
		})
		return
	}
}
