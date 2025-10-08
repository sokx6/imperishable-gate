package cmd

import (
	"fmt"
	"imperishable-gate/internal/client/service/search"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search links in the server database",
	Long: `Search command supports multiple operations:
  - Search links by keyword: --keyword <keyword>
  - Search with pagination: --page <page> --page-size <size>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Executing search command...")

		// 获取命令行参数
		keyword, _ := cmd.Flags().GetString("keyword")
		page, _ := cmd.Flags().GetInt("page")
		pageSize, _ := cmd.Flags().GetInt("page-size")

		// 通过关键字搜索
		if keyword != "" {
			return search.HandleSearchByKeyword(keyword, page, pageSize, addr, accessToken)
		}

		return fmt.Errorf("no search criteria provided")
	},
}

func init() {
	searchCmd.Flags().StringP("keyword", "k", "", "Search links by keyword")
	searchCmd.Flags().IntP("page", "p", 1, "Page number for pagination")
	searchCmd.Flags().IntP("page-size", "s", 20, "Number of items per page")
	rootCmd.AddCommand(searchCmd)
}
