package request

type SendResetPasswordEmailByRequest struct {
	Email string `json:"email" validate:"required,email"`
}
