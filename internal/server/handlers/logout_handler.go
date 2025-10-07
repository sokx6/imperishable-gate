package handlers

import (
	"net/http"

	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/types/data"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"

	"github.com/labstack/echo/v4"
)

func LogoutHandler(c echo.Context) error {
	var req request.LogoutRequest
	if err := c.Bind(&req); err != nil {
		return response.InvalidRequestResponse
	}

	// 获取当前已认证用户的 ID
	userInfo := c.Get("userInfo").(data.UserInfo)

	var tokenRecord model.RefreshToken
	if err := database.DB.Where("token = ?", req.RefreshToken).First(&tokenRecord).Error; err != nil {
		return response.InvalidRequestResponse
	}

	// 验证 token 是否属于当前用户
	if tokenRecord.UserID != userInfo.UserID {
		return response.AuthenticationFailedResponse
	}

	tokenRecord.Revoked = true
	database.DB.Save(&tokenRecord)
	return c.NoContent(http.StatusOK)
}
