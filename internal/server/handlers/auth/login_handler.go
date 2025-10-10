package auth

import (
	"errors"
	authService "imperishable-gate/internal/server/service/auth"
	"imperishable-gate/internal/server/service/common"
	"imperishable-gate/internal/server/utils/logger"
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
		logger.Warning("Invalid login request: %v", err)
		return response.InvalidRequestResponse
	}
	// 调用服务层进行认证
	loginResp, err := authService.GenerateJWTIfAuthenticated(req.Username, req.Password)
	if err != nil {
		switch {
		// 处理不同的失败原因，返回相应的HTTP状态码
		case errors.Is(err, common.ErrUsernameNotFound):
			logger.Warning("Username not found: %s", req.Username)
			return response.UserNotFoundResponse
		case errors.Is(err, common.ErrInvalidPassword):
			logger.Warning("Invalid password for username: %s", req.Username)
			return response.AuthenticationFailedResponse
		case errors.Is(err, common.ErrEmailNotVerified):
			logger.Warning("Email not verified for username: %s", req.Username)
			return response.EmailNotVerifiedResponse
		case errors.Is(err, common.ErrDatabase):
			logger.Error("Database error: %v", err)
			return response.DatabaseErrorResponse
		default:
			logger.Error("Login failed: %v", err)
			return response.InternalServerErrorResponse
		}

	}

	logger.Success("User logged in successfully: %s", req.Username)
	return c.JSON(http.StatusOK, loginResp)
}
