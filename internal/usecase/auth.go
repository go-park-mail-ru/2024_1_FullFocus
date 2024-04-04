package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

const (
	_minLoginLength    = 4
	_maxLoginLength    = 32
	_minPasswordLength = 8
	_maxPasswordLength = 32
)

type AuthUsecase struct {
	userRepo    repository.Users
	sessionRepo repository.Sessions
}

func NewAuthUsecase(ur repository.Users, sr repository.Sessions) *AuthUsecase {
	return &AuthUsecase{
		userRepo:    ur,
		sessionRepo: sr,
	}
}

func (u *AuthUsecase) Login(ctx context.Context, login string, password string) (string, error) {
	if err := helper.ValidateField(login, _minLoginLength, _maxLoginLength); err != nil {
		return "", models.NewValidationError("invalid login input",
			"Логин должен содержать от 4 до 32 букв английского алфавита или цифр")
	}
	if err := helper.ValidateField(password, _minPasswordLength, _maxPasswordLength); err != nil {
		return "", models.NewValidationError("invalid password input",
			"Пароль должен содержать от 8 до 32 букв английского алфавита или цифр")
	}

	user, err := u.userRepo.GetUser(ctx, login)
	if err != nil {
		return "", models.ErrNoUser
	}

	// TODO: add hasher from helper
	// err = helper.CheckPassword(password, user.PasswordHash)
	if password != user.PasswordHash {
		return "", models.ErrWrongPassword
	}
	return u.sessionRepo.CreateSession(ctx, user.ID), nil
}

func (u *AuthUsecase) Signup(ctx context.Context, login string, password string) (string, error) {
	if err := helper.ValidateField(login, _minLoginLength, _maxLoginLength); err != nil {
		return "", models.NewValidationError("invalid login input",
			"Логин должен содержать от 4 до 32 букв английского алфавита или цифр")
	}
	if err := helper.ValidateField(password, _minPasswordLength, _maxPasswordLength); err != nil {
		return "", models.NewValidationError("invalid password input",
			"Пароль должен содержать от 8 до 32 букв английского алфавита или цифр")
	}

	// TODO: add hasher from helper
	// passwordHash, err := helper.HashPassword(password)
	// if err != nil {
	// 	return "", err
	// }
	user := models.User{
		Username:     login,
		PasswordHash: password,
	}
	uID, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return "", models.ErrUserAlreadyExists
	}
	sID := u.sessionRepo.CreateSession(ctx, uID)
	return sID, nil
}

func (u *AuthUsecase) GetUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	return u.sessionRepo.GetUserIDBySessionID(ctx, sID)
}

func (u *AuthUsecase) Logout(ctx context.Context, sID string) error {
	return u.sessionRepo.DeleteSession(ctx, sID)
}

func (u *AuthUsecase) IsLoggedIn(ctx context.Context, sID string) bool {
	return u.sessionRepo.SessionExists(ctx, sID)
}
