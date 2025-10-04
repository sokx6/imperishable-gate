package service

import (
	"errors"
)

var ErrNameAlreadyExists = errors.New("name already exists for another link")
var ErrInvalidRequest = errors.New("invalid request")
