package store

import (
	"errors"
	"fmt"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/constants"
	"github.com/LiteyukiStudio/spage/spage/models"
	"github.com/LiteyukiStudio/spage/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"plugin"
	"runtime"
)

var DB *gorm.DB

// DBConfig 数据库配置结构体
type DBConfig struct {
	Driver   string // 数据库驱动类型，例如 "sqlite" 或 "postgres" Database driver type, e.g., "sqlite" or "postgres"
	Path     string // SQLite 路径 SQLite path
	Host     string // PostgreSQL 主机名 PostgreSQL hostname
	Port     int    // PostgreSQL 端口 PostgreSQL port
	User     string // PostgreSQL 用户名 PostgreSQL username
	Password string // PostgreSQL 密码 PostgreSQL password
	DBName   string // PostgreSQL 数据库名 PostgreSQL database name
	SSLMode  string // PostgreSQL SSL 模式 PostgreSQL SSL mode
}

// loadDBConfig 从配置文件加载数据库配置
func loadDBConfig() DBConfig {
	return DBConfig{
		Driver:   config.GetString("database.driver", "sqlite"),
		Path:     config.GetString("database.path", "./data/data.db"),
		Host:     config.GetString("database.host", "postgres"),
		Port:     config.GetInt("database.port", 5432),
		User:     config.GetString("database.user", "spage"),
		Password: config.GetString("database.password", "spage"),
		DBName:   config.GetString("database.dbname", "spage"),
		SSLMode:  config.GetString("database.sslmode", "disable"),
	}
}

// Init 手动初始化数据库连接
func Init() error {
	dbConfig := loadDBConfig()
	// 创建通用的 GORM 配置
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	var err error

	switch dbConfig.Driver {
	case "postgres":
		if DB, err = initPostgres(dbConfig, gormConfig); err != nil {
			return fmt.Errorf("postgres initialization failed: %w", err)
		}
		logrus.Infoln("postgres initialization succeeded", dbConfig)
	case "sqlite":
		if DB, err = InitSQLiteDynamic(dbConfig, gormConfig); err != nil {
			return fmt.Errorf("sqlite initialization failed: %w", err)
		}
		logrus.Infoln("sqlite initialization succeeded", dbConfig)
	default:
		return errors.New("unsupported database driver, only sqlite and postgres are supported")
	}

	// 迁移模型
	if err = models.Migrate(DB); err != nil {
		logrus.Error("Failed to migrate models:", err)
		return err
	}
	// 执行初始化数据
	// 创建管理员账户
	hashedPassword, err := utils.Password.HashPassword(config.AdminPassword, config.JwtSecret)
	if err != nil {
		logrus.Error("Failed to hash password:", err)
		return err
	}
	user := &models.User{
		Name:     config.AdminUsername,
		Password: &hashedPassword,
		Role:     constants.RoleAdmin,
	}
	if err = User.UpdateSystemAdmin(user); err != nil {
		logrus.Error("Failed to update admin user:", err)
		return err
	}
	return nil
}

// initPostgres 初始化PostgreSQL连接
func initPostgres(config DBConfig, gormConfig *gorm.Config) (db *gorm.DB, err error) {
	if config.Host == "" || config.User == "" || config.Password == "" || config.DBName == "" {
		err = errors.New("PostgreSQL configuration is incomplete: host, user, password, and dbname are required")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err = gorm.Open(postgres.Open(dsn), gormConfig)
	return
}

func InitSQLiteDynamic(config DBConfig, gormConfig *gorm.Config) (*gorm.DB, error) {
	// 根据操作系统选择插件后缀
	var ext string
	switch runtime.GOOS {
	case "darwin":
		ext = "dylib"
	case "linux":
		ext = "so"
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
	// 构建插件路径
	pluginPath := fmt.Sprintf("./build/sqlite.%s", ext)
	// 加载 SQLite 插件
	plug, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load SQLite plugin: %w", err)
	}
	// 查找导出的 Init 函数
	symbol, err := plug.Lookup("InitSQLite")
	logrus.Infof("Found symbol: %v", symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to find Init function in plugin: %w", err)
	}
	// 调试符号类型
	// 转换为函数类型
	initFunc, ok := symbol.(func(string, *gorm.Config) (*gorm.DB, error))
	if !ok {
		return nil, errors.New("invalid Init function signature in plugin")
	}
	// 调用插件中的 Init 函数
	return initFunc(config.Path, gormConfig)
}
