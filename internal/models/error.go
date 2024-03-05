package models

import "github.com/pkg/errors"

var (
	ErrNoSession         = errors.New("no session")
	ErrNoUser            = errors.New("no user")
	ErrUserAlreadyExists = errors.New("user exists")
	ErrWrongPassword     = errors.New("wrong password")
	ErrNoProduct         = errors.New("no product")
	ErrWrongUsername     = errors.New("incorrect username")
	ErrWeakPassword      = errors.New("too weak password")
	ErrNoUserID          = errors.New("no user ID")
)
