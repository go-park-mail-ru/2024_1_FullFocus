package models

import (
	"github.com/pkg/errors"
)

var (
	ErrNoSession                 = errors.New("no session")
	ErrUserNotFound              = errors.New("user not found")
	ErrUserAlreadyExists         = errors.New("user exists")
	ErrWrongPassword             = errors.New("wrong password")
	ErrNoProduct                 = errors.New("no product")
	ErrNoUserID                  = errors.New("no user ID")
	ErrInvalidField              = errors.New("invalid field input")
	ErrNoAvatar                  = errors.New("no avatar found")
	ErrNoAccess                  = errors.New("no access")
	ErrInvalidParameters         = errors.New("invalid parameters")
	ErrNoRowsFound               = errors.New("no rows found")
	ErrNoProfile                 = errors.New("no profile")
	ErrProfileAlreadyExists      = errors.New("profile exists")
	ErrEmptyCart                 = errors.New("no cart items found")
	ErrCantUpload                = errors.New("can't upload")
	ErrInternal                  = errors.New("internal server error")
	ErrNoReviews                 = errors.New("no reviews found")
	ErrReviewAlreadyExists       = errors.New("review exists")
	ErrNoPolls                   = errors.New("no polls found")
	ErrPollAlreadyRated          = errors.New("poll is already rated")
	ErrInvalidBenefitValue       = errors.New("invalid benefit value")
	ErrPromoProductAlreadyExists = errors.New("promo product already added")
	ErrProductNotFound           = errors.New("product not found")
)
