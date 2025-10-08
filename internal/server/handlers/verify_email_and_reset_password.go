package handlers

import (
	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// VerifyEmailAndResetPasswordHandler 处理邮箱验证和重置密码请求
func VerifyEmailAndResetPasswordHandler(c echo.Context) error {
	var req request.ResetPasswordByEmailRequest
	if err := c.Bind(&req); err != nil {
		return response.InvalidRequestResponse
	}

	if req.Email == "" || req.Code == "" || req.NewPassword == "" {
		return response.EmailOrCodeCannotBeEmptyResponse
	}

	// 调用验证服务
	err := service.VerifyEmailAndResetPassword(req.Email, req.Code, req.NewPassword)
	if err != nil {
		switch err {
		case service.ErrInvalidVerificationCode:
			return response.InvalidVerificationCodeResponse
		case service.ErrVerificationExpired:
			return response.VerificationExpiredResponse
		case service.ErrTooManyAttempts:
			return response.TooManyAttemptsResponse
		case service.ErrDatabase:
			return response.DatabaseErrorResponse
		default:
			return response.VerificationFailedResponse
		}
	}
	return c.JSON(http.StatusOK, response.PasswordResetSuccessResponse)
}
