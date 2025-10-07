package list

import (
	"fmt"

	"imperishable-gate/internal/types/data"
)

// HandleListByTag 通过标签查询链接
func HandleListByTag(tag string, page int, pageSize int, addr string, accessToken string) error {
	fmt.Printf("Listing links with tag: %s\n", tag)

	result, err := ListByTag(addr, accessToken, tag, page, pageSize)
	if err != nil {
		return fmt.Errorf("failed to list links by tag: %w", err)
	}

	printLinksList(result.Data)
	return nil
}

// HandleListByName 通过名称查询链接
func HandleListByName(name string, page int, pageSize int, addr string, accessToken string) error {
	fmt.Printf("Listing link with name: %s\n", name)

	result, err := ListByName(addr, accessToken, name, page, pageSize)
	if err != nil {
		return fmt.Errorf("failed to list link by name: %w", err)
	}

	// ListByNameResponse 返回的是单个链接
	if result.Data.ID != 0 {
		fmt.Printf("ID: %d\nURL: %s\nTags: %v\nNames: %v\nRemark: %s\n\n",
			result.Data.ID, result.Data.Url, result.Data.Tags, result.Data.Names, result.Data.Remark)
	} else {
		fmt.Println("No link found with the specified name.")
	}

	return nil
}

// HandleListAllLinks 列出所有链接
func HandleListAllLinks(page int, pageSize int, addr string, accessToken string) error {
	fmt.Println("Listing all links...")

	result, err := ListLinks(addr, accessToken, page, pageSize)
	if err != nil {
		return fmt.Errorf("failed to list links: %w", err)
	}

	printLinksList(result.Data)
	return nil
}

// printLinksList 打印链接列表的辅助函数
func printLinksList(links []data.Link) {
	if len(links) == 0 {
		fmt.Println("No links found.")
		return
	}

	fmt.Printf("Found %d link(s):\n\n", len(links))
	for i, link := range links {
		fmt.Printf("[%d] ID: %d\n", i+1, link.ID)
		fmt.Printf("    URL: %s\n", link.Url)

		if len(link.Tags) > 0 {
			fmt.Printf("    Tags: %v\n", link.Tags)
		} else {
			fmt.Printf("    Tags: (none)\n")
		}

		if len(link.Names) > 0 {
			fmt.Printf("    Names: %v\n", link.Names)
		} else {
			fmt.Printf("    Names: (none)\n")
		}

		if link.Remark != "" {
			fmt.Printf("    Remark: %s\n", link.Remark)
		} else {
			fmt.Printf("    Remark: (none)\n")
		}

		// 显示元数据（如果存在）
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
