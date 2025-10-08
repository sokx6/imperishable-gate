// routes/routes.go
package routes

import (
	"imperishable-gate/internal/server/handlers/auth"
	"imperishable-gate/internal/server/handlers/common"
	"imperishable-gate/internal/server/handlers/email"
	"imperishable-gate/internal/server/handlers/links"
	"imperishable-gate/internal/server/handlers/names"
	"imperishable-gate/internal/server/handlers/remarks"
	"imperishable-gate/internal/server/handlers/tags"
	"imperishable-gate/internal/server/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	// 公共路由：不需要认证
	v1.POST("/register", auth.RegisterUserHandler)
	v1.POST("/login", auth.LoginHandler)
	v1.POST("/ping", common.PingHandler)
	v1.POST("/refresh", auth.RefreshTokenHandler)
	v1.POST("/logout", auth.LogoutHandler)
	v1.POST("/verify-email", email.VerifyEmailAndRegisterHandler)
	v1.POST("/resend-verification", email.ResendVerificationEmailHandler)
	v1.PATCH("/email/password", email.VerifyEmailAndResetPasswordHandler)
	v1.PATCH("/username/password", email.VerifyEmailByUsernameAndResetPasswordHandler)
	v1.PATCH("/email/password/request", email.SendResetPasswordEmailByEmailHandler)
	v1.PATCH("/username/password/request", email.SendResetPasswordEmailByUsernameHandler)

	protected := v1.Group("", middlewares.JwtAuthMiddleware)

	protected.GET("/names/:name", links.ListByNameHandler)
	protected.GET("/links", links.ListHandler)
	protected.GET("/links/search", links.SearchByKeywordHandler)
	protected.GET("/tags/:tag", links.ListByTagHandler)

	protected.POST("/links", links.AddHandler)
	protected.POST("/names", names.AddNamesHandler)
	protected.POST("/remarks", remarks.AddRemarkHandler)
	protected.POST("/tags", tags.AddTagsByLinkHandler)
	protected.POST("/name/:name/remark", remarks.AddRemarkByNameHandler)
	protected.POST("/name/:name/tags", tags.AddTagsByNameHandler)

	protected.PATCH("/links/watch", links.WatchByUrlHandler)
	protected.PATCH("/name/watch", links.WatchByNameHandler)
	protected.PATCH("/links/names/remove", names.DeleteNamesByLinkHandler)
	protected.PATCH("/links/by-url/tags/remove", tags.DeleteTagsByLinkHandler)
	protected.PATCH("/:name/tags/remove", tags.DeleteTagsByNameHandler)

	protected.DELETE("/links/name/:name", links.DeleteByNameHandler)
	protected.DELETE("/links", links.DeleteHandler)
}
