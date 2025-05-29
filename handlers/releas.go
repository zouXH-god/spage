package handlers

import (
	"context"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/models"
	"github.com/LiteyukiStudio/spage/resps"
	"github.com/LiteyukiStudio/spage/store"
	"github.com/LiteyukiStudio/spage/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"os"
	"time"
)

type ReleaseApi struct {
}

var Release = ReleaseApi{}

func (ReleaseApi) ToDTO(release *models.SiteRelease) ReleaseDTO {
	return ReleaseDTO{
		ID:   release.ID,
		Site: Site.ToDTO(&release.Site, false),
		Tag:  release.Tag,
	}
}

func (ReleaseApi) ReleaseList(ctx context.Context, c *app.RequestContext) {
	site := getSite(ctx)
	if site == nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	releaseList, err := store.Site.GetReleaseList(site.ID)
	if err != nil {
		resps.InternalServerError(c, resps.ParameterError)
		return
	}
	resps.Ok(c, resps.OK, map[string]any{
		"releases": func(releases []*models.SiteRelease) []ReleaseDTO {
			var releasesDTO []ReleaseDTO
			for _, release := range releases {
				releasesDTO = append(releasesDTO, Release.ToDTO(release))
			}
			return releasesDTO
		}(releaseList),
	})
}

func (ReleaseApi) Create(ctx context.Context, c *app.RequestContext) {
	req := CreateReleaseReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	site := getSite(ctx)
	if site == nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	// 检查 zip 文件
	valid, err := utils.IsValidZipFile(req.File)
	if !valid || err != nil {
		resps.BadRequest(c, "file is not a zip or zip file is invalid")
		return
	}
	// 拼接并生成保存路径
	releaseSaveDir := config.ReleaseSavePath + "/" + site.Name + "/" + req.Tag
	if err := os.MkdirAll(releaseSaveDir, os.ModePerm); err != nil {
		resps.InternalServerError(c, "create release directory error")
		return
	}
	releaseName := time.Now().Format("20060102150405") + ".zip"
	releaseSavePath := releaseSaveDir + "/" + releaseName
	// 保存文件
	if err := c.SaveUploadedFile(req.File, releaseSavePath); err != nil {
		resps.InternalServerError(c, "create release file error")
		return
	}
	// 创建文件记录
	file := models.File{
		Path: releaseSavePath,
	}
	if err := store.File.Create(&file); err != nil {
		resps.InternalServerError(c, "create file record error")
		return
	}
	// 创建发布记录
	release := models.SiteRelease{
		SiteID: site.ID,
		Tag:    req.Tag,
		FileID: file.ID,
	}
	if err := store.Site.CreateRelease(&release); err != nil {
		resps.InternalServerError(c, "create release record error")
		return
	}
	// TODO 创建发布任务
	resps.Ok(c, resps.OK, map[string]any{
		"release": Release.ToDTO(&release),
	})
}

func (ReleaseApi) Delete(ctx context.Context, c *app.RequestContext) {
	req := DeleteReleaseReq{}
	if err := c.BindAndValidate(&req); err != nil {
		resps.BadRequest(c, resps.ParameterError)
		return
	}
	// 获取 release
	release, err := store.Site.GetReleaseById(req.ID)
	if err != nil {
		resps.NotFound(c, resps.TargetNotFound)
		return
	}
	// 删除文件
	err = os.RemoveAll(release.File.Path)
	if err != nil {
		resps.InternalServerError(c, "delete file error")
		return
	}
	// 删除 release 记录
	err = store.Site.DeleteRelease(release)
	if err != nil {
		resps.InternalServerError(c, "delete release record error")
		return
	}
	resps.Ok(c, resps.OK)
}
