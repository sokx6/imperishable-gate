package handlers

import (
	types "imperishable-gate/internal"

	"imperishable-gate/internal/server/service"

	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUserHandler(c echo.Context) error {
	var req types.UserRegisterRequest
	if err := c.Bind(&req); err != nil || req.Username == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, types.InvalidRequestResponse)
	}

	// 调用服务层函数注册用户
	err := service.RegisterUser(req.Username, req.Email, req.Password)
	if err != nil {
		if err == service.ErrNameAlreadyExists {
			return c.JSON(http.StatusConflict, types.UserNameAlreadyExistsResponse)
		}
		if err == service.ErrEmailAlreadyExists {
			return c.JSON(http.StatusConflict, types.EmailAlreadyExistsResponse)
		}
		return c.JSON(http.StatusInternalServerError, types.DatabaseErrorResponse)
	}

	return c.JSON(http.StatusOK, types.RegisterSuccessResponse)
}
