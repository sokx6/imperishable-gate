package response

type LoginResult struct {
	Success      bool   `json:"success"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Message      string `json:"message"`
}
