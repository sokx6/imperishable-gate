package cmd

import (
	"fmt"

	"imperishable-gate/internal/client/service/add"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new link to the server database",
	Long: `Add command supports multiple operations:
  - Add a new link: --link <url>
  - Add names to a link: --link <url> --name <name1> [--name <name2> ...]
  - Add tags by link: --link <url> --tag <tag1> [--tag <tag2> ...]
  - Add tags by name: --name <name> --tag <tag1> [--tag <tag2> ...]
  - Add remark by link: --link <url> --remark <remark>
  - Add remark by name: --name <name> --remark <remark>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Executing add command...")

		// 获取命令行参数
		link, _ := cmd.Flags().GetString("link")
		tags, _ := cmd.Flags().GetStringSlice("tag")
		names, _ := cmd.Flags().GetStringSlice("name")
		remark, _ := cmd.Flags().GetString("remark")

		// 参数验证：至少需要 link 或 name
		if link == "" && len(names) == 0 {
			fmt.Println("Error: Either --link or --name must be provided.")
			return fmt.Errorf("either --link or --name must be provided")
		}

		// 场景1: 只提供了 name (没有 link)
		if link == "" && len(names) > 0 {
			return add.HandleAddByName(names, tags, remark, addr, accessToken)
		}

		// 场景2: 只提供了 link (没有 name)
		if link != "" && len(names) == 0 {
			return add.HandleAddByLink(link, tags, remark, addr, accessToken)
		}

		// 场景3: 同时提供了 link 和 name - 添加新链接及其名称
		if link != "" && len(names) > 0 {
			return add.HandleAddLinkWithNames(link, names, tags, remark, addr, accessToken)
		}

		return fmt.Errorf("invalid parameter combination")
	},
}

// 初始化命令行参数
func init() {
	// 为 add 命令添加参数
	addCmd.Flags().StringP("link", "l", "", "Link URL to add or modify")
	addCmd.Flags().StringSliceP("tag", "t", []string{}, "Tags for the link (can specify multiple)")
	addCmd.Flags().StringSliceP("name", "n", []string{}, "Name(s) for the link (can specify multiple)")
	addCmd.Flags().StringP("remark", "r", "", "Remark for the link")
	rootCmd.AddCommand(addCmd)
}
