package handlers

import (
	"net/http"
	"time"

	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"

	"github.com/labstack/echo/v4"
)

func RefreshTokenHandler(c echo.Context) error {
	var req request.RefreshRequest
	if err := c.Bind(&req); err != nil {
		return response.InvalidRequestResponse
	}

	var tokenRecord model.RefreshToken
	if err := database.DB.
		Where("token = ? AND expires_at > ? AND revoked = false", req.RefreshToken, time.Now()).
		First(&tokenRecord).Error; err != nil {

		return response.AuthenticationFailedResponse
	}

	// 验证通过，生成新的 Access Token
	var userInfo model.User
	if err := database.DB.Where("id = ?", tokenRecord.UserID).First(&userInfo).Error; err != nil {
		return response.DatabaseErrorResponse
	}

	newAccessToken, err := service.GenerateAccessToken(userInfo.ID, userInfo.Username)
	if err != nil {
		return response.InternalServerErrorResponse
	}

	tokenRecord.ExpiresAt = time.Now().Add(service.RefreshExpiry)
	database.DB.Save(&tokenRecord)

	return c.JSON(http.StatusOK, response.RefreshResponse{
		AccessToken: newAccessToken,
	})
}
