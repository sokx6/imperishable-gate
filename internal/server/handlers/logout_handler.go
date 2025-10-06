package handlers

import (
	"net/http"

	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"

	"github.com/labstack/echo/v4"
)

func LogoutHandler(c echo.Context) error {
	var req request.LogoutRequest
	if err := c.Bind(&req); err != nil {
		return response.InvalidRequestResponse
	}

	var tokenRecord model.RefreshToken
	if err := database.DB.Where("token = ?", req.RefreshToken).First(&tokenRecord).Error; err != nil {
		return response.InvalidRequestResponse
	}

	// 标记为已撤销
	tokenRecord.Revoked = true
	database.DB.Save(&tokenRecord)

	return c.NoContent(http.StatusOK)
}
