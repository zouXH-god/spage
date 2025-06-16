package filedriver

import (
	"bytes"
	"io"
	"os"
	"path"

	"github.com/studio-b12/gowebdav"
)

type WebDAVClientDriver struct {
	client   *gowebdav.Client
	BasePath string
}

func NewWebDAVClientDriver(driverConfig *DriverConfig) *WebDAVClientDriver {
	c := gowebdav.NewClient(driverConfig.WebDavUrl, driverConfig.WebDavUser, driverConfig.WebDavPassword)
	return &WebDAVClientDriver{client: c, BasePath: driverConfig.BasePath}
}

func (d *WebDAVClientDriver) fullPath(p string) string {
	if d.BasePath == "" {
		return p
	}
	return path.Join(d.BasePath, p)
}

func (d *WebDAVClientDriver) Save(p string, r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return d.client.Write(d.fullPath(p), data, 0644)
}

func (d *WebDAVClientDriver) Open(p string) (io.ReadCloser, error) {
	data, err := d.client.Read(d.fullPath(p))
	if err != nil {
		return nil, err
	}
	return io.NopCloser(bytes.NewReader(data)), nil
}

func (d *WebDAVClientDriver) Delete(p string) error {
	return d.client.Remove(d.fullPath(p))
}

func (d *WebDAVClientDriver) Stat(p string) (os.FileInfo, error) {
	return d.client.Stat(d.fullPath(p))
}

func (d *WebDAVClientDriver) ListDir(p string) ([]os.FileInfo, error) {
	return d.client.ReadDir(d.fullPath(p))
}
