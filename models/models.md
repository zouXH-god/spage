# 数据模型汇总

## User 用户模型

| 字段名           | 类型              | GORM标签                                   | 注释                          |
|---------------|-----------------|------------------------------------------|-----------------------------|
| Model         | gorm.Model      |                                          | 内嵌GORM基础模型                  |
| Name          | string          | `gorm:"not null;unique"`                 | 用户的唯一名称                     |
| DisplayName   | *string         | `gorm:"column:display_name"`             | 用户的显示名称                     |
| Email         | *string         | `gorm:"unique"`                          | 用户的电子邮件地址(仅用户邮箱唯一，用于OIDC认证) |
| Description   | string          | `gorm:"default:'No description.'"`       | 用户描述                        |
| Avatar        | *string         | `gorm:"column:avatar"`                   | 头像URL，留空则使用Gravatar         |
| Role          | string          | `gorm:"not null;default:member"`         | 用户的全局角色                     |
| Organizations | []*Organization | `gorm:"many2many:organization_members;"` | 用户所属的组织                     |
| ProjectLimit  | int             | `gorm:"default:0"`                       | 用户的项目限制，0表示无限制              |
| Language      | string          | `gorm:"default:'zh-cn'"`                 | 用户语言，默认为中文                  |
| Flag          | string          | `gorm:"default:'0'"`                     | 系统管理员的另一个标志位                |
| Password      | *string         | `gorm:"column:password"`                 | 用户密码(哈希值)，仅用于本地认证           |

表名: `users`

## Organization 组织模型

| 字段名          | 类型         | GORM标签                                   | 注释                    |
|--------------|------------|------------------------------------------|-----------------------|
| Model        | gorm.Model |                                          | 内嵌GORM基础模型            |
| Name         | string     | `gorm:"not null;unique"`                 | 组织的唯一名称               |
| DisplayName  | *string    | `gorm:"column:display_name"`             | 组织的显示名称               |
| Email        | *string    | `gorm:"column:email"`                    | 组织的电子邮件地址             |
| Description  | string     | `gorm:"default:'No description.'"`       | 组织描述                  |
| Avatar       | *string    | `gorm:"column:avatar"`                   | 头像URL，留空则使用Gravatar   |
| Members      | []*User    | `gorm:"many2many:organization_members;"` | 组织成员(包含创建者)           |
| Owners       | []User     | `gorm:"many2many:organization_owners;"`  | 组织所有者(无反向关系，包含创建者)    |
| ProjectLimit | int        | `gorm:"default:0"`                       | 组织的项目限制，0:遵循策略，-1:无限制 |

表名: `organizations`

## Project 项目模型

| 字段名         | 类型         | GORM标签                             | 注释                         |
|-------------|------------|------------------------------------|----------------------------|
| Model       | gorm.Model |                                    | 内嵌GORM基础模型                 |
| Name        | string     | `gorm:"not null;unique"`           | 项目的唯一名称                    |
| DisplayName | *string    | `gorm:"column:display_name"`       | 项目的显示名称                    |
| Description | string     | `gorm:"default:'No description.'"` | 项目描述                       |
| OwnerID     | uint       | `gorm:"not null"`                  | 所有者ID(用户ID或组织ID)           |
| OwnerType   | string     | `gorm:"not null"`                  | 所有者类型，可以是user或organization |
| Owners      | []User     | `gorm:"many2many:project_owners;"` | 项目所有者(无反向关系)               |
| SiteLimit   | int        | `gorm:"default:0"`                 | 项目的站点限制，0:遵循策略，-1:无限制      |

表名: `projects`

## File 文件模型

| 字段名   | 类型         | GORM标签              | 注释               |
|-------|------------|---------------------|------------------|
| Model | gorm.Model |                     | 内嵌GORM基础模型       |
| ID    | uint       | `gorm:"primaryKey"` | 文件ID             |
| Path  | string     | `gorm:"not null"`   | 文件路径，相较于根目录的相对路径 |

表名: `projects`

## OIDCConfig OIDC配置模型

| 字段名              | 类型         | GORM标签                                                     | 注释                                                                          |
|------------------|------------|------------------------------------------------------------|-----------------------------------------------------------------------------|
| Model            | gorm.Model |                                                            | 内嵌GORM基础模型                                                                  |
| AdminGroups      | []string   | `gorm:"type:json;column:admin_groups;default:'[]'"`        | 平台管理员组，默认为：[]string{}，*为匹配所有组，储存为逗号分隔的字符串                                   |
| AllowedGroups    | []string   | `gorm:"type:json;column:allowed_groups;default:'[\"*\"]'"` | 允许登录的组，默认为：[]string{"*"}，*为匹配所有组，储存为逗号分隔的字符串                                |
| ClientID         | string     | `gorm:"column:client_id"`                                  | 客户端ID                                                                       |
| ClientSecret     | string     | `gorm:"column:client_secret"`                              | 客户端密钥                                                                       |
| DisplayName      | string     | `gorm:"column:display_name"`                               | 显示名称，例如：轻雪通行证                                                               |
| GroupsClaim      | *string    | `gorm:"default:groups"`                                    | 组声明，默认为："groups"                                                            |
| Icon             | *string    | `gorm:"column:icon"`                                       | 图标url，为空则使用内置默认图标                                                           |
| OidcDiscoveryURL | string     | `gorm:"column:oidc_discovery_url"`                         | OpenID自动发现URL，例如：https://pass.liteyuki.icu/.well-known/openid-configuration |

表名: `oidc_configs`

## Site 站点模型

| 字段名         | 类型         | GORM标签                                                                     | 注释           |
|-------------|------------|----------------------------------------------------------------------------|--------------|
| Model       | gorm.Model |                                                                            | 内嵌GORM基础模型   |
| Name        | string     | `gorm:"unique"`                                                            | 站点名称         |
| Description | string     | `gorm:"size:255"`                                                          | 站点描述         |
| ProjectID   | uint       | `gorm:"not null"`                                                          | 所属项目ID       |
| Project     | Project    | `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` | 所属项目         |
| SubDomain   | string     | `gorm:"unique;size:255"`                                                   | 子域前缀         |
| Domains     | []string   | `gorm:"type:json;default:'[]'"`                                            | 允许的域名，json格式 |

表名: `sites`

## SiteRelease 站点发布模型

| 字段名    | 类型         | GORM标签                                                                   | 注释         |
|--------|------------|--------------------------------------------------------------------------|------------|
| Model  | gorm.Model |                                                                          | 内嵌GORM基础模型 |
| SiteID | uint       | `gorm:"not null"`                                                        | 站点ID       |
| Site   | Site       | `gorm:"foreignKey:SiteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`  | 所属站点       |
| Tag    | string     | `gorm:"not null"`                                                        | 版本标签       |
| FileID | uint       | `gorm:"not null"`                                                        | 版本文件ID     |
| File   | File       | `gorm:"foreignKey:FileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` | 版本文件       |

表名: `site_releases`