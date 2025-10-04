package routes

import (
	"imperishable-gate/internal/server/handlers"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes 注册所有 API 路由
func RegisterRoutes(e *echo.Echo) {
	// v1 API 分组
	v1 := e.Group("/api/v1")

	v1.GET("/names/:name", handlers.ListByNameHandler)
	v1.GET("/links", handlers.ListHandler)
	v1.GET("/tags/:tag", handlers.ListByTagHandler)

	v1.POST("/ping", handlers.PingHandler)
	v1.POST("/links", handlers.AddHandler)
	v1.POST("/names", handlers.AddNamesHandler)
	v1.POST("/remarks", handlers.AddRemarkHandler)
	v1.POST("/tags", handlers.AddTagsByLinkHandler)
	v1.POST("/name/:name/remark", handlers.AddRemarkByNameHandler)
	v1.POST("/name/:name/tags", handlers.AddTagsByNameHandler)

	v1.DELETE("/links/name/:name", handlers.DeleteByNameHandler)
	v1.DELETE("/links/delete", handlers.DeleteHandler)
}
