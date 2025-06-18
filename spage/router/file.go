package router

import (
	"github.com/LiteyukiStudio/spage/spage/handlers"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerFileGroup(group *route.RouterGroup, groupWithoutAuth *route.RouterGroup) {
	fileGroup := group.Group("/file")
	fileGroupWithoutAuth := groupWithoutAuth.Group("/file")
	{
		fileGroup.POST("", handlers.File.UploadFileStream)      // 上传文件 Upload file
		fileGroup.DELETE("")                                    // 删除文件 Delete file
		fileGroupWithoutAuth.GET("/:id", handlers.File.GetFile) // 下载文件 Download file
	}
}
