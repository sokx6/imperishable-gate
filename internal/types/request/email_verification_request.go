package request

// EmailVerificationRequest 邮箱验证请求
type EmailVerificationRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

// ResendVerificationRequest 重新发送验证邮件请求
type ResendVerificationRequest struct {
	Email string `json:"email"`
}

type ResetPasswordByEmailRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Code        string `json:"code" validate:"required,len=6"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=64"`
}

type ResetPasswordByUsernameRequest struct {
	Username    string `json:"username" validate:"required"`
	Code        string `json:"code" validate:"required,len=6"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=64"`
}
