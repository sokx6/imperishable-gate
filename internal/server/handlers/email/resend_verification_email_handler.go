package email

import (
	emailService "imperishable-gate/internal/server/service/email"
	"imperishable-gate/internal/server/service/common"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ResendVerificationEmailHandler 重新发送验证邮件
func ResendVerificationEmailHandler(c echo.Context) error {
	var req request.ResendVerificationRequest
	if err := c.Bind(&req); err != nil {
		return response.InvalidRequestResponse
	}

	if req.Email == "" {
		return response.EmailCannotBeEmptyResponse
	}

	// 使用 ResendVerificationEmail 服务，包含2分钟冷却时间检查
	err := emailService.ResendVerificationEmail(req.Email)
	if err != nil {
		switch err {
		case common.ErrUserNotFound:
			return response.EmailNotRegisteredResponse
		case common.ErrEmailAlreadyVerified:
			return response.EmailAlreadyVerifiedResponse
		case common.ErrResendTooSoon:
			return response.ResendTooSoonResponse
		default:
			return response.ResendVerificationEmailFailedResponse
		}
	}

	return c.JSON(http.StatusOK, response.VerificationEmailResentSuccessResponse)
}
