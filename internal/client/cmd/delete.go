package cmd

import (
	"fmt"

	"imperishable-gate/internal/client/service/delete"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete links, names, or tags from the server database",
	Long: `Delete command supports multiple operations:
  - Delete links: --link <url1> [--link <url2> ...]
  - Delete a link by name: --name <name>
  - Delete names from a link: --link <url> --name <name1> [--name <name2> ...]
  - Delete tags from a link: --link <url> --tag <tag1> [--tag <tag2> ...]
  - Delete tags by name: --name <name> --tag <tag1> [--tag <tag2> ...]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Executing delete command...")

		// 获取命令行参数
		links, _ := cmd.Flags().GetStringSlice("link")
		names, _ := cmd.Flags().GetStringSlice("name")
		tags, _ := cmd.Flags().GetStringSlice("tag")

		// 参数验证：至少需要 link 或 name
		if len(links) == 0 && len(names) == 0 {
			return fmt.Errorf("either --link or --name must be provided")
		}

		// 场景1: 只提供了 name (没有 link)
		if len(links) == 0 && len(names) > 0 {
			return delete.HandleDeleteByName(names, tags, addr, accessToken)
		}

		// 场景2: 只提供了 link (没有 name)
		if len(links) > 0 && len(names) == 0 {
			return delete.HandleDeleteByLink(links, tags, addr, accessToken)
		}

		// 场景3: 同时提供了 link 和 name - 删除链接的名称
		if len(links) > 0 && len(names) > 0 {
			return delete.HandleDeleteNamesFromLink(links, names, addr, accessToken)
		}

		return fmt.Errorf("invalid parameter combination")
	},
}

// 初始化命令行参数
func init() {
	deleteCmd.Flags().StringSliceP("link", "l", []string{}, "Link(s) to delete or modify (can specify multiple)")
	deleteCmd.Flags().StringSliceP("name", "n", []string{}, "Name(s) to delete (can specify multiple)")
	deleteCmd.Flags().StringSliceP("tag", "t", []string{}, "Tag(s) to delete from a link or name (can specify multiple)")
	rootCmd.AddCommand(deleteCmd)
}
