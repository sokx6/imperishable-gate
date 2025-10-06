package utils

import (
	"net/http"
)

func AddAccessTokenToRequest(req *http.Request, accessToken string) {
	if req != nil && accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}
}
