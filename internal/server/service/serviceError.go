package service

import (
	"errors"
)

var ErrNameAlreadyExists = errors.New("name already exists for another link")
var ErrInvalidRequest = errors.New("invalid request")
var ErrNameNotFound = errors.New("name not found")
var ErrDatabase = errors.New("database error")
var ErrLinkNotFound = errors.New("link not found")
var ErrCrawlFailed = errors.New("failed to crawl metadata from the URL")
var ErrUserNameExists = errors.New("username already exists")
var ErrEmailAlreadyExists = errors.New("email already registered")
var ErrInvalidRegister = errors.New("invalid register request")
var ErrUsernameNotFound = errors.New("username not found")
var ErrInvalidPassword = errors.New("invalid password")
var ErrAuthenticationFailed = errors.New("authentication failed")
var ErrLinkAlreadyExists = errors.New("link already exists")
var ErrUserNotFound = errors.New("user not found")
var ErrEmailAlreadyVerified = errors.New("email already verified")
var ErrInvalidVerificationCode = errors.New("invalid verification code")
var ErrVerificationExpired = errors.New("verification code expired")
var ErrEmailNotVerified = errors.New("email not verified")
var ErrVerificationCodeNotExpired = errors.New("verification code is still valid, please wait until it expires")
var ErrTooManyAttempts = errors.New("too many verification attempts, please request a new code")
var ErrResendTooSoon = errors.New("please wait at least 2 minutes before requesting a new code")
