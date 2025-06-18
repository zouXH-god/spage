package models

import "gorm.io/gorm"

type OIDCConfig struct {
	gorm.Model
	AdminGroups []string `gorm:"type:json;column:admin_groups;default:'[]'"` // 平台管理员组，默认为：[]string{}，*为匹配所有组，储存为逗号分隔的字符串
	// Admin groups, default is: []string{}, * matches all groups, stored as a comma-separated string
	AllowedGroups []string `gorm:"type:json;column:allowed_groups;default:'[\"*\"]'"` // 允许登录的组，默认为：[]string{"*"}，*为匹配所有组，储存为逗号分隔的字符串
	// Allowed groups for login, default is: []string{"*"}, * matches all groups, stored as a comma-separated string
	ClientID string `gorm:"column:client_id"` // 客户端ID
	// Client ID
	ClientSecret string `gorm:"column:client_secret"` // 客户端密钥
	// Client Secret
	DisplayName string `gorm:"column:display_name"` // 显示名称，例如：轻雪通行证
	// Display name, e.g., Light Snow Passport
	GroupsClaim *string `gorm:"default:groups"` // 组声明，默认为："groups"
	// Groups claim, default is: "groups"
	Icon *string `gorm:"column:icon"` // 图标url，为空则使用内置默认图标
	// Icon URL, if empty use the built-in default icon
	OidcDiscoveryURL string `gorm:"column:oidc_discovery_url"` // OpenID自动发现URL，例如 ：https://pass.liteyuki.icu/.well-known/openid-configuration
	// OpenID auto-discovery URL, e.g., https://pass.liteyuki.icu/.well-known/openid-configuration
}

// TableName 重写表名
// Rewrite table name
func (OIDCConfig) TableName() string {
	return "oidc_configs"
}
