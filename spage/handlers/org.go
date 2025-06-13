package handlers

import (
	"context"
	"github.com/LiteyukiStudio/spage/spage/constants"
	"github.com/LiteyukiStudio/spage/spage/middle"
	"github.com/LiteyukiStudio/spage/spage/models"
	store "github.com/LiteyukiStudio/spage/spage/store"
	"strconv"

	"github.com/LiteyukiStudio/spage/resps"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrgApi struct{}

var Org = OrgApi{}

// OrganizationDTO 组织信息数据传输对象
// Organization Information Data Transfer Object (DTO)
func getOrg(ctx context.Context) *models.Organization {
	org, ok := ctx.Value("userOrg").(*models.Organization)
	if !ok {
		return nil
	}
	return org
}

// UserOrgAuth 组织权限检查
// Organization permission check
func (OrgApi) UserOrgAuth(ctx context.Context, c *app.RequestContext) {
	user := middle.Auth.GetUser(ctx, c)
	// 当 id 为空的 POST 请求默认为 create
	orgIdStr := c.Param("id")
	if orgIdStr == "" && string(c.Method()) == "POST" {
		return
	}
	// 查询 Check
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
	// Determine permissions (GET requests require user permissions, other requests require admin permissions)
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

// OrganizationDTO 组织信息数据传输对象
// Organization Information Data Transfer Object (DTO)
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

// CreateOrganization 创建组织
// Create Organization
func (OrgApi) CreateOrganization(ctx context.Context, c *app.RequestContext) {
	// 绑定参数
	// Bind parameters
	req := &CreateOrgReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	// 检验组织名称是否存在
	// Check if the organization name already exists
	if store.Org.OrgNameIsExist(req.Name) {
		resps.BadRequest(c, "organization name already exists")
		return
	}
	// 创建组织
	// Create organization
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

// UpdateOrganization 更新组织信息
// Update Organization Information
func (OrgApi) UpdateOrganization(ctx context.Context, c *app.RequestContext) {
	req := &UpdateOrgReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	org := getOrg(ctx)
	// 更新 Update
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

// GetOrganizationProject 获取组织项目
// Get Organization Projects
func (OrgApi) GetOrganizationProject(ctx context.Context, c *app.RequestContext) {
	req := &GetOrgProjectReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	org := getOrg(ctx)
	// 查询 Query
	projects, total, err := store.Project.ListByOwner(constants.OwnerTypeOrg, strconv.Itoa(int(org.ID)), req.Page, req.Limit)
	if err != nil {
		resps.InternalServerError(c, "Failed to get projects")
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"projects": func() (projectDTOs []ProjectDTO) {
			for _, project := range projects {
				projectDTOs = append(projectDTOs, Project.toDTO(&project, false)) // 这里naloveyuki尝试修复了下gitd7b49ff的问题
			}
			return
		}(),
		"total": total,
	})
}

// DeleteOrganization 删除组织
// Delete Organization
func (OrgApi) DeleteOrganization(ctx context.Context, c *app.RequestContext) {
	org := getOrg(ctx)
	// 删除组织
	if err := store.Org.DeleteOrg(org); err != nil {
		resps.InternalServerError(c, err.Error())
		return
	}
	resps.Ok(c, resps.OK)
}

// GetOrganization 获取组织信息
// Get Organization Information
func (OrgApi) GetOrganization(ctx context.Context, c *app.RequestContext) {
	org := getOrg(ctx)
	resps.Ok(c, resps.OK, map[string]any{
		"organization": Org.ToDTO(org),
	})
}

// GetOrganizationUsers 获取组织用户
// Get Organization Users
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

// AddOrganizationUser 添加组织用户
// Add Organization User
func (OrgApi) AddOrganizationUser(ctx context.Context, c *app.RequestContext) {
	req := &OrgUserReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	// 查询 Query
	user, err := store.User.GetByID(req.UserID)
	if err != nil || user == nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	// 新增 Add
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

// DeleteOrganizationUser 删除组织用户
// Delete Organization User
func (OrgApi) DeleteOrganizationUser(ctx context.Context, c *app.RequestContext) {
	req := &OrgUserReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	// 查询 Query
	user, err := store.User.GetByID(req.UserID)
	if err != nil || user == nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	// 删除 Delete
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
