package email

import (
	emailService "imperishable-gate/internal/server/service/email"
	"imperishable-gate/internal/server/service/common"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// VerifyEmailByUsernameAndResetPasswordHandler 处理邮箱验证和重置密码请求
func VerifyEmailByUsernameAndResetPasswordHandler(c echo.Context) error {
	var req request.ResetPasswordByUsernameRequest
	if err := c.Bind(&req); err != nil {
		return response.InvalidRequestResponse
	}

	if req.Username == "" || req.Code == "" || req.NewPassword == "" {
		return response.UsernameOrCodeCannotBeEmptyResponse
	}

	// 调用验证服务
	err := emailService.VerifyEmailByUsernameAndResetPassword(req.Username, req.Code, req.NewPassword)
	if err != nil {
		switch err {
		case common.ErrInvalidVerificationCode:
			return response.InvalidVerificationCodeResponse
		case common.ErrVerificationExpired:
			return response.VerificationExpiredResponse
		case common.ErrTooManyAttempts:
			return response.TooManyAttemptsResponse
		case common.ErrDatabase:
			return response.DatabaseErrorResponse
		default:
			return response.VerificationFailedResponse
		}
	}
	return c.JSON(http.StatusOK, response.PasswordResetSuccessResponse)
}
