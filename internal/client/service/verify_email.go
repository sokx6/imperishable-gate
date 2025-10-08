package service

import (
	"fmt"
	"imperishable-gate/internal/client/utils"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
)

// VerifyEmail 向服务器发送邮箱验证请求
func VerifyEmail(addr, email, code string) error {
	// 创建 API 客户端（不需要 token，因为还未登录）
	client := utils.NewAPIClient(addr, "")

	// 构建验证请求体
	reqBody := request.EmailVerificationRequest{
		Email: email,
		Code:  code,
	}

	// 使用 APIClient 发送请求
	var verifyResp response.Response
	err := client.DoRequest("POST", "/api/v1/verify-email", reqBody, &verifyResp)
	if err != nil {
		return fmt.Errorf("failed to verify email: %w", err)
	}

	return nil
}

// ResendVerificationEmail 重新发送验证邮件
func ResendVerificationEmail(addr, email string) error {
	// 创建 API 客户端
	client := utils.NewAPIClient(addr, "")

	// 构建重发请求体
	reqBody := request.ResendVerificationRequest{
		Email: email,
	}

	// 发送请求
	var resp response.Response
	err := client.DoRequest("POST", "/api/v1/resend-verification", reqBody, &resp)
	if err != nil {
		return fmt.Errorf("failed to resend verification email: %w", err)
	}

	return nil
}
