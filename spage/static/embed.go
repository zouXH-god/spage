package static

import (
	"embed"
)

//go:embed dist/*
var WebFS embed.FS

//go:embed assets/*
var AssetsFS embed.FS
