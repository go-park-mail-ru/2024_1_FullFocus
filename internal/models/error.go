package models

import "github.com/pkg/errors"

var ErrNoSession = errors.New("no session")
var ErrNoUserID = errors.New("no user ID")