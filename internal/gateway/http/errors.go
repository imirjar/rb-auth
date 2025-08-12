package http

import "errors"

var (
	errInvalidUser     = errors.New("user isn't valid")
	errIncorrectUser   = errors.New("user isn't correct")
	errUserIsNotExists = errors.New("user isn't exists")
)

var (
	errInvalidToken  = errors.New("token isn't valid")
	errMissingHeader = errors.New("missing Authorization Header")
)

var (
	errInternal = errors.New("smth went wrong")
)
