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

// VerifyEmailByUsernameAndResetPasswordHandler 处理邮箱验证和重置密码请求
func VerifyEmailByUsernameAndResetPasswordHandler(c echo.Context) error {
	var req request.ResetPasswordByUsernameRequest
	if err := c.Bind(&req); err != nil {
		logger.Warning("Invalid reset password request: %v", err)
		return response.InvalidRequestResponse
	}

	if req.Username == "" || req.Code == "" || req.NewPassword == "" {
		logger.Warning("Username, code or new password cannot be empty")
		return response.UsernameOrCodeCannotBeEmptyResponse
	}

	// 调用验证服务
	err := emailService.VerifyEmailByUsernameAndResetPassword(req.Username, req.Code, req.NewPassword)
	if err != nil {
		switch err {
		case common.ErrInvalidVerificationCode:
			logger.Warning("Invalid verification code for username: %s", req.Username)
			return response.InvalidVerificationCodeResponse
		case common.ErrVerificationExpired:
			logger.Warning("Verification expired for username: %s", req.Username)
			return response.VerificationExpiredResponse
		case common.ErrTooManyAttempts:
			logger.Warning("Too many attempts for username: %s", req.Username)
			return response.TooManyAttemptsResponse
		case common.ErrDatabase:
			logger.Error("Database error: %v", err)
			return response.DatabaseErrorResponse
		default:
			logger.Error("Email verification failed: %v", err)
			return response.VerificationFailedResponse
		}
	}

	logger.Success("Email verified and password reset successfully for username: %s", req.Username)
	return c.JSON(http.StatusOK, response.PasswordResetSuccessResponse)
}
