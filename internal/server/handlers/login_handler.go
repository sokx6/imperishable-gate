package handlers

import (
	types "imperishable-gate/internal"
	"imperishable-gate/internal/server/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// LoginHandler 处理登录请求
func LoginHandler(c echo.Context) error {
	var req types.LoginRequest
	if err := c.Bind(&req); err != nil || req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}

	auth_result := service.GenerateJWTIfAuthenticated(req.Username, req.Password)
	if !auth_result.Success {
		switch auth_result.Message {
		case "用户未找到":
			return c.JSON(http.StatusNotFound, types.UserNotFoundResponse)
		case "用户名或密码不正确":
			return c.JSON(http.StatusUnauthorized, types.AuthenticationFailedResponse)
		case "数据库错误，请稍后重试":
			return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
		case "认证服务内部错误":
			return c.JSON(http.StatusInternalServerError, types.InternalServerErrorResponse)
		default:
			return c.JSON(http.StatusInternalServerError, types.InternalServerErrorResponse)
		}

	}

	return c.JSON(http.StatusOK, types.Response{
		Code:    0,
		Message: "Login successful",
		Data:    auth_result.Token,
	})
}
