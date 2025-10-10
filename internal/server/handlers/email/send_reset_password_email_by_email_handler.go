package email

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/service/common"
	emailService "imperishable-gate/internal/server/service/email"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SendResetPasswordEmailByEmailHandler(c echo.Context) error {
	var req request.SendResetPasswordEmailByRequest
	if err := c.Bind(&req); err != nil {
		logger.Warning("Invalid send reset password email request: %v", err)
		return response.InvalidRequestResponse
	}

	if req.Email == "" {
		logger.Warning("Email cannot be empty")
		return response.EmailCannotBeEmptyResponse
	}
	var user model.User
	if database.DB.Where("email = ?", req.Email).First(&user).Error != nil {
		logger.Warning("Email not registered: %s", req.Email)
		return response.EmailNotRegisteredResponse
	}

	// 调用服务发送重置密码邮件
	err := emailService.SendVerificationEmail(user.ID, req.Email)
	if err != nil {
		switch err {
		case common.ErrUserNotFound:
			logger.Warning("User not found: %s", req.Email)
			return response.EmailNotRegisteredResponse
		case common.ErrDatabase:
			logger.Error("Database error: %v", err)
			return response.DatabaseErrorResponse
		default:
			logger.Error("Failed to send reset password email: %v", err)
			return response.SendResetPasswordEmailFailedResponse
		}
	}

	logger.Success("Reset password email sent successfully to %s", req.Email)
	return c.JSON(http.StatusOK, response.ResetPasswordEmailSentSuccessResponse)
}
