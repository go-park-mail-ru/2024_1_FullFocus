package models

import "github.com/pkg/errors"

var (
	ErrProfileAlreadyExists = errors.New("profile already exists")
	ErrInvalidInput         = errors.New("invalid input")
	ErrNoProfile            = errors.New("profile not found")
	ErrInternal             = errors.New("internal error")
)
