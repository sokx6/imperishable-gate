package pwd

import (
	"fmt"
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

// HandleResetPasswordByEmail 处理通过邮箱重置密码的完整流程
func HandleResetPasswordByEmail(addr string) error {
	// 1. 读取并验证邮箱
	email, err := utils.ReadEmail("")
	if err != nil {
		return err
	}

	// 2. 请求发送重置密码邮件
	if err := sendResetEmailNotification(addr, email, SendResetPasswordEmail); err != nil {
		return err
	}

	// 3. 处理验证和重置密码
	return handlePasswordResetWithRetry(addr, email, "", ResetPasswordByEmail, SendResetPasswordEmail)
}

// HandleResetPasswordByUsername 处理通过用户名重置密码的完整流程
func HandleResetPasswordByUsername(addr string) error {
	// 1. 读取用户名
	username, err := utils.ReadUsername("")
	if err != nil {
		return err
	}

	// 2. 请求发送重置密码邮件
	if err := sendResetEmailNotification(addr, username, SendResetPasswordEmailByUsername); err != nil {
		return err
	}

	// 3. 处理验证和重置密码
	return handlePasswordResetWithRetry(addr, "", username, ResetPasswordByUsername, SendResetPasswordEmailByUsername)
}

// SendResetPasswordEmail 请求发送重置密码邮件（通过邮箱）
func SendResetPasswordEmail(addr, email string) error {
	client := utils.NewAPIClient(addr, "")
	reqBody := request.SendResetPasswordEmailByRequest{Email: email}
	var resp response.Response
	return client.DoRequest("PATCH", "/api/v1/email/password/request", reqBody, &resp)
}

// SendResetPasswordEmailByUsername 请求发送重置密码邮件（通过用户名）
func SendResetPasswordEmailByUsername(addr, username string) error {
	client := utils.NewAPIClient(addr, "")
	reqBody := request.SendResetPasswordEmailByUsernameRequest{Username: username}
	var resp response.Response
	return client.DoRequest("PATCH", "/api/v1/username/password/request", reqBody, &resp)
}

// ResetPasswordByEmail 通过邮箱重置密码
func ResetPasswordByEmail(addr, email, code, newPassword string) error {
	client := utils.NewAPIClient(addr, "")
	reqBody := request.ResetPasswordByEmailRequest{
		Email:       email,
		Code:        code,
		NewPassword: newPassword,
	}
	var resp response.Response
	if err := client.DoRequest("PATCH", "/api/v1/email/password", reqBody, &resp); err != nil {
		return fmt.Errorf("reset password failed: %w", err)
	}
	return nil
}

// ResetPasswordByUsername 通过用户名重置密码
func ResetPasswordByUsername(addr, username, code, newPassword string) error {
	client := utils.NewAPIClient(addr, "")
	reqBody := request.ResetPasswordByUsernameRequest{
		Username:    username,
		Code:        code,
		NewPassword: newPassword,
	}
	var resp response.Response
	if err := client.DoRequest("PATCH", "/api/v1/username/password", reqBody, &resp); err != nil {
		return fmt.Errorf("reset password failed: %w", err)
	}
	return nil
}
