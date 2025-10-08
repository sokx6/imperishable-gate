package request

type SendResetPasswordEmailByUsernameRequest struct {
	Username string `json:"username" validate:"required"`
}
