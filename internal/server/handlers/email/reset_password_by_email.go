package email

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	emailService "imperishable-gate/internal/server/service/email"
	"imperishable-gate/internal/server/service/common"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SendResetPasswordEmailByEmailHandler(c echo.Context) error {
	var req request.SendResetPasswordEmailByRequest
	if err := c.Bind(&req); err != nil {
		return response.InvalidRequestResponse
	}

	if req.Email == "" {
		return response.EmailCannotBeEmptyResponse
	}
	var user model.User
	if database.DB.Where("email = ?", req.Email).First(&user).Error != nil {
		return response.EmailNotRegisteredResponse
	}

	// 调用服务发送重置密码邮件
	err := emailService.SendVerificationEmail(user.ID, req.Email)
	if err != nil {
		switch err {
		case common.ErrUserNotFound:
			return response.EmailNotRegisteredResponse
		case common.ErrDatabase:
			return response.DatabaseErrorResponse
		default:
			return response.SendResetPasswordEmailFailedResponse
		}
	}

	return c.JSON(http.StatusOK, response.ResetPasswordEmailSentSuccessResponse)
}
