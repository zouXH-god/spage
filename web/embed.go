package web

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var StaticFS embed.FS

// GetFS 返回前端构建文件的子文件系统
func GetFS() (fs.FS, error) {
	return fs.Sub(StaticFS, "dist")
}
