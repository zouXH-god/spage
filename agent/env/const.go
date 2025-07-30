package env

import "github.com/LiteyukiStudio/spage/pkg/utils"

var (
	LogLevel           = utils.GetEnv("LOG_LEVEL", "info")                           // 日志级别
	AgentApiToken      = utils.GetEnv("AGENT_API_TOKEN", "")                         // Agent API Token，用于主控和代理通信的身份验证
	SitesRoot          = utils.GetEnv("SITES_ROOT", "./data/sites/")                 // 站点根目录
	SiteGroupPrefix    = utils.GetEnv("SITE_GROUP_PREFIX", "site_")                  // 站点组前缀
	CaddyBin           = utils.GetEnv("CADDY_BIN", "caddy")                          // Caddy二进制文件路径
	CaddyCommandArgs   = utils.GetEnv("CADDY_COMMAND_ARGS", "run")                   // Caddy命令行参数
	CaddyApiEndpoint   = utils.GetEnv("CADDY_API_ENDPOINT", "http://localhost:2019") // Caddy API 端点地址
	CaddyCheckInterval = utils.GetEnvInt("CADDY_CHECK_INTERVAL", 30)                 // Caddy健康检查间隔（秒
)
