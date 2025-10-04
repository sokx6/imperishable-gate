package utils

import "strings"

func removeEmptyAndDuplicate(strs []string) []string {
	seen := make(map[string]struct{}) // 使用 map 来去重
	var result []string

	for _, s := range strs {
		trimmed := strings.TrimSpace(s) // 去除首尾空格后判断是否为空
		if trimmed == "" {
			continue // 跳过空字符串（包括全是空格的）
		}
		if _, exists := seen[trimmed]; !exists {
			seen[trimmed] = struct{}{}
			result = append(result, trimmed)
		}
	}

	return result
}
