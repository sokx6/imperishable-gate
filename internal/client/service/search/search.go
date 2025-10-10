package search

import (
	"fmt"
	"imperishable-gate/internal/client/utils"
)

func HandleSearchByKeyword(keyword string, page int, pageSize int, addr string, accessToken string, concise bool) error {
	fmt.Printf("Searching links with keyword: %s\n", keyword)

	result, err := SearchLinks(addr, accessToken, keyword, page, pageSize)
	if err != nil {
		return fmt.Errorf("failed to search links by keyword: %w", err)
	}

	utils.PrintLinksList(result.Links, concise)
	return nil
}
