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

// UserOrgAuth 组织权限检查
func (OrgApi) UserOrgAuth(ctx context.Context, c *app.RequestContext) {
	user := middle.Auth.GetUser(ctx, c)
	// 当 id 为空的 POST 请求默认为 create
	orgIdStr := c.Param("id")
	if orgIdStr == "" && string(c.Method()) == "POST" {
		return
	}
	// 查询
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
	// 判断权限 (GET 请求需要用户权限，其他请求需要管理员权限)
	var authType string
	if string(c.Method()) == "GET" {
		authType = "member"
	} else {
		authType = "owner"
	}
	if authType == store.Org.GetUserAuth(org, user.ID) {
		context.WithValue(ctx, "userOrg", org)
		return
	} else {
		resps.BadRequest(c, resps.ParameterError)
		c.Abort()
		return
	}
}

func (OrgApi) ToDTO(org *models.Organization) OrganizationDTO {
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
		"organization": Org.ToDTO(&org),
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
		"organization": Org.ToDTO(org),
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
				projectDTOs = append(projectDTOs, Project.toDTO(&project))
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
		"organization": Org.ToDTO(org),
	})
}

func (OrgApi) GetOrganizationUsers(ctx context.Context, c *app.RequestContext) {
	org := getOrg(ctx)
	resps.Ok(c, resps.OK, map[string]any{
		"members": func() (users []UserDTO) {
			for _, user := range org.Members {
				users = append(users, User.ToDTO(user, false))
			}
			return
		},
		"owners": func() (users []UserDTO) {
			for _, user := range org.Owners {
				users = append(users, User.ToDTO(&user, false))
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
