package utils

import (
	"os"
	"path"
	"testing"

	"github.com/LiteyukiStudio/spage/config"
)

func TestFileHash(t *testing.T) {
	testFile := path.Join(config.UploadsPath, "test.txt")
	testDir := path.Dir(testFile)
	testContent := []byte("hello world")

	// 递归创建测试文件所在的目录
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("无法创建测试目录: %v", err)
	}

	// 创建测试文件
	err = os.WriteFile(testFile, testContent, 0644)
	if err != nil {
		t.Fatalf("无法创建测试文件: %v", err)
	}
	defer os.Remove(testFile)   // 测试结束后删除文件
	defer os.RemoveAll(testDir) // 测试结束后删除目录

	// 计算哈希
	hash, err := FileHash(testFile)
	if err != nil {
		t.Fatalf("FileHash 错误: %v", err)
	}

	// 计算期望哈希
	expected := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9" // "hello world" 的 SHA256

	if hash != expected {
		t.Errorf("哈希不匹配，got: %s, want: %s", hash, expected)
	}
}
