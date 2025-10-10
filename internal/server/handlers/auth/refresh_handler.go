package auth

import (
	"net/http"
	"time"

	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	authService "imperishable-gate/internal/server/service/auth"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RefreshTokenHandler(c echo.Context) error {
	var req request.RefreshRequest
	// 解析请求体
	if err := c.Bind(&req); err != nil {
		logger.Warning("Invalid refresh token request: %v", err)
		return response.InvalidRequestResponse
	}

	var tokenRecord model.RefreshToken
	// 查找并验证刷新令牌
	if err := database.DB.
		Where("token = ? AND expires_at > ? AND revoked = false", req.RefreshToken, time.Now()).
		First(&tokenRecord).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.AuthenticationFailedResponse
		}
		logger.Error("Database error: %v", err)
		return response.DatabaseErrorResponse
	}

	// 验证通过，生成新的 Access Token
	var userInfo model.User
	if err := database.DB.Where("id = ?", tokenRecord.UserID).First(&userInfo).Error; err != nil {
		logger.Error("Database error: %v", err)
		return response.DatabaseErrorResponse
	}

	newAccessToken, err := authService.GenerateAccessToken(userInfo.ID, userInfo.Username)
	if err != nil {
		logger.Error("Failed to generate access token: %v", err)
		return response.InternalServerErrorResponse
	}

	logger.Success("Access token refreshed successfully for user %d", userInfo.ID)
	return c.JSON(http.StatusOK, response.RefreshResponse{
		AccessToken: newAccessToken,
	})
}
