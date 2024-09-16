package http

import "errors"

var (
	errInvalidUser   = errors.New("user isn't valid")
	errIncorrectUser = errors.New("user isn't correct")
)
