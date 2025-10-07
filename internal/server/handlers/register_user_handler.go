package handlers

import (
	"errors"
	"imperishable-gate/internal/server/service"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUserHandler(c echo.Context) error {
	var req request.UserRegisterRequest
	if err := c.Bind(&req); err != nil || req.Username == "" || req.Email == "" || req.Password == "" {
		return response.InvalidRequestResponse
	}

	// 调用服务层函数注册用户
	err := service.RegisterUser(req.Username, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrNameAlreadyExists) {
			return response.UserNameAlreadyExistsResponse
		}
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			return response.EmailAlreadyExistsResponse
		}
		return response.DatabaseErrorResponse
	}

	return c.JSON(http.StatusOK, response.RegisterSuccessResponse)
}
