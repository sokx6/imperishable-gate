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
var ErrInternalServer = errors.New("internal server error")
var ErrUnknown = errors.New("unknown error")
