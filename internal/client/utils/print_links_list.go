package utils

import (
	"fmt"

	"imperishable-gate/internal/types/data"
)

// PrintLinksList 打印链接列表的辅助函数
// concise: 为 true 时只显示链接和名字
func PrintLinksList(links []data.Link, concise bool) {
	if len(links) == 0 {
		fmt.Println("No links found.")
		return
	}

	fmt.Printf("Found %d link(s):\n\n", len(links))
	for i, link := range links {
		fmt.Printf("[%d] URL: %s\n", i+1, link.Url)

		if len(link.Names) > 0 {
			fmt.Printf("    Names: %v\n", link.Names)
		} else {
			fmt.Printf("    Names: (none)\n")
		}

		// 简洁模式下只显示链接和名字
		if concise {
			fmt.Println()
			continue
		}

		if len(link.Tags) > 0 {
			fmt.Printf("    Tags: %v\n", link.Tags)
		} else {
			fmt.Printf("    Tags: (none)\n")
		}

		if link.Remark != "" {
			fmt.Printf("    Remark: %s\n", link.Remark)
		} else {
			fmt.Printf("    Remark: (none)\n")
		}

		// 显示元数据（如果存在）
		fmt.Printf("    Watching: %t\n", link.Watching)
		if link.Title != "" {
			fmt.Printf("    Title: %s\n", link.Title)
		}
		if link.Description != "" {
			fmt.Printf("    Description: %s\n", link.Description)
		}
		if link.Keywords != "" {
			fmt.Printf("    Keywords: %s\n", link.Keywords)
		}
		if link.StatusCode != 0 {
			fmt.Printf("    Status Code: %d\n", link.StatusCode)
		}

		fmt.Println()
	}
}
