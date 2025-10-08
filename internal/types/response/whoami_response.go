package response

import "imperishable-gate/internal/types/data"

type WhoamiResponse struct {
	Message  string         `json:"message"`
	UserInfo *data.UserInfo `json:"user_info"`
}
