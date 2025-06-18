package models

import "gorm.io/gorm"

type OIDCConfig struct {
	gorm.Model
	AdminGroups      []string `gorm:"type:json;column:admin_groups;default:'[]'"`        // 平台管理员组，默认为：[]string{}，*为匹配所有组，储存为逗号分隔的字符串
	AllowedGroups    []string `gorm:"type:json;column:allowed_groups;default:'[\"*\"]'"` // 允许登录的组，默认为：[]string{"*"}，*为匹配所有组，储存为逗号分隔的字符串
	ClientID         string   `gorm:"column:client_id"`                                  // 客户端ID
	ClientSecret     string   `gorm:"column:client_secret"`                              // 客户端密钥
	DisplayName      string   `gorm:"column:display_name"`                               // 显示名称，例如：轻雪通行证
	GroupsClaim      *string  `gorm:"default:groups"`                                    // 组声明，默认为："groups"
	Icon             *string  `gorm:"column:icon"`                                       // 图标url，为空则使用内置默认图标
	OidcDiscoveryURL string   `gorm:"column:oidc_discovery_url"`                         // OpenID自动发现URL，例如 ：https://pass.liteyuki.icu/.well-known/openid-configuration
}

// TableName 重写表名
func (OIDCConfig) TableName() string {
	return "oidc_configs"
}
