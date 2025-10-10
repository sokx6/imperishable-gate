package cmd

import (
	"fmt"

	"imperishable-gate/internal/client/service/list"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List links from the server database",
	Long: `List command supports multiple operations:
  - List all links: (no parameters)
  - List links by tag: --tag <tag>
  - List link by name: --name <name>
  - List with pagination: --page <page> --page-size <size>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Executing list command...")

		// 获取命令行参数
		tag, _ := cmd.Flags().GetString("tag")
		name, _ := cmd.Flags().GetString("name")
		page, _ := cmd.Flags().GetInt("page")
		pageSize, _ := cmd.Flags().GetInt("page-size")
		concise, _ := cmd.Flags().GetBool("concise")

		// 场景1: 通过标签查询
		if tag != "" {
			return list.HandleListByTag(tag, page, pageSize, addr, accessToken, concise)
		}

		// 场景2: 通过名称查询
		if name != "" {
			return list.HandleListByName(name, page, pageSize, addr, accessToken, concise)
		}

		// 场景3: 列出所有链接（默认）
		return list.HandleListAllLinks(page, pageSize, addr, accessToken, concise)
	},
}

// 初始化命令行参数
func init() {
	listCmd.Flags().StringP("tag", "t", "", "Filter links by tag")
	listCmd.Flags().StringP("name", "n", "", "Filter link by name")
	listCmd.Flags().IntP("page", "p", 1, "Page number for pagination")
	listCmd.Flags().IntP("page-size", "s", 20, "Number of items per page")
	listCmd.Flags().BoolP("concise", "c", false, "Show only URL and names (concise output)")
	rootCmd.AddCommand(listCmd)
}
