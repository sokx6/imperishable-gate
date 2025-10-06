package utils

import (
	"errors"
)

var ErrNoAccessToken = errors.New("no access token found")

var ErrNoRefreshToken = errors.New("no refresh token found")

var ErrTokenExpired = errors.New("access token expired, please refresh")

var ErrInvalidToken = errors.New("invalid token, please login again")
