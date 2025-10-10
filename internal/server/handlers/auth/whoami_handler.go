package auth

import (
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/data"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// WhoamiHandler 返回当前认证用户的信息
func WhoamiHandler(c echo.Context) error {
	// 从context中获取用户信息（由JWT中间件设置）
	userInfoValue := c.Get("userInfo")
	if userInfoValue == nil {
		logger.Warning("User information not found")
		return echo.NewHTTPError(http.StatusUnauthorized, "User information not found")
	}

	userInfo, ok := userInfoValue.(data.UserInfo)
	if !ok {
		logger.Warning("User information not found")
		return echo.NewHTTPError(http.StatusUnauthorized, "User information not found")
	}

	logger.Success("User info retrieved successfully for user %d", userInfo.UserID)
	return c.JSON(http.StatusOK, response.WhoamiResponse{
		Message:  "Success",
		UserInfo: &userInfo,
	})
}
