package open

import (
	"fmt"
	"imperishable-gate/internal/client/service/list"

	"github.com/skratchdot/open-golang/open"
)

func HandleOpenByName(addr string, accessToken string, name string) error {
	resp, err := list.ListByName(addr, accessToken, name)
	if err != nil {
		return err
	}

	if len(resp.Links) == 0 {
		fmt.Println("No links found for name:", name)
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
