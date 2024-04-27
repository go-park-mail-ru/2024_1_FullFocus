package error

import "errors"

var (
	ErrInternal      = errors.New("internal server error")
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)
