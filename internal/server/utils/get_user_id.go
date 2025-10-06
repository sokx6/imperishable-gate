// utils/user.go 或 middlewares/auth.go

package utils

import (
	"imperishable-gate/internal/types/data"

	"github.com/labstack/echo/v4"
)

// GetUserID 从 Echo 上下文中提取用户 ID
func GetUserID(c echo.Context) (uint, bool) {
	if raw := c.Get("userInfo"); raw != nil {
		if userInfo, ok := raw.(data.UserInfo); ok {
			return userInfo.UserID, true
		}
	}
	return 0, false
}
