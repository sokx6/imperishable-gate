package open

import (
	"fmt"
	"imperishable-gate/internal/client/service/list"

	"github.com/skratchdot/open-golang/open"
)

func HandleOpenByTag(addr string, accessToken string, tag string, page int, pageSize int) error {
	resp, err := list.ListByTag(addr, accessToken, tag, page, pageSize)
	if err != nil {
		return err
	}

	if len(resp.Links) == 0 {
		fmt.Printf("No links found for tag: %s (page: %d, page-size: %d)\n", tag, page, pageSize)
		return nil
	}

	successCount := 0
	var errors []string

	for _, link := range resp.Links {
		if err := open.Run(link.Url); err != nil {
			errors = append(errors, fmt.Sprintf("Failed to open %s: %v", link.Url, err))
		} else {
			successCount++
		}
	}

	fmt.Printf("Successfully opened %d/%d links\n", successCount, len(resp.Links))

	if len(errors) > 0 {
		fmt.Println("\nErrors occurred:")
		for _, errMsg := range errors {
			fmt.Println("  -", errMsg)
		}
	}

	return nil
}
