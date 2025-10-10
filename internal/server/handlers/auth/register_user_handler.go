package auth

import (
	"errors"
	"imperishable-gate/internal/server/service/common"
	emailService "imperishable-gate/internal/server/service/email"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUserHandler(c echo.Context) error {
	var req request.UserRegisterRequest
	if err := c.Bind(&req); err != nil || req.Username == "" || req.Email == "" || req.Password == "" {
		logger.Warning("Invalid registration request: %v", err)
		return response.InvalidRequestResponse
	}

	// 使用事务确保用户创建和邮件发送的原子性
	err := emailService.RegisterUserWithVerification(req.Username, req.Email, req.Password)
	if err != nil {
		// 用户名已存在
		if errors.Is(err, common.ErrNameAlreadyExists) {
			return response.UserNameAlreadyExistsResponse
		}
		// 邮箱已存在
		if errors.Is(err, common.ErrEmailAlreadyExists) {
			return response.EmailAlreadyExistsResponse
		}
		// 数据库错误
		if errors.Is(err, common.ErrDatabase) {
			logger.Error("Database error: %v", err)
			return response.DatabaseErrorResponse
		}
		// 邮件发送失败时，返回特殊错误提示用户使用重发功能
		logger.Error("Failed to send verification email: %v", err)
		return response.SendVerificationEmailFailedResponse
	}

	// 注册成功（包括邮件已发送）
	logger.Success("User registered successfully: %s", req.Username)
	return c.JSON(http.StatusOK, response.RegistrationSuccessResponse)
}
