package utils

import (
	"regexp"
	"strings"
)

// IsValidEntityName 检查字符串是否符合Git仓库命名规范
func IsValidEntityName(name string) bool {
	// 仓库名不能为空
	if name == "" {
		return false
	}
	// 仓库名不能是 "." 或 ".."
	if name == "." || name == ".." {
		return false
	}
	// 仓库名不能以 "." 开头或结尾
	if strings.HasPrefix(name, ".") || strings.HasSuffix(name, ".") {
		return false
	}
	// 仓库名不能包含连续的 "."
	if strings.Contains(name, "..") {
		return false
	}
	// 仓库名只能包含字母、数字、连字符、下划线和单个句点
	validPattern := regexp.MustCompile(`^[a-zA-Z0-9\-_.]+$`)
	if !validPattern.MatchString(name) {
		return false
	}
	// 检查是否包含非法字符
	invalidChars := []string{
		" ", "\\", "/", ":", "*", "?", "\"", "<", ">", "|", "@", "{", "}", "[", "]",
	}
	for _, char := range invalidChars {
		if strings.Contains(name, char) {
			return false
		}
	}
	return true
}
