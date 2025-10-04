package handlers

import (
	"net/http"

	types "imperishable-gate/internal"

	"github.com/labstack/echo/v4"
)

func PingHandler(c echo.Context) error {
	var req types.PingRequest

	if err := c.Bind(&req); err != nil || req.Action != "ping" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "action must be 'ping'",
		})
	}

	return c.JSON(http.StatusOK, types.PongResponse)
}
