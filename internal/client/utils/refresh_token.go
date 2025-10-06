package utils

import (
	"bytes"
	"encoding/json"
	"imperishable-gate/internal/types/request"
	"imperishable-gate/internal/types/response"
	"net/http"
)

func RefreshToken(refreshToken, addr string) (string, error) {
	reqBody := request.RefreshRequest{RefreshToken: refreshToken}
	reqBytes, _ := json.Marshal(reqBody)

	resp, err := http.Post(addr+"/api/v1/refresh", "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var accessToken response.RefreshResponse
	json.NewDecoder(resp.Body).Decode(&accessToken)

	if accessToken.AccessToken == "" {
		return "", ErrNoAccessToken
	}

	return accessToken.AccessToken, nil
}
