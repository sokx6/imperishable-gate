package handlers

import (
	"net/http"

	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"

	"github.com/labstack/echo/v4"
)

func PingHandler(c echo.Context) error {
	var req request.PingRequest

	if err := c.Bind(&req); err != nil || req.Action != "ping" {
		return response.InvalidRequestResponse
	}

	return c.JSON(http.StatusOK, response.PongResponse)
}
