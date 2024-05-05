package helper

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
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

type ValidationError struct {
	msg    string
	msgRus string
}

func NewValidationError(msg, msgRus string) *ValidationError {
	return &ValidationError{
		msg:    msg,
		msgRus: msgRus,
	}
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("error: %s, rus: %s", ve.msg, ve.msgRus)
}

func (ve *ValidationError) WithCode(code int) *dto.ErrResponse {
	return &dto.ErrResponse{
		Status: code,
		Msg:    ve.msg,
		MsgRus: ve.msgRus,
	}
}

func ValidateNumber(value string, length int) error {
	onlyValidSymbols := regexp.MustCompile(`^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`).MatchString
	if !onlyValidSymbols(value) {
		return models.ErrInvalidField
	}
	if len(value) < length {
		return models.ErrInvalidField
	}
	return nil
}

func ValidateEmail(value string) error {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return models.ErrInvalidField
	}
	return nil
}
