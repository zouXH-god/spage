package handlers

import (
	"context"
	"github.com/LiteyukiStudio/spage/constants"
	"github.com/LiteyukiStudio/spage/middle"
	"github.com/LiteyukiStudio/spage/models"
	"github.com/LiteyukiStudio/spage/resps"
	"github.com/LiteyukiStudio/spage/store"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

type OrgApi struct{}

var Org = OrgApi{}

func getOrg(ctx context.Context) *models.Organization {
	org, ok := ctx.Value("userOrg").(*models.Organization)
	if !ok {
		return nil
	}
	return org
}

// 组织权限检查
func orgAuth(ctx context.Context, c *app.RequestContext, authType string) {
	user := middle.Auth.GetUser(ctx, c)
	// 查询
	orgIdStr := c.Param("id")
	orgId, err := strconv.Atoi(orgIdStr)
	if err != nil {
		resps.BadRequest(c, resps.ParameterError)
		c.Abort()
		return
	}
	org, err := store.Org.GetOrgById(uint(orgId))
	if err != nil || org == nil {
		resps.NotFound(c, resps.TargetNotFound)
		c.Abort()
		return
	}
	if authType == "owner" {
		// 检查用户是否为组织所有者
		for _, owner := range org.Owners {
			if owner.ID == user.ID {
				context.WithValue(ctx, "userOrg", org)
				return
			}
		}
	} else if authType == "member" {
		// 检查用户是否为组织成员
		for _, member := range org.Members {
			if member.ID == user.ID {
				// 存储组织信息
				context.WithValue(ctx, "userOrg", org)
				return
			}
		}
	} else {
		resps.BadRequest(c, resps.ParameterError)
		c.Abort()
		return
	}
}

// UserIsOrgOwner 验证用户是否为组织所有者
func (OrgApi) UserIsOrgOwner(ctx context.Context, c *app.RequestContext) {
	orgAuth(ctx, c, "owner")
}

// UserIsOrgMember 验证用户是否为组织成员
func (OrgApi) UserIsOrgMember(ctx context.Context, c *app.RequestContext) {
	orgAuth(ctx, c, "member")
}

func organizationModelToDTO(org *models.Organization) OrganizationDTO {
	return OrganizationDTO{
		ID:           org.ID,
		Name:         org.Name,
		DisplayName:  org.DisplayName,
		Email:        org.Email,
		Description:  org.Description,
		AvatarURL:    org.AvatarURL,
		ProjectLimit: org.ProjectLimit,
	}
}

func (OrgApi) CreateOrganization(ctx context.Context, c *app.RequestContext) {
	// 绑定参数
	req := &CreateOrgReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	// 检验组织名称是否存在
	if store.Org.OrgNameIsExist(req.Name) {
		resps.BadRequest(c, "organization name already exists")
		return
	}
	// 创建组织
	user := middle.Auth.GetUser(ctx, c)
	org := models.Organization{
		Name:        req.Name,
		DisplayName: &req.DisplayName,
		Email:       req.Email,
		Description: req.Description,
		AvatarURL:   req.AvatarURL,
		Members:     []*models.User{user},
		Owners:      []models.User{*user},
	}
	if err := store.Org.CreateOrg(&org); err != nil {
		resps.InternalServerError(c, err.Error())
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"organization": organizationModelToDTO(&org),
	})
}

func (OrgApi) UpdateOrganization(ctx context.Context, c *app.RequestContext) {
	req := &UpdateOrgReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	org := getOrg(ctx)
	// 更新
	org.DisplayName = req.DisplayName
	org.Email = req.Email
	org.Description = *req.Description
	org.AvatarURL = req.AvatarURL
	if err := store.Org.UpdateOrg(org); err != nil {
		resps.InternalServerError(c, err.Error())
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"organization": organizationModelToDTO(org),
	})
}

func (OrgApi) GetOrganizationProject(ctx context.Context, c *app.RequestContext) {
	req := &GetOrgProjectReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	org := getOrg(ctx)
	// 查询
	projects, total, err := store.Project.ListByOwner(constants.OwnerTypeOrg, strconv.Itoa(int(org.ID)), req.Page, req.Limit)
	if err != nil {
		resps.InternalServerError(c, "Failed to get projects")
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"projects": func() (projectDTOs []ProjectDTO) {
			for _, project := range projects {
				projectDTOs = append(projectDTOs, ProjectDTO{
					ID:          project.ID,
					Name:        project.Name,
					DisplayName: project.DisplayName,
					Description: project.Description,
					OwnerID:     project.OwnerID,
					OwnerType:   project.OwnerType,
					SiteLimit:   project.SiteLimit,
				})
			}
			return
		}(),
		"total": total,
	})
}

func (OrgApi) DeleteOrganization(ctx context.Context, c *app.RequestContext) {
	org := getOrg(ctx)
	// 删除组织
	if err := store.Org.DeleteOrg(org); err != nil {
		resps.InternalServerError(c, err.Error())
		return
	}
	resps.Ok(c, resps.OK)
}

func (OrgApi) GetOrganization(ctx context.Context, c *app.RequestContext) {
	org := getOrg(ctx)
	resps.Ok(c, resps.OK, map[string]any{
		"organization": organizationModelToDTO(org),
	})
}

func (OrgApi) GetOrganizationUsers(ctx context.Context, c *app.RequestContext) {
	org := getOrg(ctx)
	err := store.Org.LoadOrgUsers(org)
	if err != nil {
		resps.InternalServerError(c, err.Error())
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"members": func() (users []UserDTO) {
			for _, user := range org.Members {
				u := UserDTO{}
				u.FromUser(user)
				users = append(users, u)
			}
			return
		},
		"owners": func() (users []UserDTO) {
			for _, user := range org.Owners {
				u := UserDTO{}
				u.FromUser(&user)
				users = append(users, u)
			}
			return
		},
	})
}

func (OrgApi) AddOrganizationUser(ctx context.Context, c *app.RequestContext) {
	req := &OrgUserReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	// 查询用户
	user, err := store.User.GetByID(req.UserID)
	if err != nil || user == nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	// 新增
	org := getOrg(ctx)
	if req.Role == "member" {
		org.Members = append(org.Members, user)
		if err = store.Org.UpdateOrg(org); err != nil {
			resps.InternalServerError(c, err.Error())
			return
		}
	} else if req.Role == "owner" {
		org.Owners = append(org.Owners, *user)
		if err = store.Org.UpdateOrg(org); err != nil {
			resps.InternalServerError(c, err.Error())
			return
		}
	}
	resps.Ok(c, resps.OK)
}

func (OrgApi) DeleteOrganizationUser(ctx context.Context, c *app.RequestContext) {
	req := &OrgUserReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	// 查询用户
	user, err := store.User.GetByID(req.UserID)
	if err != nil || user == nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	// 删除
	org := getOrg(ctx)
	if req.Role == "member" {
		for i, member := range org.Members {
			if member.ID == user.ID {
				org.Members = append(org.Members[:i], org.Members[i+1:]...)
				break
			}
		}
		if err = store.Org.UpdateOrg(org); err != nil {
			resps.InternalServerError(c, err.Error())
			return
		}
	} else if req.Role == "owner" {
		for i, owner := range org.Owners {
			if owner.ID == user.ID {
				org.Owners = append(org.Owners[:i], org.Owners[i+1:]...)
				break
			}
			if err = store.Org.UpdateOrg(org); err != nil {
				resps.InternalServerError(c, err.Error())
				return
			}
		}
	}
	resps.Ok(c, resps.OK)
}
