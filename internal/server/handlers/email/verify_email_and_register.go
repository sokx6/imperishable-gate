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

// VerifyEmailAndRegisterHandler 处理邮箱验证请求（带暴力破解防护）
func VerifyEmailAndRegisterHandler(c echo.Context) error {
	var req request.EmailVerificationRequest
	if err := c.Bind(&req); err != nil {
		logger.Warning("Invalid email verification request: %v", err)
		return response.InvalidRequestResponse
	}

	if req.Email == "" || req.Code == "" {
		logger.Warning("Email or code cannot be empty")
		return response.EmailOrCodeCannotBeEmptyResponse
	}

	// 调用验证服务
	err := emailService.VerifyEmailAndRegister(req.Email, req.Code)
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

	logger.Success("Email verified and user registered successfully: %s", req.Email)
	return c.JSON(http.StatusOK, response.EmailVerifiedSuccessResponse)
}
