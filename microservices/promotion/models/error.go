package models

import (
	"errors"
)

var (
	ErrProductNotFound    = errors.New("product not found")
	ErrPromoAlreadyExists = errors.New("tnis product already has sales")
)
