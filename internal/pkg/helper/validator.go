package helper

import (
	"unicode"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

func ValidateField(value string, minLength, maxLength int) error {
	if len(value) < minLength || len(value) > maxLength {
		return models.ErrInvalidField
	}
	for _, r := range value {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return models.ErrInvalidField
		}
	}
	return nil
}
