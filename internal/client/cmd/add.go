package cmd

import (
	"fmt"

	"imperishable-gate/internal/client/service/add"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new link to the server database",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Executing add command...")
		// 获取命令行参数
		link, _ := cmd.Flags().GetString("link")
		tags, _ := cmd.Flags().GetStringArray("tag")
		name, _ := cmd.Flags().GetString("name")
		remark, _ := cmd.Flags().GetString("remark")
		if link == "" {
			if name == "" {
				fmt.Println("Either link or name must be provided.")
				return fmt.Errorf("either link or name must be provided")
			}
			if len(tags) > 0 {
				fmt.Println("Adding tags to link with name:", name)
				return add.AddTagsByName(name, tags, addr, accessToken)
			}
			if remark != "" {
				fmt.Println("Adding remark to link with name:", name)
				return add.AddRemarkByName(name, remark, addr, accessToken)
			}
			fmt.Println("No tags or remark to add for the given name.")
			return fmt.Errorf("no tags or remark to add for the given name")
		}

		if name == "" {
			if len(tags) > 0 {
				fmt.Println("Adding tags to link:", link)
				return add.AddTagsByLink(link, tags, addr, accessToken)
			}
			if remark != "" {
				fmt.Println("Adding remark to link:", link)
				return add.AddRemarkByLink(link, remark, addr, accessToken)
			}
			fmt.Println("No tags or remark to add for the given link.")
			return fmt.Errorf("no tags or remark to add for the given link")
		}

		fmt.Println("Adding link:", link)
		fmt.Println("server:", addr)
		return add.AddLink(link, addr, accessToken)
	},
}

// 初始化命令行参数
func init() {
	// 为 add 命令添加参数link，用来指定要添加的链接
	addCmd.Flags().StringP("link", "l", "", "link to add")
	addCmd.Flags().StringArrayP("tag", "t", []string{}, "tags for the link")
	addCmd.Flags().StringP("name", "n", "", "name for the link")
	addCmd.Flags().StringP("remark", "r", "", "remark for the link")
	rootCmd.AddCommand(addCmd)
}
