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

type ProjectApi struct {
}

var Project = ProjectApi{}

func (ProjectApi) toDTO(project *models.Project, full bool) ProjectDTO {
	projectDto := ProjectDTO{
		Description: project.Description,
		DisplayName: project.DisplayName,
		ID:          project.ID,
		Name:        project.Name,
		OwnerType:   project.OwnerType,
	}
	if full {
		projectDto.OwnerID = project.OwnerID
		projectDto.Owners = func([]models.User) (owners []UserDTO) {
			for _, owner := range project.Owners {
				owners = append(owners, User.ToDTO(&owner, false))
			}
			return
		}(project.Owners)
		projectDto.SiteLimit = project.SiteLimit
	}
	return projectDto
}

func getProject(ctx context.Context) *models.Project {
	project, ok := ctx.Value("userProject").(*models.Project)
	if !ok {
		return nil
	}
	return project
}

func (ProjectApi) UserProjectAuth(ctx context.Context, c *app.RequestContext) {
	user := middle.Auth.GetUser(ctx, c)
	projectIdStr := c.Param("id")
	// 当id为空默认为创建
	if projectIdStr == "" && string(c.Method()) == "POST" {
		return
	}
	// 查询项目
	projectId, err := strconv.Atoi(projectIdStr)
	if err != nil {
		resps.BadRequest(c, resps.ParameterError)
		c.Abort()
		return
	}
	project, err := store.Project.GetByID(uint(projectId))
	if err != nil || project == nil {
		resps.NotFound(c, resps.TargetNotFound)
		c.Abort()
		return
	}
	// 项目权限判断
	if store.Project.UserIsOwner(project, user.ID) {
		context.WithValue(ctx, "userProject", project)
		return
	}
	// 组织权限判断
	if project.OwnerType == constants.OwnerTypeOrg {
		// 组织查询
		org, err := store.Org.GetOrgById(project.OwnerID)
		if err != nil || org == nil {
			resps.InternalServerError(c, resps.ParameterError)
			c.Abort()
			return
		}
		// 请求类型判断
		var authType string
		if string(c.Method()) == "GET" {
			authType = "member"
		} else {
			authType = "owner"
		}
		if authType == store.Org.GetUserAuth(org, user.ID) {
			context.WithValue(ctx, "userOrg", org)
			context.WithValue(ctx, "userProject", project)
			return
		} else {
			resps.BadRequest(c, resps.ParameterError)
			c.Abort()
			return
		}
	} else {
		resps.BadRequest(c, resps.ParameterError)
		c.Abort()
	}
}

func (ProjectApi) Create(ctx context.Context, c *app.RequestContext) {
	req := CreateProjectReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	user := middle.Auth.GetUser(ctx, c)
	// 校验权限
	if req.OwnerType == constants.OwnerTypeOrg {
		// 如果为组织，需要具有组织管理员权限
		org, err := store.Org.GetOrgById(req.OwnerID)
		if err != nil || org == nil {
			resps.InternalServerError(c, resps.ParameterError)
			return
		}
		if store.Org.GetUserAuth(org, user.ID) != "member" {
			resps.Forbidden(c, resps.PermissionDenied)
		}
	} else if req.OwnerType == constants.OwnerTypeUser {
		// 如果为用户，仅允许为自己添加
		req.OwnerID = user.ID
	} else {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	project := &models.Project{
		Description: req.Description,
		DisplayName: req.DisplayName,
		Name:        req.Name,
		OwnerID:     req.OwnerID,
		OwnerType:   req.OwnerType,
		Owners:      []models.User{*user},
	}
	if err := store.Project.Create(project); err != nil {
		resps.InternalServerError(c, resps.ParameterError)
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"project": Project.toDTO(project, true),
	})
}

func (ProjectApi) Update(ctx context.Context, c *app.RequestContext) {
	req := UpdateProjectReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	project := getProject(ctx)
	if project == nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	// 更新数据
	project.Description = *req.Description
	project.DisplayName = req.DisplayName
	project.Name = *req.Name
	if err := store.Project.Update(project); err != nil {
		resps.InternalServerError(c, resps.ParameterError)
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"project": Project.toDTO(project, true),
	})
}

func (ProjectApi) Delete(ctx context.Context, c *app.RequestContext) {
	project := getProject(ctx)
	if project == nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	if err := store.Project.Delete(project); err != nil {
		resps.InternalServerError(c, resps.ParameterError)
		return
	}
	resps.Ok(c, resps.OK)
}

func (ProjectApi) Info(ctx context.Context, c *app.RequestContext) {
	project := getProject(ctx)
	if project == nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"project": Project.toDTO(project, true),
	})
}

func (ProjectApi) GetOwners(ctx context.Context, c *app.RequestContext) {
	project := getProject(ctx)
	if project == nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"owners": func([]models.User) (owners []UserDTO) {
			for _, owner := range project.Owners {
				owners = append(owners, User.ToDTO(&owner, false))
			}
			return
		}(project.Owners),
	})
}

func (ProjectApi) AddOwner(ctx context.Context, c *app.RequestContext) {
	project := getProject(ctx)
	if project == nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	req := ProjectUserReq{}
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
	// 判断用户是否已存在权限列表
	for _, owner := range project.Owners {
		if owner.ID == user.ID {
			resps.BadRequest(c, "User already exists in the permission list")
			return
		}
	}
	// 添加用户
	if err := store.Project.AddOwner(project, user); err != nil {
		resps.InternalServerError(c, resps.ParameterError)
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"project": Project.toDTO(project, true),
	})
}

func (ProjectApi) DeleteOwner(ctx context.Context, c *app.RequestContext) {
	project := getProject(ctx)
	if project == nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	req := ProjectUserReq{}
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
	// 删除用户
	if err := store.Project.DeleteOwner(project, user); err != nil {
		resps.InternalServerError(c, resps.ParameterError)
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"project": Project.toDTO(project, true),
	})
}

func (ProjectApi) GetSites(ctx context.Context, c *app.RequestContext) {
	req := GetSiteListReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	project := getProject(ctx)
	if project == nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	sites, total, err := store.Project.GetSiteList(project, req.Page, req.Limit)
	if err != nil {
		resps.InternalServerError(c, "Failed to get sites")
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"sites": func([]models.Site) (siteDTOs []SiteDTO) {
			for _, site := range sites {
				siteDTOs = append(siteDTOs, Site.ToDTO(&site, false))
			}
			return
		}(sites),
		"total": total,
	})
}
