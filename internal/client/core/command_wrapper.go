package core

import "github.com/spf13/cobra"

// CommandSilencer 命令静默处理器
type CommandSilencer struct {
	InInteractiveMode *bool
}

// Apply 递归设置所有命令的静默模式和错误处理
func (s *CommandSilencer) Apply(cmd *cobra.Command) {
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true

	// 为每个命令设置自定义的 RunE 包装器（仅用于非交互模式）
	if cmd.RunE != nil {
		s.wrapRunE(cmd)
	}

	// 包装 PersistentPreRunE 以在非交互模式下显示错误
	if cmd.PersistentPreRunE != nil {
		s.wrapPersistentPreRunE(cmd)
	}

	// 递归处理所有子命令
	for _, subCmd := range cmd.Commands() {
		s.Apply(subCmd)
	}
}

// wrapRunE 包装 RunE 函数以在非交互模式下显示错误
func (s *CommandSilencer) wrapRunE(cmd *cobra.Command) {
	originalRunE := cmd.RunE
	cmd.RunE = func(c *cobra.Command, args []string) error {
		err := originalRunE(c, args)
		if err != nil && !*s.InInteractiveMode {
			// 非交互模式：先打印 Usage，再打印错误
			c.Println(c.UsageString())
			c.PrintErrln("Error:", err.Error())
		}
		return err
	}
}

// wrapPersistentPreRunE 包装 PersistentPreRunE 函数以在非交互模式下显示错误
func (s *CommandSilencer) wrapPersistentPreRunE(cmd *cobra.Command) {
	originalPersistentPreRunE := cmd.PersistentPreRunE
	cmd.PersistentPreRunE = func(c *cobra.Command, args []string) error {
		err := originalPersistentPreRunE(c, args)
		if err != nil && !*s.InInteractiveMode {
			// 非交互模式：先打印 Usage，再打印错误
			c.Println(c.UsageString())
			c.PrintErrln("Error:", err.Error())
		}
		return err
	}
}
