package models

import "github.com/pkg/errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrWrongPassword     = errors.New("wrong password")
	ErrUnableToHash      = errors.New("unable to hash")
	ErrInvalidInput      = errors.New("invalid input")
)
