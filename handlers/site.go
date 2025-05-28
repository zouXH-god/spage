package handlers

import (
	"context"
	"github.com/LiteyukiStudio/spage/models"
	"github.com/LiteyukiStudio/spage/resps"
	"github.com/LiteyukiStudio/spage/store"
	"github.com/cloudwego/hertz/pkg/app"
)

type SiteApi struct {
}

var Site = SiteApi{}

func (SiteApi) ToDTO(site *models.Site, full bool) SiteDTO {
	siteDTO := SiteDTO{
		Description: site.Description,
		ID:          site.ID,
		Name:        site.Name,
	}
	if full {
		siteDTO.Project = Project.toDTO(&site.Project, full)
		siteDTO.SubDomain = &site.SubDomain
		siteDTO.Domains = site.Domains
	}
	return siteDTO
}

func (SiteApi) Create(ctx context.Context, c *app.RequestContext) {
	req := CreateSiteReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	site := models.Site{
		Name:        req.Name,
		Description: req.Description,
		Domains:     req.Domains,
		ProjectID:   req.ProjectID,
		SubDomain:   *req.SubDomain,
	}
	if err := store.Site.Create(&site); err != nil {
		resps.InternalServerError(c, err.Error())
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"site": Site.ToDTO(&site, true),
	})
}
