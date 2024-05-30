package models

import "errors"

var (
	ErrNoPolls          = errors.New("no polls found")
	ErrPollAlreadyRated = errors.New("poll is already rated")
)
