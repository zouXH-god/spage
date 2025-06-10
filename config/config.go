package config

import (
	"embed"
	"errors"
	"os"
	"path/filepath"

	"github.com/LiteyukiStudio/spage/constants"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const ()

var (
	ServerPort string
	// 服务器端口 Server Port

	Mode = constants.ModeProd
	// 运行模式，支持dev和prod
	// Running Mode, support dev and prod

	JwtSecret string
	// JWT密钥 JWT Secret

	FrontEndURL string
	// 前端URL Frontend URL

	LogLevel = "info"
	// 日志级别 Log Level

	AdminUsername = "admin"
	// 管理员用户名 Admin Username

	AdminPassword = "admin"
	// 管理员密码 Admin Password

	EmailEnable   bool   // 是否启用邮箱发送 Enable Email Sending
	EmailUsername string // 邮箱用户名 Email Username
	EmailAddress  string // 邮箱地址 Email Address
	EmailHost     string // 邮箱服务器地址 Email Server Address
	EmailPort     string // 邮箱服务器端口 Email Server Port
	EmailPassword string // 邮箱密码 Email Password
	EmailSSL      bool   // 是否启用SSL Enable SSL

	PageLimit int = 40
	// 每页显示的文章数量，默认为40
	// Number of articles displayed per page, default 40

	CaptchaType = constants.CaptchaTypeDisable
	// 验证码类型，支持turnstile、recaptcha和hcaptcha
	// Captcha Type, support turnstile, recaptcha and hcaptcha

	CaptchaSiteKey   string // reCAPTCHA v3的站点密钥
	CaptchaSecretKey string // reCAPTCHA v3的密钥
	CaptchaUrl       string // for mcaptcha

	TokenExpireTime = 3600 * 24
	// session过期时间，单位秒
	// session expiration time, in seconds

	RefreshTokenExpireTime = 3600 * 144
	// 刷新token过期时间，单位秒
	// refresh token expiration time, in seconds

	// CommitHash 构件时注入的git commit hash
	CommitHash = "develop"
	// git commit hash 构建时注入
	// git commit hash injected at build time

	BuildTime = "0000-00-00 00:00:00" // 构建时间 Build Time
	Version   = "0.0.0"         // 版本号 Version

	ReleaseSavePath = "data/releases"

	//go:embed config.example.yaml
	configExample embed.FS
)

// InitConfig 初始化配置文件
// Initialize the configuration file
func InitConfig() error {
	configPath := "config.yaml"
	// 目标配置文件路径
	// Target configuration file path

	// 如果 config.yaml 已存在，直接返回
	// If config.yaml already exists, return directly
	if _, err := os.Stat(configPath); err == nil {
		return nil
	}

	// 读取嵌入的示例配置
	// Read the embedded example configuration
	data, err := configExample.ReadFile("config.example.yaml")
	if err != nil {
		return errors.New("failed to read embedded config: " + err.Error())
	}

	// 确保目录存在（如果 config.yaml 不在当前目录）
	// Ensure the directory exists (if config.yaml is not in the current directory)
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return errors.New("failed to create config directory: " + err.Error())
	}

	// 写入 config.yaml
	// Write config.yaml
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return errors.New("failed to write config file: " + err.Error())
	}

	return nil
}

// Init 初始化配置文件和常量
// Initialize the configuration file and constants
func Init() error {
	// 设置配置文件路径
	// Set the configuration file path
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// 读取配置文件
	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			err := InitConfig()
			if err != nil {
				return err
			}
			return errors.New("config file not found")
		}
	}
	// 初始化配置常量
	// Initialize configuration constants
	ServerPort = GetString("server.port", "8888")
	FrontEndURL = GetString("frontend.url", "http://localhost:5173")
	Mode = GetString("mode", "prod")
	LogLevel = GetString("log.level", "info")

	// Admin配置项
	// Admin configuration items
	AdminUsername = GetString("admin.username", "admin")
	AdminPassword = GetString("admin.password", "admin")

	// Captcha配置项
	// Captcha configuration items
	CaptchaType = GetString("captcha.type", CaptchaType)
	CaptchaSiteKey = GetString("captcha.site-key", "")
	CaptchaSecretKey = GetString("captcha.secret-key", "")
	CaptchaUrl = GetString("captcha.url", "")

	// Email配置项
	// Email configuration items
	EmailEnable = GetBool("email.enable", false)
	EmailUsername = GetString("email.username", "")
	EmailAddress = GetString("email.address", "")
	EmailHost = GetString("email.host", "")
	EmailPort = GetString("email.port", "465")
	EmailPassword = GetString("email.password", "")
	EmailSSL = GetBool("email.ssl", true)

	// File存储配置项
	ReleaseSavePath = GetString("file.release-path", "data/releases")

	// 分页查询限制
	// Pagination query limit
	PageLimit = GetInt("page-limit", PageLimit)

	// Session过期时间
	// Session expiration time
	TokenExpireTime = GetInt("token.expire", TokenExpireTime)
	RefreshTokenExpireTime = GetInt("token.refresh-expire", RefreshTokenExpireTime)
	JwtSecret = GetString("token.secret", "none-secret")

	// 从启动参数拿取一些配置项mode frontend-url
	// Get some configuration items from the startup parameters mode frontend-url
	argsMap := Cmd.GetArgsMap(os.Args[1:])
	queryKeys := []string{"mode", "frontend-url", "port"}
	for _, key := range queryKeys {
		if value, ok := argsMap[key]; ok {
			switch key {
			case "mode":
				Mode = value
			case "port":
				ServerPort = value
			case "frontend-url":
				FrontEndURL = value
			}
		}
	}
	logrus.Info("Configuration loaded successfully, mode: ", Mode)

	// 设置日志级别
	// Set log level
	logLevel, err := logrus.ParseLevel(LogLevel)
	if err != nil {
		logrus.Error("Invalid log level, using default level: info")
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)
	logrus.Info("Log level set to: ", logLevel)
	logrus.Debugln("LogLevel is: ", LogLevel)
	// 其他配置项合法校验流程
	// Other configuration item validation process
	// ...
	return nil
}

// Get 返回配置项的值，如果不存在则返回默认值
// Return the value of the configuration item, or the default value if it does not exist
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
// Return the string value of the configuration item
func GetString(key string, defaultValue ...string) string {
	if len(defaultValue) > 0 {
		return Get(key, defaultValue[0])
	}
	return viper.GetString(key)
}

// GetInt 返回配置项的整数值
// Return the integer value of the configuration item
func GetInt(key string, defaultValue ...int) int {
	if len(defaultValue) > 0 {
		return Get(key, defaultValue[0])
	}
	return viper.GetInt(key)
}

// GetBool 返回配置项的布尔值
// Return the boolean value of the configuration item
func GetBool(key string, defaultValue ...bool) bool {
	if len(defaultValue) > 0 {
		return Get(key, defaultValue[0])
	}
	return viper.GetBool(key)
}

// GetFloat64 返回配置项的浮点数值
// Return the floating-point value of the configuration item
func GetFloat64(key string, defaultValue ...float64) float64 {
	if len(defaultValue) > 0 {
		return Get(key, defaultValue[0])
	}
	return viper.GetFloat64(key)
}

// GetStringSlice 返回配置项的字符串切片值
// Return the string slice value of the configuration item
func GetStringSlice(key string, defaultValue ...[]string) []string {
	if len(defaultValue) > 0 {
		return Get(key, defaultValue[0])
	}
	return viper.GetStringSlice(key)
}
