package models

import (
	"github.com/LiteyukiStudio/spage/agent/env"
	"github.com/LiteyukiStudio/spage/pkg/orm"
	"gorm.io/gorm"
	"strconv"
)

type Site struct {
	gorm.Model
	SiteId    uint                         `gorm:"not null"`
	ReleaseId uint                         `gorm:"not null"` // 用于和主控校验状态
	Hosts     orm.GenericJsonArray[string] `gorm:"type:json;default:'[]'"`
}

func (s *Site) GroupName() string {
	return env.SiteGroupPrefix + strconv.Itoa(int(s.SiteId))
}

func (s *Site) SiteIdString() string {
	return strconv.Itoa(int(s.SiteId))
}

func (s *Site) FileServerBase() string {
	return env.SitesRoot + s.SiteIdString()
}

func (s *Site) GenerateConfig() map[string]any {
	return map[string]any{
		"group": s.GroupName(),
		"match": []any{
			map[string]any{
				"host": s.Hosts,
			},
		},
		"handle": []any{
			map[string]any{
				"handler": "subroute",
				"routes": []any{
					map[string]any{
						"handle": []any{
							map[string]any{
								"handler": "file_server",
								"root":    s.FileServerBase(),
							},
						},
					},
				},
			},
		},
	}
}
