package filedriver

import (
	"io"
	"os"
	"path/filepath"
)

type LocalDriver struct {
	BasePath string
}

func NewLocalDriver(driverConfig *DriverConfig) *LocalDriver {
	return &LocalDriver{
		BasePath: driverConfig.BasePath,
	}
}

func (d *LocalDriver) fullPath(path string) string {
	return filepath.Join(d.BasePath, path)
}

func (d *LocalDriver) Save(path string, r io.Reader) error {
	full := d.fullPath(path)
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		return err
	}
	f, err := os.Create(full)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, r)
	return err
}

func (d *LocalDriver) Open(path string) (io.ReadCloser, error) {
	return os.Open(d.fullPath(path))
}

func (d *LocalDriver) Delete(path string) error {
	return os.Remove(d.fullPath(path))
}

func (d *LocalDriver) Stat(path string) (os.FileInfo, error) {
	return os.Stat(d.fullPath(path))
}

func (d *LocalDriver) ListDir(path string) ([]os.FileInfo, error) {
	entries, err := os.ReadDir(d.fullPath(path))
	if err != nil {
		return nil, err
	}
	infos := make([]os.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}
	return infos, nil
}
