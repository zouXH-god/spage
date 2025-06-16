package filedriver

import (
	"fmt"
	"github.com/LiteyukiStudio/spage/constants"
	"io"
	"os"
)

type FileDriver interface {
	Save(path string, r io.Reader) error
	Open(path string) (io.ReadCloser, error)
	Delete(path string) error
	Stat(path string) (os.FileInfo, error)
	ListDir(path string) ([]os.FileInfo, error)
}

type DriverConfig struct {
	Type           string `mapstructure:"file.driver.type"`
	BasePath       string `mapstructure:"file.driver.base_path"`
	WebDavUrl      string `mapstructure:"file.driver.webdav.url"`
	WebDavUser     string `mapstructure:"file.driver.webdav.user"`
	WebDavPassword string `mapstructure:"file.driver.webdav.password"`
}

func GetFileDriver(driverConfig *DriverConfig) (FileDriver, error) {
	switch driverConfig.Type {
	case constants.FileDriverLocal:
		return NewLocalDriver(driverConfig), nil
	case constants.FileDriverWebdav:
		return NewWebDAVClientDriver(driverConfig), nil
	default:
		return nil, fmt.Errorf("unsupported file driver type: %s", driverConfig.Type)
	}
}
