package cmd

import (
	"imperishable-gate/internal/client/service/pwd"

	"github.com/spf13/cobra"
)

var resetPasswordCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Reset your account password via email or username verification",
	Long: `Reset your account password using email or username verification.
You will receive a verification code via email to complete the password reset.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// addr 已经在 PersistentPreRunE 中通过 core.LoadServerAddr 设置了
		// 不需要再次处理，直接使用全局变量 addr

		// 检查是否使用用户名模式
		useUsername, _ := cmd.Flags().GetBool("username")

		if useUsername {
			return pwd.HandleResetPasswordByUsername(addr)
		}
		return pwd.HandleResetPasswordByEmail(addr)
	},
}

func init() {
	rootCmd.AddCommand(resetPasswordCmd)
	resetPasswordCmd.Flags().BoolP("username", "u", false, "Reset password using username instead of email")
}
