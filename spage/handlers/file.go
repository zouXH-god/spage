package handlers

import (
	"context"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/resps"
	"github.com/LiteyukiStudio/spage/spage/middle"
	"github.com/LiteyukiStudio/spage/spage/models"
	"github.com/LiteyukiStudio/spage/spage/store"
	"github.com/LiteyukiStudio/spage/utils"
	"github.com/LiteyukiStudio/spage/utils/filedriver"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/sirupsen/logrus"
	"io"
	"path/filepath"
	"strconv"
)

type FileApi struct{}

var File = FileApi{}

func (FileApi) UploadFileStream(ctx context.Context, c *app.RequestContext) {
	// 获取文件信息
	file, err := c.FormFile("file")
	if err != nil {
		logrus.Error("无法读取文件: ", err)
		resps.BadRequest(c, err.Error())
		return
	}

	// 初始化文件驱动
	driver, err := filedriver.GetFileDriver(config.FileDriverConfig)
	if err != nil {
		logrus.Error("获取文件驱动失败: ", err)
		resps.InternalServerError(c, "获取文件驱动失败")
		return
	}

	// 校验文件哈希
	if hashForm := string(c.FormValue("hash")); hashForm != "" {
		dir, fileName := utils.FilePath(hashForm)
		storagePath := filepath.Join(dir, fileName)
		if _, err := driver.Stat(c, storagePath); err == nil {
			resps.Ok(c, "文件已存在", map[string]any{"hash": hashForm})
			return
		}
	}

	// 打开文件
	src, err := file.Open()
	if err != nil {
		logrus.Error("无法打开文件: ", err)
		resps.BadRequest(c, err.Error())
		return
	}
	defer src.Close()

	// 计算文件哈希值
	hash, err := utils.FileHashFromStream(src)
	if err != nil {
		logrus.Error("计算文件哈希失败: ", err)
		resps.BadRequest(c, err.Error())
		return
	}

	// 根据哈希值生成存储路径
	dir, fileName := utils.FilePath(hash)
	storagePath := filepath.Join(dir, fileName)
	// 保存文件
	if _, err := src.Seek(0, io.SeekStart); err != nil {
		logrus.Error("无法重置文件流位置: ", err)
		resps.BadRequest(c, err.Error())
		return
	}
	if err := driver.Save(c, storagePath, src); err != nil {
		logrus.Error("保存文件失败: ", err)
		resps.InternalServerError(c, err.Error())
		return
	}

	// 数据库索引建立
	crtUser := middle.Auth.GetUserWithBlock(ctx, c)
	fileModel := &models.File{
		Hash:       hash,
		UploaderID: crtUser.ID,
	}

	if err := store.File.Create(fileModel); err != nil {
		logrus.Error("数据库索引建立失败: ", err)
		resps.InternalServerError(c, "数据库索引建立失败")
		return
	}
	resps.Ok(c, "文件上传成功", map[string]any{"hash": hash, "id": fileModel.ID})
}

func (FileApi) GetFile(ctx context.Context, c *app.RequestContext) {
	// TODO 私有文件访问控制，目前暂时没有私有化文件，可以先不做
	fileIdString := c.Param("id")
	fileId, err := strconv.ParseUint(fileIdString, 10, 64)
	if err != nil {
		logrus.Error("无效的文件ID: ", err)
		resps.BadRequest(c, "无效的文件ID")
		return
	}
	fileModel, err := store.File.GetByID(uint(fileId))
	if err != nil {
		logrus.Error("获取文件信息失败: ", err)
		resps.InternalServerError(c, "获取文件信息失败")
		return
	}
	driver, err := filedriver.GetFileDriver(config.FileDriverConfig)
	if err != nil {
		logrus.Error("获取文件驱动失败: ", err)
		resps.InternalServerError(c, "获取文件驱动失败")
		return
	}
	filePath := filepath.Join(utils.FilePath(fileModel.Hash))
	driver.Get(c, filePath)
}
