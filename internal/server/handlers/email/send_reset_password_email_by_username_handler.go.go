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

func SendResetPasswordEmailByUsernameHandler(c echo.Context) error {
	var req request.SendResetPasswordEmailByUsernameRequest
	if err := c.Bind(&req); err != nil {
		logger.Warning("Invalid send reset password email request: %v", err)
		return response.InvalidRequestResponse
	}

	if req.Username == "" {
		logger.Warning("Username cannot be empty")
		return response.UsernameCannotBeEmptyResponse
	}
	var user model.User
	if database.DB.Where("username = ?", req.Username).First(&user).Error != nil {
		logger.Warning("Username not registered: %s", req.Username)
		return response.UsernameNotRegisteredResponse
	}
	// 调用服务发送重置密码邮件
	err := emailService.SendVerificationEmail(user.ID, user.Email)
	if err != nil {
		switch err {
		case common.ErrUserNotFound:
			logger.Warning("User not found: %s", req.Username)
			return response.UsernameNotRegisteredResponse
		case common.ErrDatabase:
			logger.Error("Database error: %v", err)
			return response.DatabaseErrorResponse
		default:
			logger.Error("Failed to send reset password email: %v", err)
			return response.SendResetPasswordEmailFailedResponse
		}
	}

	logger.Success("Reset password email sent successfully to %s", user.Email)
	return c.JSON(http.StatusOK, response.ResetPasswordEmailSentSuccessResponse)
}
