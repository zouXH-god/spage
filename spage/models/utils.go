package models

import (
	"database/sql/driver"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/json"
)

// JsonObject 是一个表示 JSON 对象数组的类型
type JsonObject []map[string]any

// Scan 实现 sql.Scanner 接口，用于从数据库读取 JSON 数据到结构体
func (j *JsonObject) Scan(value interface{}) error {
	if value == nil {
		*j = JsonObject{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New("无法将类型转换为 JsonObject")
	}

	return json.Unmarshal(bytes, j)
}

// Value 实现 driver.Valuer 接口，用于将结构体数据保存为 JSON 到数据库
func (j JsonObject) Value() (driver.Value, error) {
	if j == nil {
		return "[]", nil
	}
	return json.Marshal(j)
}

type GenericJsonArray[T any] []T

// Scan 实现 sql.Scanner 接口
func (j *GenericJsonArray[T]) Scan(value interface{}) error {
	if value == nil {
		*j = GenericJsonArray[T]{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New("无法将类型转换为 GenericJsonArray")
	}

	return json.Unmarshal(bytes, j)
}

// Value 实现 driver.Valuer 接口
func (j GenericJsonArray[T]) Value() (driver.Value, error) {
	if j == nil {
		return "[]", nil
	}
	return json.Marshal(j)
}
