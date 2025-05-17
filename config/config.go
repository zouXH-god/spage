package config

import (
	"errors"
	"github.com/LiteyukiStudio/spage/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const (
	ModeDev  = "dev"
	ModeProd = "prod"
)

var (
	ServerPort  string
	Mode        = "prod"
	JwtSecret   string
	FrontendUrl string = "http://localhost:5173" // 前端开发服务器地址，仅在开发模式下使用

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
	ServerPort = GetString("serverPort", "8888")
	JwtSecret = GetString("jwtSecret", "none-secret")
	Mode = GetString("mode", "prod")

	// 从启动参数拿取一些配置项mode frontend-url
	argsMap := utils.Cmd.GetArgsMap(os.Args[1:])
	queryKeys := []string{"mode", "frontend-url", "port"}
	for _, key := range queryKeys {
		if value, ok := argsMap[key]; ok {
			switch key {
			case "mode":
				Mode = value
			case "port":
				ServerPort = value
			case "frontend-url":
				FrontendUrl = value
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
