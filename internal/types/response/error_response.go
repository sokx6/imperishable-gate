package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 错误提示消息（前端可通过 HTTP Status 区分错误类型）
var InvalidUrlFormatResponse = echo.NewHTTPError(http.StatusBadRequest, "Invalid URL format")
var InvalidRequestResponse = echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
var DatabaseErrorResponse = echo.NewHTTPError(http.StatusInternalServerError, "Database error")
var RemarkExistsResponse = echo.NewHTTPError(http.StatusConflict, "Remark already exists")
var NameNotFoundResponse = echo.NewHTTPError(http.StatusNotFound, "Name not found")
var NameExistsResponse = echo.NewHTTPError(http.StatusConflict, "Name already exists")
var LinkNotFoundResponse = echo.NewHTTPError(http.StatusNotFound, "Link not found")
var LinkExistsResponse = echo.NewHTTPError(http.StatusConflict, "Link already exists")
var TagNotFoundResponse = echo.NewHTTPError(http.StatusNotFound, "Tag not found")
var UserNameAlreadyExistsResponse = echo.NewHTTPError(http.StatusConflict, "Username already exists")
var EmailAlreadyExistsResponse = echo.NewHTTPError(http.StatusConflict, "Email already registered")
var UserNotFoundResponse = echo.NewHTTPError(http.StatusNotFound, "User not found")
var AuthenticationFailedResponse = echo.NewHTTPError(http.StatusUnauthorized, "Authentication failed")
var InternalServerErrorResponse = echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
var UnknownErrorResponse = echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
var EmailNotVerifiedResponse = echo.NewHTTPError(http.StatusForbidden, "Email not verified")

// 邮箱验证相关错误
var EmailCannotBeEmptyResponse = echo.NewHTTPError(http.StatusBadRequest, "Email cannot be empty")
var EmailOrCodeCannotBeEmptyResponse = echo.NewHTTPError(http.StatusBadRequest, "Email or code cannot be empty")
var EmailNotRegisteredResponse = echo.NewHTTPError(http.StatusBadRequest, "Email not registered")
var EmailAlreadyVerifiedResponse = echo.NewHTTPError(http.StatusBadRequest, "Email is already verified")
var InvalidVerificationCodeResponse = echo.NewHTTPError(http.StatusBadRequest, "Invalid or already used verification code")
var VerificationExpiredResponse = echo.NewHTTPError(http.StatusBadRequest, "Verification code has expired, please request a new one")
var VerificationCodeNotExpiredResponse = echo.NewHTTPError(http.StatusTooManyRequests, "Verification code is still valid. Please wait until it expires (15 minutes) before requesting a new one.")
var ResendTooSoonResponse = echo.NewHTTPError(http.StatusTooManyRequests, "Please wait at least 2 minutes before requesting a new verification code")
var TooManyAttemptsResponse = echo.NewHTTPError(http.StatusTooManyRequests, "Too many verification attempts. Please request a new verification code")
var InvalidVerificationTokenResponse = echo.NewHTTPError(http.StatusBadRequest, "Invalid verification token")
var VerificationFailedResponse = echo.NewHTTPError(http.StatusInternalServerError, "Verification failed")
var ResendVerificationEmailFailedResponse = echo.NewHTTPError(http.StatusInternalServerError, "Failed to resend verification email")
var SendVerificationEmailFailedResponse = echo.NewHTTPError(http.StatusInternalServerError, "Account created, but failed to send verification email. Please use 'resend' to get a new code.")
var UsernameOrCodeCannotBeEmptyResponse = echo.NewHTTPError(http.StatusBadRequest, "Username or code cannot be empty")
var SendResetPasswordEmailFailedResponse = echo.NewHTTPError(http.StatusInternalServerError, "Failed to send reset password email")
var UsernameCannotBeEmptyResponse = echo.NewHTTPError(http.StatusBadRequest, "Username cannot be empty")
var UsernameNotRegisteredResponse = echo.NewHTTPError(http.StatusBadRequest, "Username not registered")
