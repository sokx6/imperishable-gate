package utils

import (
	"github.com/zalando/go-keyring"
)

const (
	service = "imperishable-gate"
	userKey = "refresh-token"
)

func SaveRefreshToken(refreshToken string) error {
	return keyring.Set(service, userKey, refreshToken)
}

func LoadRefreshToken() (string, error) {
	return keyring.Get(service, userKey)
}

func ClearTokens() {
	_ = keyring.Delete(service, userKey)
}
