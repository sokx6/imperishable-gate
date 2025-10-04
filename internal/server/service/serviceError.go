package service

import (
	"errors"
)

var ErrNameAlreadyExists = errors.New("name already exists for another link")
var ErrInvalidRequest = errors.New("invalid request")
var ErrNameNotFound = errors.New("name not found")
var ErrDatabase = errors.New("database error")
var ErrLinkNotFound = errors.New("link not found")
