package storage

import "errors"

var (
	errFindUser   = errors.New("user can't be matched")
	errUserExists = errors.New("can't add becouse user is exists")
)
