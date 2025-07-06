package caddy

import (
	"github.com/LiteyukiStudio/spage/agent/env"
	"resty.dev/v3"
)

// GetConfig 导出指定路径的配置
func GetConfig(path string) (any, error) {
	var result any
	client := resty.New()
	res, err := client.R().SetResult(&result).Get(env.CaddyApiEndpoint + "/config" + path)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, res.Err
	}
	return result, nil
}

// SetConfig 设置或替换对象；向数组追加
func SetConfig(path string, config any) error {
	client := resty.New()
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(config).
		Post(env.CaddyBin + "/config" + path)
	if err != nil {
		return err
	}
	if res.IsError() {
		return res.Err
	}
	return nil
}

// CreateConfig 创建新对象；插入到数组
func CreateConfig(path string, config any) error {
	client := resty.New()
	res, err := client.R().SetBody(config).Put(env.CaddyApiEndpoint + "/config" + path)
	if err != nil {
		return err
	}
	if res.IsError() {
		return res.Err
	}
	return nil
}

// UpdateConfig 替换现有的对象或数组元素
func UpdateConfig(path string, config any) error {
	client := resty.New()
	res, err := client.R().SetBody(config).Patch(env.CaddyApiEndpoint + "/config" + path)
	if err != nil {
		return err
	}
	if res.IsError() {
		return res.Err
	}
	return nil
}

// DeleteConfig 删除指定路径的值
func DeleteConfig(path string) error {
	client := resty.New()
	res, err := client.R().Delete(env.CaddyApiEndpoint + "/config" + path)
	if err != nil {
		return err
	}
	if res.IsError() {
		return res.Err
	}
	return nil
}
