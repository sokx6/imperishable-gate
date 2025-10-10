package list

import (
	"fmt"
	"imperishable-gate/internal/client/utils"
)

// HandleListByTag 通过标签查询链接
func HandleListByTag(tag string, page int, pageSize int, addr string, accessToken string, concise bool) error {
	fmt.Printf("Listing links with tag: %s\n", tag)

	result, err := ListByTag(addr, accessToken, tag, page, pageSize)
	if err != nil {
		return fmt.Errorf("failed to list links by tag: %w", err)
	}

	utils.PrintLinksList(result.Links, concise)
	return nil
}

// HandleListByName 通过名称查询链接
func HandleListByName(name string, page int, pageSize int, addr string, accessToken string, concise bool) error {
	fmt.Printf("Listing link with name: %s\n", name)

	result, err := ListByName(addr, accessToken, name)
	if err != nil {
		return fmt.Errorf("failed to list link by name: %w", err)
	}

	// Response 返回的是单个链接

	utils.PrintLinksList(result.Links, concise)

	return nil
}

// HandleListAllLinks 列出所有链接
func HandleListAllLinks(page int, pageSize int, addr string, accessToken string, concise bool) error {
	fmt.Println("Listing all links...")

	result, err := ListLinks(addr, accessToken, page, pageSize)
	if err != nil {
		return fmt.Errorf("failed to list links: %w", err)
	}

	utils.PrintLinksList(result.Links, concise)
	return nil
}
