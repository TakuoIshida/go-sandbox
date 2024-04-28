package errors

import "errors"

var (
	ErrNotExist     = errors.New("file does not exist")
	ErrUnAuthorized = errors.New("unauthorized")
)
