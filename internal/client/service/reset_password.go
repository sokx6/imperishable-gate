package service

import (
	"fmt"
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

// HandleResetPassword 处理重置密码完整流程
func HandleResetPassword(addr string) error {
	// 1. 读取邮箱
	email, err := utils.ReadLine("Please enter your email: ")
	if err != nil {
		return err
	}
	if err := utils.ValidateEmail(email); err != nil {
		return err
	}

	// 2. 请求发送重置密码邮件
	fmt.Println("\n📤 Sending password reset email...")
	if err := SendResetPasswordEmail(addr, email); err != nil {
		return fmt.Errorf("failed to send reset email: %w", err)
	}
	fmt.Println("✓ Password reset email sent! Please check your inbox.")
	fmt.Println("💡 The code is valid for 15 minutes.")

	// 3. 验证码+新密码输入，支持重试和resend
	maxAttempts := 5
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("\n[Attempt %d/%d] ", attempt, maxAttempts)
		code, err := utils.ReadLine("Please enter the verification code (or 'resend' to get a new code): ")
		if err != nil {
			return err
		}
		if code == "" {
			return fmt.Errorf("verification code cannot be empty")
		}
		if code == "resend" {
			fmt.Println("\n📧 Resending password reset email...")
			if err := SendResetPasswordEmail(addr, email); err != nil {
				fmt.Println("❌ Failed to resend reset email:", err)
				continue
			}
			fmt.Println("✓ Password reset email has been resent!")
			attempt--
			continue
		}
		if len(code) != 6 {
			fmt.Println("Verification code must be 6 digits.")
			continue
		}
		newPassword, err := utils.ReadLine("Please enter your new password (min 6 chars): ")
		if err != nil {
			return err
		}
		if err := utils.ValidatePassword(newPassword, 6); err != nil {
			fmt.Println(err)
			continue
		}
		if err := utils.ConfirmInput("Please confirm your new password: ", newPassword); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("\n🔄 Resetting your password...")
		err = ResetPassword(addr, email, code, newPassword)
		if err != nil {
			if attempt < maxAttempts {
				fmt.Printf("Reset failed: %v\nYou have %d attempt(s) remaining.\n", err, maxAttempts-attempt)
				fmt.Println("Tip: Enter 'resend' to get a new verification code.")
				continue
			}
			return fmt.Errorf("password reset failed after %d attempts: %w", maxAttempts, err)
		}
		fmt.Println("\n✓ Password reset successful! You can now login with your new password.")
		return nil
	}
	return fmt.Errorf("password reset failed: maximum attempts exceeded")
}

// SendResetPasswordEmail 请求发送重置密码邮件
func SendResetPasswordEmail(addr, email string) error {
	client := utils.NewAPIClient(addr, "")
	reqBody := request.SendResetPasswordEmailByRequest{Email: email}
	var resp response.Response
	// PATCH /api/v1/email/password/request
	return client.DoRequest("PATCH", "/api/v1/email/password/request", reqBody, &resp)
}

// ResetPassword 发送重置密码请求
func ResetPassword(addr, email, code, newPassword string) error {
	client := utils.NewAPIClient(addr, "")
	reqBody := request.ResetPasswordByEmailRequest{
		Email:       email,
		Code:        code,
		NewPassword: newPassword,
	}
	var resp response.Response
	// PATCH /api/v1/email/password
	err := client.DoRequest("PATCH", "/api/v1/email/password", reqBody, &resp)
	if err != nil {
		return fmt.Errorf("reset password failed: %w", err)
	}
	return nil
}
