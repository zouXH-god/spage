package config

import (
	"testing"
)

// TestCmdUtilsType_GetArgsMap 测试GetArgsMap函数
// Test the GetArgsMap function
func TestCmdUtilsType_GetArgsMap(t *testing.T) {
	args := []string{"--mode=dev", "--port=8080", "--name=spage"}
	expected := map[string]string{
		"mode": "dev",
		"port": "8080",
		"name": "spage",
	}

	cmdUtils := CmdUtilsType{}
	result := cmdUtils.GetArgsMap(args)

	if len(result) != len(expected) {
		t.Errorf("Expected %d arguments, got %d", len(expected), len(result))
	}

	for key, value := range expected {
		if result[key] != value {
			t.Errorf("Expected %s for key %s, got %s", value, key, result[key])
		}
	}
}
