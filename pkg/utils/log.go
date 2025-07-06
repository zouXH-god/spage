package utils

import (
	"github.com/sirupsen/logrus"
	"strings"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,                         // 启用完整时间戳
		TimestampFormat: "2006/01/02 15:04:05.000000", // 时间格式
	})
	logrus.SetLevel(getLogLevel(GetEnv("LOG_LEVEL", "info")))
}

func getLogLevel(level string) logrus.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel // 默认级别
	}
}
