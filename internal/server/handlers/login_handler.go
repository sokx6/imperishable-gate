package handlers

import (
	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// LoginHandler 处理登录请求
func LoginHandler(c echo.Context) error {
	var req request.LoginRequest
	if err := c.Bind(&req); err != nil || req.Username == "" || req.Password == "" {
		return response.InvalidRequestResponse
	}

	auth_result := service.GenerateJWTIfAuthenticated(req.Username, req.Password)
	if !auth_result.Success {
		switch auth_result.Message {
		case "用户未找到":
			return response.UserNotFoundResponse
		case "用户名或密码不正确":
			return response.AuthenticationFailedResponse
		case "数据库错误，请稍后重试":
			return response.DatabaseErrorResponse
		case "认证服务内部错误":
			return response.InternalServerErrorResponse
		default:
			return response.InternalServerErrorResponse
		}

	}

	return c.JSON(http.StatusOK, auth_result)
}
