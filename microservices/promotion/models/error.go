package models

import (
	"errors"
)

var (
	ErrProductNotFound           = errors.New("product not found")
	ErrPromoProductAlreadyExists = errors.New("promo product already added")
)
