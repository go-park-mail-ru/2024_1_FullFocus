package helper

import (
	"regexp"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

func ValidateField(value string, minLength, maxLength int) error {
	onlyValidSymbols := regexp.MustCompile(`^[A-Za-z0-9_]*$`).MatchString
	if !onlyValidSymbols(value) {
		return models.ErrInvalidField
	}
	if len(value) < minLength || len(value) > maxLength {
		return models.ErrInvalidField
	}
	return nil
}
