// routes/routes.go
package routes

import (
	"imperishable-gate/internal/server/handlers"
	"imperishable-gate/internal/server/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	// 公共路由：不需要认证
	v1.POST("/register", handlers.RegisterUserHandler)
	v1.POST("/login", handlers.LoginHandler)
	v1.POST("/ping", handlers.PingHandler)
	v1.POST("/refresh", handlers.RefreshTokenHandler)
	v1.POST("/logout", handlers.LogoutHandler)
	v1.POST("/verify-email", handlers.VerifyEmailAndRegisterHandler)         // 新增：验证邮箱
	v1.POST("/resend-verification", handlers.ResendVerificationEmailHandler) // 新增：重发验证码
	v1.PATCH("/email/password", handlers.VerifyEmailAndResetPasswordHandler)
	v1.PATCH("/username/password", handlers.VerifyEmailByUsernameAndResetPasswordHandler)
	v1.PATCH("/email/password/request", handlers.SendResetPasswordEmailByEmailHandler)
	v1.PATCH("/username/password/request", handlers.SendResetPasswordEmailByUsernameHandler)
	// 创建需要认证的子分组（只加一次中间件）
	protected := v1.Group("", middlewares.JwtAuthMiddleware)

	// 下面这些路由都自动受保护，无需手动加中间件
	protected.GET("/names/:name", handlers.ListByNameHandler)
	protected.GET("/links", handlers.ListHandler)
	protected.GET("/links/search", handlers.SearchByKeywordHandler)
	protected.GET("/tags/:tag", handlers.ListByTagHandler)

	protected.POST("/links", handlers.AddHandler)
	protected.POST("/names", handlers.AddNamesHandler)
	protected.POST("/remarks", handlers.AddRemarkHandler)
	protected.POST("/tags", handlers.AddTagsByLinkHandler)
	protected.POST("/name/:name/remark", handlers.AddRemarkByNameHandler)
	protected.POST("/name/:name/tags", handlers.AddTagsByNameHandler)

	protected.PATCH("/links/watch", handlers.WatchByUrlHandler)
	protected.PATCH("/name/watch", handlers.WatchByNameHandler)
	protected.PATCH("/links/names/remove", handlers.DeleteNamesByLinkHandler)
	protected.PATCH("/links/by-url/tags/remove", handlers.DeleteTagsByLinkHandler)
	protected.PATCH("/:name/tags/remove", handlers.DeleteTagsByNameHandler)

	protected.DELETE("/links/name/:name", handlers.DeleteByNameHandler)
	protected.DELETE("/links", handlers.DeleteHandler)
}
