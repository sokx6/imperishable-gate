package routes

import (
	"imperishable-gate/internal/server/handlers"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes 注册所有 API 路由
func RegisterRoutes(e *echo.Echo) {
	// v1 API 分组
	v1 := e.Group("/api/v1")

	v1.POST("/ping", handlers.PingHandler)
	v1.POST("/links/add", handlers.AddHandler)
	v1.DELETE("/links/delete", handlers.DeleteHandler)
	v1.GET("/links/list", handlers.ListHandler)
	v1.POST("/names/add", handlers.AddNamesHandler)
	v1.POST("/remarks/add", handlers.AddRemarkHandler)
	v1.GET("/names/:name", handlers.ListByNameHandler)
	v1.POST("/links/tags", handlers.AddTagsByLinkHandler)
}
