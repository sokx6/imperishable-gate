package email

import (
	"imperishable-gate/internal/server/service/common"
	emailService "imperishable-gate/internal/server/service/email"
	"imperishable-gate/internal/server/utils/logger"

	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// VerifyEmailAndResetPasswordHandler 处理邮箱验证和重置密码请求
func VerifyEmailAndResetPasswordHandler(c echo.Context) error {
	var req request.ResetPasswordByEmailRequest
	if err := c.Bind(&req); err != nil {
		logger.Warning("Invalid reset password request: %v", err)
		return response.InvalidRequestResponse
	}

	if req.Email == "" || req.Code == "" || req.NewPassword == "" {
		logger.Warning("Email, code or new password cannot be empty")
		return response.EmailOrCodeCannotBeEmptyResponse
	}

	// 调用验证服务
	err := emailService.VerifyEmailAndResetPassword(req.Email, req.Code, req.NewPassword)
	if err != nil {
		switch err {
		case common.ErrInvalidVerificationCode:
			logger.Warning("Invalid verification code for email: %s", req.Email)
			return response.InvalidVerificationCodeResponse
		case common.ErrVerificationExpired:
			logger.Warning("Verification expired for email: %s", req.Email)
			return response.VerificationExpiredResponse
		case common.ErrTooManyAttempts:
			logger.Warning("Too many attempts for email: %s", req.Email)
			return response.TooManyAttemptsResponse
		case common.ErrDatabase:
			logger.Error("Database error: %v", err)
			return response.DatabaseErrorResponse
		default:
			logger.Error("Email verification failed: %v", err)
			return response.VerificationFailedResponse
		}
	}

	logger.Success("Email verified and password reset successfully: %s", req.Email)
	return c.JSON(http.StatusOK, response.PasswordResetSuccessResponse)
}
