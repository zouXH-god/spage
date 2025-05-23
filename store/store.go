package store

import (
	"errors"
	"fmt"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/constants"
	"github.com/LiteyukiStudio/spage/models"
	"github.com/LiteyukiStudio/spage/utils"
	"github.com/glebarez/sqlite" // 基于Go的 SQLite 驱动
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
)

var DB *gorm.DB

// DBConfig 数据库配置结构体
type DBConfig struct {
	Driver   string
	Path     string // SQLite 路径
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
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
		if err = initPostgres(dbConfig, gormConfig); err != nil {
			return fmt.Errorf("postgres initialization failed: %w", err)
		}
	case "sqlite":
		if err = initSQLite(dbConfig, gormConfig); err != nil {
			return fmt.Errorf("sqlite initialization failed: %w", err)
		}
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
	if err = User.UpdateSystemAdminUser(user); err != nil {
		logrus.Error("Failed to update admin user:", err)
		return err
	}
	return nil
}

// initPostgres 初始化PostgreSQL连接
func initPostgres(config DBConfig, gormConfig *gorm.Config) error {
	if config.Host == "" || config.User == "" || config.Password == "" || config.DBName == "" {
		return errors.New("PostgreSQL configuration is incomplete")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
	return err
}

// initSQLite 初始化SQLite连接
func initSQLite(config DBConfig, gormConfig *gorm.Config) error {
	if config.Path == "" {
		config.Path = "./data/data.db"
	}
	// 创建 SQLite 数据库文件的目录
	if err := os.MkdirAll(filepath.Dir(config.Path), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory for SQLite database: %w", err)
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(config.Path), gormConfig)
	return err
}
