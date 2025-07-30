package config

import (
	"os"
	"strconv"
	"strings"
)

type CmdUtilsType struct {
	data map[string]string
}

var Cmd = CmdUtilsType{}

// GetArgsSlice 输入--mode=dev --port=8080时，解析为["mode","port"]
// When inputting --mode=dev --port=8080, it is parsed as ["mode","port"]
func (*CmdUtilsType) GetArgsSlice() []string {
	return os.Args[1:]
}

// GetArgsMap 输入--mode=dev --port=8080时，解析为map[string]string{"mode":"dev", "port":"8080"}
// When inputting --mode=dev --port=8080, it is parsed as map[string]string{"mode":"dev", "port":"8080"}
func (c *CmdUtilsType) GetArgsMap(args []string) map[string]string {
	argsMap := make(map[string]string)
	for _, arg := range args {
		if len(arg) > 2 && arg[:2] == "--" {
			parts := splitArg(arg)
			if len(parts) == 2 {
				argsMap[parts[0][2:]] = parts[1]
			}
		}
	}
	return argsMap
}

// ParseArgs 解析命令行参数为 map
func (c *CmdUtilsType) ParseArgs() {
	args := os.Args[1:]
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			parts := splitArg(arg)
			if len(parts) == 2 {
				c.data[parts[0][2:]] = parts[1]
			}
		}
	}
}

func (c *CmdUtilsType) GetString(key string, defaultValue string) string {
	value := c.data[key]
	if value == "" {
		return defaultValue
	}
	return value
}

func (c *CmdUtilsType) GetInt(key string, defaultValue int) int {
	value := c.data[key]
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}

// splitArg 输入--mode=dev时，解析为["mode","dev"]
// When inputting --mode=dev, it is parsed as ["mode","dev"]
func splitArg(arg string) []string {
	parts := make([]string, 2)
	splitIndex := -1
	for i, char := range arg {
		if char == '=' {
			splitIndex = i
			break
		}
	}
	if splitIndex != -1 {
		parts[0] = arg[:splitIndex]
		parts[1] = arg[splitIndex+1:]
	}
	return parts
}
