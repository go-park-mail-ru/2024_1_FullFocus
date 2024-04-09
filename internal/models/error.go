package models

import (
	"fmt"

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
	ErrNoAccess          = errors.New("no access")
	ErrInvalidParameters = errors.New("invalid parameters")
	ErrNoRowsFound       = errors.New("no rows found")
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

func (ve *ValidationError) WithCode(code int) *ErrResponse {
	return &ErrResponse{
		Status: code,
		Msg:    ve.msg,
		MsgRus: ve.msgRus,
	}
}
