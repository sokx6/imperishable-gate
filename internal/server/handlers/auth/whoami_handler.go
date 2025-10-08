package auth

import (
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
		return echo.NewHTTPError(http.StatusUnauthorized, "User information not found")
	}

	userInfo, ok := userInfoValue.(data.UserInfo)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User information not found")
	}

	resp := response.WhoamiResponse{
		Message:  "Success",
		UserInfo: &userInfo,
	}

	return c.JSON(http.StatusOK, resp)
}
