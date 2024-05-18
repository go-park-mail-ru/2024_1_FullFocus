package models

import (
	"errors"
)

var (
	ErrProductNotFound           = errors.New("product not found")
	ErrPromoProductAlreadyExists = errors.New("tnis product already has sales")
)
