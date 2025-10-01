package handlers

import (
	"net/http"

	types "imperishable-gate/internal"

	"github.com/labstack/echo/v4"
)

func PingHandler(c echo.Context) error {
	// 定义请求体结构
	var req types.PingRequest

	// 解析请求体到req
	if err := c.Bind(&req); err != nil || req.Action != "ping" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "action must be 'ping'",
		})
	}

	// 讲req解析为JSON返回
	return c.JSON(http.StatusOK, types.PingResponse{
		Action:  "ping",
		Message: "pong",
	})
}
