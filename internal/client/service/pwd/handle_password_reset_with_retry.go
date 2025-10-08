package pwd

import (
	"fmt"
	"imperishable-gate/internal/client/utils"
)

// handlePasswordResetWithRetry 处理验证码输入、密码重置的重试逻辑
// identifier 可以是 email 或 username
func handlePasswordResetWithRetry(
	addr, email, username string,
	resetFunc func(string, string, string, string) error,
	resendFunc func(string, string) error,
) error {
	identifier := email
	if username != "" {
		identifier = username
	}

	// 使用通用的验证码处理逻辑，并在验证成功后执行密码重置
	return utils.HandleVerificationWithRetry(utils.VerificationConfig{
		MaxAttempts: 5,
		ResendFunc: func() error {
			return resendFunc(addr, identifier)
		},
		ValidateFunc: func(code string) error {
			// 读取并确认新密码
			newPassword, err := utils.ReadPasswordWithConfirm(
				"Please enter your new password (min 6 chars): ",
				"Please confirm your new password: ",
				utils.MinPasswordLength,
			)
			if err != nil {
				return err
			}

			// 执行重置密码
			fmt.Println("\nResetting your password...")
			return resetFunc(addr, identifier, code, newPassword)
		},
		SuccessMessage: "Password reset successful! You can now login with your new password.",
		FailureMessage: "Password reset failed",
	})
}
