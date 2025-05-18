package config

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const (
	ModeDev  = "dev"
	ModeProd = "prod"
)

var (
	ServerPort string
	Mode       = "prod"
	JwtSecret  string

	AdminUsername = "admin" // 管理员用户名
	AdminPassword = "admin" // 管理员密码

	EmailEnable   bool   // 是否启用邮箱发送
	EmailUsername string // 邮箱用户名
	EmailAddress  string // 邮箱地址
	EmailHost     string // 邮箱服务器地址
	EmailPort     string // 邮箱服务器端口
	EmailPassword string // 邮箱密码
	EmailSSL      bool   // 是否启用SSL

	TokenExpireTime        = 3600 * 24  // session过期时间，单位秒
	RefreshTokenExpireTime = 3600 * 144 // 刷新token过期时间，单位秒

	// CommitHash 构件时注入的git commit hash
	CommitHash = "develop"             // git commit hash 构建时注入
	BuildTime  = "0000-00-00 00:00:00" // 构建时间
)

func Init() error {
	// 设置配置文件路径
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return errors.New("config file not found")
		}
	}
	// 初始化配置常量
	ServerPort = GetString("server.port", "8888")

	Mode = GetString("mode", "prod")
	// Admin配置项
	AdminUsername = GetString("admin.username", "admin")
	AdminPassword = GetString("admin.password", "admin")
	// Email配置项
	EmailEnable = GetBool("email.enable", false)
	EmailUsername = GetString("email.username", "")
	EmailAddress = GetString("email.address", "")
	EmailHost = GetString("email.host", "")
	EmailPort = GetString("email.port", "465")
	EmailPassword = GetString("email.password", "")
	EmailSSL = GetBool("email.ssl", true)
	// Session过期时间
	TokenExpireTime = GetInt("token.expire", TokenExpireTime)
	RefreshTokenExpireTime = GetInt("token.refresh-expire", RefreshTokenExpireTime)
	JwtSecret = GetString("token.secret", "none-secret")

	// 从启动参数拿取一些配置项mode frontend-url
	argsMap := Cmd.GetArgsMap(os.Args[1:])
	queryKeys := []string{"mode", "frontend-url", "port"}
	for _, key := range queryKeys {
		if value, ok := argsMap[key]; ok {
			switch key {
			case "mode":
				Mode = value
			case "port":
				ServerPort = value
			}
		}
	}
	logrus.Info("Configuration loaded successfully, mode: ", Mode)
	return nil
}

func Get[T any](key string, defaultValue T) T {
	if !viper.IsSet(key) {
		return defaultValue
	}

	value := viper.Get(key)
	if v, ok := value.(T); ok {
		return v
	}
	return defaultValue
}

// GetString 返回配置项的字符串值
func GetString(key string, defaultValue ...string) string {
	if len(defaultValue) > 0 {
		return Get(key, defaultValue[0])
	}
	return viper.GetString(key)
}

// GetInt 返回配置项的整数值
func GetInt(key string, defaultValue ...int) int {
	if len(defaultValue) > 0 {
		return Get(key, defaultValue[0])
	}
	return viper.GetInt(key)
}

// GetBool 返回配置项的布尔值
func GetBool(key string, defaultValue ...bool) bool {
	if len(defaultValue) > 0 {
		return Get(key, defaultValue[0])
	}
	return viper.GetBool(key)
}

// GetFloat64 返回配置项的浮点数值
func GetFloat64(key string, defaultValue ...float64) float64 {
	if len(defaultValue) > 0 {
		return Get(key, defaultValue[0])
	}
	return viper.GetFloat64(key)
}

// GetStringSlice 返回配置项的字符串切片值
func GetStringSlice(key string, defaultValue ...[]string) []string {
	if len(defaultValue) > 0 {
		return Get(key, defaultValue[0])
	}
	return viper.GetStringSlice(key)
}
