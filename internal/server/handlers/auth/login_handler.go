package auth

import (
	"errors"
	authService "imperishable-gate/internal/server/service/auth"
	"imperishable-gate/internal/server/service/common"
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
	loginResp, err := authService.GenerateJWTIfAuthenticated(req.Username, req.Password)
	if err != nil {
		switch {
		// 处理不同的失败原因，返回相应的HTTP状态码
		case errors.Is(err, common.ErrUsernameNotFound):
			return response.UserNotFoundResponse
		case errors.Is(err, common.ErrInvalidPassword):
			return response.AuthenticationFailedResponse
		case errors.Is(err, common.ErrEmailNotVerified):
			return response.EmailNotVerifiedResponse
		case errors.Is(err, common.ErrDatabase):
			return response.DatabaseErrorResponse
		default:
			return response.InternalServerErrorResponse
		}

	}

	return c.JSON(http.StatusOK, loginResp)
}
