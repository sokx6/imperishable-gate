package handlers

import (
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/service"
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
	err := service.SendVerificationEmail(user.ID, req.Email)
	if err != nil {
		switch err {
		case service.ErrUserNotFound:
			return response.EmailNotRegisteredResponse
		case service.ErrDatabase:
			return response.DatabaseErrorResponse
		default:
			return response.SendResetPasswordEmailFailedResponse
		}
	}

	return c.JSON(http.StatusOK, response.ResetPasswordEmailSentSuccessResponse)
}
