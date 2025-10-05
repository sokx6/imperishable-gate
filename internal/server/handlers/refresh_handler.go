package handlers

import (
	"net/http"
	"time"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/service"

	"github.com/labstack/echo/v4"
)

func RefreshTokenHandler(c echo.Context) error {
	var req types.RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "无效请求")
	}

	var tokenRecord model.RefreshToken
	if err := database.DB.
		Where("token = ? AND expires_at > ? AND revoked = false", req.RefreshToken, time.Now()).
		First(&tokenRecord).Error; err != nil {

		return echo.NewHTTPError(http.StatusUnauthorized, "无效或已过期的刷新令牌")
	}

	// 验证通过，生成新的 Access Token
	var userInfo model.User
	if err := database.DB.Where("id = ?", tokenRecord.UserID).First(&userInfo).Error; err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "用户不存在")
	}

	newAccessToken, err := service.GenerateAccessToken(userInfo.ID, userInfo.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "无法生成新访问令牌")
	}

	tokenRecord.ExpiresAt = time.Now().Add(service.RefreshExpiry)
	database.DB.Save(&tokenRecord)

	return c.JSON(http.StatusOK, map[string]string{
		"token": newAccessToken,
	})
}
