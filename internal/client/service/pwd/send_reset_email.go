package pwd

import (
	"fmt"
)

// sendResetEmailNotification 发送重置密码邮件并显示通知
func sendResetEmailNotification(addr, identifier string, sendFunc func(string, string) error) error {
	fmt.Println("\nSending password reset email...")
	if err := sendFunc(addr, identifier); err != nil {
		return fmt.Errorf("failed to send reset email: %w", err)
	}
	fmt.Println("Password reset email sent! Please check your inbox.")
	fmt.Println("The code is valid for 15 minutes.")
	return nil
}
