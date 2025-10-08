package cmd

import (
	"fmt"
	"os"

	"imperishable-gate/internal/client/core"

	"github.com/spf13/cobra"
)

var accessToken string = ""
var addr string
var inInteractiveMode bool = false

var rootCmd = &cobra.Command{
	Use:   "gate",
	Short: "gate is a CLI link management tool",
	Long:  "gate is a CLI link management tool . . .",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// 获取服务器地址
		addr = core.LoadServerAddr(cmd)

		// 对于其他命令，预先验证 token
		return core.EnsureAuthentication(cmd, addr, &accessToken)
	},
	// Run 执行根命令
	Run: func(cmd *cobra.Command, args []string) {
		// 如果已经在交互模式中，什么都不做（避免干扰子命令）
		if inInteractiveMode {
			return
		}

		// 如果没有提供任何参数，进入交互模式
		if len(args) == 0 && !cmd.Flags().Changed("help") && !cmd.Flags().Changed("version") {
			handler := &core.InteractiveModeHandler{
				RootCmd:           cmd,
				InInteractiveMode: &inInteractiveMode,
			}
			handler.Run()
			return
		}
		fmt.Println("Welcome to gate CLI tool. Use -h for help.")
	},
	Version: "1.0.0",
}

func init() {
	// 定义全局持久标志addr（服务器地址）
	rootCmd.PersistentFlags().StringP("addr", "a", "", "Server address (host:port)")
}

// Execute 执行根命令
func Execute() {
	// 为所有命令设置静默模式，我们自己处理错误和用法的显示顺序
	silencer := &core.CommandSilencer{
		InInteractiveMode: &inInteractiveMode,
	}
	silencer.Apply(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		// 不打印错误，因为已经在命令执行时处理了
		os.Exit(1)
	}
}
