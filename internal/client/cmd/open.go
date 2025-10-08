package cmd

import (
	"fmt"
	"imperishable-gate/internal/client/service/open"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open links by name or tag",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		tag, _ := cmd.Flags().GetString("tag")
		page, _ := cmd.Flags().GetInt("page")
		pageSize, _ := cmd.Flags().GetInt("page-size")
		if name != "" {
			return open.HandleOpenByName(addr, accessToken, name)
		} else if tag != "" {
			return open.HandleOpenByTag(addr, accessToken, tag, page, pageSize)
		} else {
			return fmt.Errorf("either --name or --tag must be provided")
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	openCmd.Flags().StringP("name", "n", "", "Name of the link to open")
	openCmd.Flags().StringP("tag", "t", "", "Tag of the links to open")
	openCmd.Flags().IntP("page", "p", 1, "Page number for pagination (used with --tag)")
	openCmd.Flags().IntP("page-size", "s", 10, "Number of items per page (used with --tag)")
}
