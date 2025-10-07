package response

import "imperishable-gate/internal/types/data"

type Response struct {
	Message string      `json:"message"`
	Links   []data.Link `json:"links,omitempty"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}
