package handlers

import (
	"context"
	"io"
	"mime"
	"path/filepath"

	"github.com/LiteyukiStudio/spage/static"
	"github.com/cloudwego/hertz/pkg/app"
)

// getMimeType 根据文件扩展名返回相应的 MIME 类型
// as the file extension returns the corresponding MIME type
func getMimeType(path string) string {
	// 这里可以根据文件扩展名返回相应的 MIME 类型
	ext := filepath.Ext(path)
	return mime.TypeByExtension(ext)
}

// WebHandler 处理静态文件请求的 Handler
// Handler for static file requests
func WebHandler(ctx context.Context, c *app.RequestContext) {
	path := "dist" + string(c.Path())
	file, err := static.WebFS.Open(path)
	if err != nil {
		// fallback 到 index.html
		file, err = static.WebFS.Open("dist/index.html")
		if err != nil {
			c.String(404, "File not found")
			return
		}
		defer file.Close()

		c.SetContentType("text/html; charset=utf-8")
	} else {
		defer file.Close()
		mimeType := getMimeType(path)
		c.SetContentType(mimeType)
	}
	// 更安全的写法，直接拷贝到响应流
	stat, err := file.Stat()
	if err != nil {
		c.String(500, "Read file stat failed")
		return
	}
	c.Status(200)
	c.Response.Header.SetContentLength(int(stat.Size()))
	_, err = io.Copy(c, file)
	if err != nil {
		c.String(500, "Read file failed")
		return
	}
}
