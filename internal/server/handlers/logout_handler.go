package handlers

import (
	"net/http"

	types "imperishable-gate/internal"
	"imperishable-gate/internal/model"
	"imperishable-gate/internal/server/database"

	"github.com/labstack/echo/v4"
)

func LogoutHandler(c echo.Context) error {
	var req types.LogoutRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "请提供 refresh_token")
	}

	var tokenRecord model.RefreshToken
	if err := database.DB.Where("token = ?", req.RefreshToken).First(&tokenRecord).Error; err != nil {
		return c.NoContent(http.StatusOK) // 无视错误，幂等退出
	}

	// 标记为已撤销
	tokenRecord.Revoked = true
	database.DB.Save(&tokenRecord)

	return c.NoContent(http.StatusOK)
}
