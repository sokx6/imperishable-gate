package handlers

import (
	"fmt"
	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// LoginHandler 处理登录请求
func LoginHandler(c echo.Context) error {
	var req request.LoginRequest
	// 解析并验证请求体
	if err := c.Bind(&req); err != nil || req.Username == "" || req.Password == "" {
		return response.InvalidRequestResponse
	}
	// 调用服务层进行认证
	login_result := service.GenerateJWTIfAuthenticated(req.Username, req.Password)
	if !login_result.Success {
		switch login_result.Message {
		// 处理不同的失败原因，返回相应的HTTP状态码
		case "User not found":
			return response.UserNotFoundResponse
		case "Invalid username or password":
			return response.AuthenticationFailedResponse
		case "Database error, please try again later":
			return response.DatabaseErrorResponse
		case "Internal authentication service error":
			fmt.Println("Internal authentication service error")
			return response.InternalServerErrorResponse
		default:
			return response.InternalServerErrorResponse
		}

	}

	return c.JSON(http.StatusOK, login_result)
}
