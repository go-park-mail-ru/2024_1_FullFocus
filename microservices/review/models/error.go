package models

import "errors"

var (
	ErrNoProduct           = errors.New("no product")
	ErrNoReviews           = errors.New("no reviews found")
	ErrReviewAlreadyExists = errors.New("review exists")
)
