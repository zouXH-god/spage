package utils

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// 域名相关工具

// NormalizeDomain 规范化域名，确保格式正确
func NormalizeDomain(domain string) string {
	// 移除协议前缀
	domain = strings.TrimPrefix(domain, "http://")
	domain = strings.TrimPrefix(domain, "https://")
	// 移除路径和参数
	if idx := strings.Index(domain, "/"); idx > 0 {
		domain = domain[:idx]
	}
	// 移除端口号
	if idx := strings.Index(domain, ":"); idx > 0 {
		domain = domain[:idx]
	}
	return domain
}

// LookupTXTRecords 查询域名的TXT记录
func LookupTXTRecords(domain string) ([]string, error) {
	if domain == "" {
		return nil, errors.New("域名不能为空")
	}
	records, err := net.LookupTXT(domain)
	if err != nil {
		return nil, fmt.Errorf("查询TXT记录失败: %w", err)
	}
	return records, nil
}

// ContainsToken 检查TXT记录中是否包含指定令牌
func ContainsToken(records []string, token string) bool {
	for _, record := range records {
		if strings.TrimSpace(record) == token {
			return true
		}
	}
	return false
}

// VerifyDomainOwnership 验证域名所有权
func VerifyDomainOwnership(domain, token string) (bool, error) {
	normalizedDomain := NormalizeDomain(domain)
	records, err := LookupTXTRecords(normalizedDomain)
	if err != nil {
		return false, err
	}
	return ContainsToken(records, token), nil
}
