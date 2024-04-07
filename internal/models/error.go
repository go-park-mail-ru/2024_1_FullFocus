package models

import (
	"fmt"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/pkg/errors"
)

var (
	ErrNoSession         = errors.New("no session")
	ErrNoUser            = errors.New("no user")
	ErrUserAlreadyExists = errors.New("user exists")
	ErrWrongPassword     = errors.New("wrong password")
	ErrNoProduct         = errors.New("no product")
	ErrNoUserID          = errors.New("no user ID")
	ErrInvalidField      = errors.New("invalid field input")
	ErrNoAvatar          = errors.New("no avatar found")
	ErrEmptyCart         = errors.New("no cart items found")
)

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

// ? перенос в dto
func (ve *ValidationError) WithCode(code int) *dto.ErrResponse {
	return &dto.ErrResponse{
		Status: code,
		Msg:    ve.msg,
		MsgRus: ve.msgRus,
	}
}
