package cmd

import (
	"fmt"

	"imperishable-gate/internal/client/service/watch"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch or unwatch links to monitor changes",
	Long: `Watch command supports multiple operations:
  - Watch a link by URL: --link <url> --watch
  - Unwatch a link by URL: --link <url> --unwatch
  - Watch a link by name: --name <name> --watch
  - Unwatch a link by name: --name <name> --unwatch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Executing watch command...")

		// 获取命令行参数
		link, _ := cmd.Flags().GetString("link")
		name, _ := cmd.Flags().GetString("name")
		isWatch, _ := cmd.Flags().GetBool("watch")
		isUnwatch, _ := cmd.Flags().GetBool("unwatch")

		// 参数验证：必须提供 link 或 name
		if link == "" && name == "" {
			fmt.Println("Error: Either --link or --name must be provided.")
			return fmt.Errorf("either --link or --name must be provided")
		}

		// 参数验证：不能同时提供 link 和 name
		if link != "" && name != "" {
			fmt.Println("Error: Cannot specify both --link and --name.")
			return fmt.Errorf("cannot specify both --link and --name")
		}

		// 参数验证：必须提供 watch 或 unwatch，且只能二选一
		if isWatch == isUnwatch {
			fmt.Println("Error: Must specify either --watch or --unwatch (but not both).")
			return fmt.Errorf("must specify either --watch or --unwatch")
		}

		// 场景1: 通过链接 URL 监控
		if link != "" {
			return watch.HandleWatchByLink(link, isWatch, addr, accessToken)
		}

		// 场景2: 通过名称监控
		if name != "" {
			return watch.HandleWatchByName(name, isWatch, addr, accessToken)
		}

		return fmt.Errorf("invalid parameter combination")
	},
}

// 初始化命令行参数
func init() {
	watchCmd.Flags().StringP("link", "l", "", "Link URL to watch/unwatch")
	watchCmd.Flags().StringP("name", "n", "", "Link name to watch/unwatch")
	watchCmd.Flags().BoolP("watch", "w", false, "Enable watching for the link")
	watchCmd.Flags().BoolP("unwatch", "u", false, "Disable watching for the link")
	rootCmd.AddCommand(watchCmd)
}
