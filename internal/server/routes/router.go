package routes

import (
	"imperishable-gate/internal/server/handlers"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes 注册所有 API 路由
func RegisterRoutes(e *echo.Echo) {
	// v1 API 分组
	v1 := e.Group("/api/v1")

	// Ping 接口
	v1.POST("/ping", handlers.PingHandler)
	v1.POST("/links/add", handlers.AddHandler)
}
