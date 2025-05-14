package config

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	ServerPort string
	Mode       string
	JwtSecret  string
	//MessageSavingDays      = 30 // 消息在数据库保存时间，单位天
	//MessageHangingSeconds  = 10 // 消息挂起时间，单位秒
	//MessageResponseTimeout = 60 // 长轮询消息响应超时时间，单位秒

	// CommitHash 构件时注入的git commit hash
	CommitHash = "develop"             // git commit hash 构建时注入
	BuildTime  = "0000-00-00 00:00:00" // 构建时间
)

func init() {
	// 设置配置文件路径
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	// 初始化配置常量
	ServerPort = GetString("serverPort", "8888")
	JwtSecret = GetString("jwtSecret", "none-secret")
	Mode = GetString("mode", "prod")
	logrus.Info("Configuration loaded successfully, mode: ", Mode)
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
