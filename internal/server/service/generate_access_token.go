package service

import (
	"imperishable-gate/internal/types/data"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userID uint, username string) (string, error) {
	claims := data.CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "imperishable-gate",
			Subject:   "access",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}
