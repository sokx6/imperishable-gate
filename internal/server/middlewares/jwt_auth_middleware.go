package middlewares

import (
	authService "imperishable-gate/internal/server/service/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

func JwtAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
		}

		tokenString := authHeader[7:]

		if userInfo, err := authService.ParseJWTAndValidate(tokenString, string(authService.JWTSecret)); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		} else {
			c.Logger().Infof("Authenticated user: %s (ID: %d)", userInfo.Username, userInfo.UserID)
			c.Set("userInfo", userInfo)
		}

		return next(c)
	}
}
