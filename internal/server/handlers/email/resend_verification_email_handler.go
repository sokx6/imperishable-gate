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

// ResendVerificationEmailHandler 重新发送验证邮件
func ResendVerificationEmailHandler(c echo.Context) error {
	var req request.ResendVerificationRequest
	if err := c.Bind(&req); err != nil {
		logger.Warning("Invalid resend verification email request: %v", err)
		return response.InvalidRequestResponse
	}

	if req.Email == "" {
		logger.Warning("Email cannot be empty")
		return response.EmailCannotBeEmptyResponse
	}

	// 使用 ResendVerificationEmail 服务，包含2分钟冷却时间检查
	err := emailService.ResendVerificationEmail(req.Email)
	if err != nil {
		switch err {
		case common.ErrUserNotFound:
			logger.Warning("User not found: %s", req.Email)
			return response.EmailNotRegisteredResponse
		case common.ErrEmailAlreadyVerified:
			logger.Warning("Email already verified: %s", req.Email)
			return response.EmailAlreadyVerifiedResponse
		case common.ErrResendTooSoon:
			logger.Warning("Resend too soon: %s", req.Email)
			return response.ResendTooSoonResponse
		default:
			logger.Error("Failed to resend verification email: %v", err)
			return response.ResendVerificationEmailFailedResponse
		}
	}

	logger.Success("Verification email resent successfully to %s", req.Email)
	return c.JSON(http.StatusOK, response.VerificationEmailResentSuccessResponse)
}
