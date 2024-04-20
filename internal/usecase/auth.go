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

const (
	_defaultEmail       = "yourawesome@mail.ru"
	_defaultPhoneNumber = "70000000000"
)

type AuthUsecase struct {
	userRepo    repository.Users
	sessionRepo repository.Sessions
	profileRepo repository.Profiles
}

func NewAuthUsecase(ur repository.Users, sr repository.Sessions, pr repository.Profiles) *AuthUsecase {
	return &AuthUsecase{
		userRepo:    ur,
		sessionRepo: sr,
		profileRepo: pr,
	}
}

func (u *AuthUsecase) Login(ctx context.Context, login string, password string) (string, error) {
	if err := helper.ValidateField(login, _minLoginLength, _maxLoginLength); err != nil {
		return "", helper.NewValidationError("invalid login input",
			"Логин должен содержать от 4 до 32 букв английского алфавита или цифр")
	}
	if err := helper.ValidateField(password, _minPasswordLength, _maxPasswordLength); err != nil {
		return "", helper.NewValidationError("invalid password input",
			"Пароль должен содержать от 8 до 32 букв английского алфавита или цифр")
	}

	user, err := u.userRepo.GetUser(ctx, login)
	if err != nil {
		return "", models.ErrNoUser
	}
	if err = helper.CheckPassword(password, user.PasswordHash); err != nil {
		return "", models.ErrWrongPassword
	}
	// if password != user.PasswordHash {
	// 	   return "", models.ErrWrongPassword
	// }
	return u.sessionRepo.CreateSession(ctx, user.ID), nil
}

func (u *AuthUsecase) Signup(ctx context.Context, login string, password string) (string, error) {
	if err := helper.ValidateField(login, _minLoginLength, _maxLoginLength); err != nil {
		return "", helper.NewValidationError("invalid login input",
			"Логин должен содержать от 4 до 32 букв английского алфавита или цифр")
	}
	if err := helper.ValidateField(password, _minPasswordLength, _maxPasswordLength); err != nil {
		return "", helper.NewValidationError("invalid password input",
			"Пароль должен содержать от 8 до 32 букв английского алфавита или цифр")
	}
	passwordHash, err := helper.HashPassword(password)
	if err != nil {
		return "", err
	}
	user := models.User{
		Username:     login,
		PasswordHash: passwordHash,
	}
	uID, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return "", models.ErrUserAlreadyExists
	}
	profile := models.Profile{
		ID:          uID,
		FullName:    login,
		Email:       _defaultEmail,
		PhoneNumber: _defaultPhoneNumber,
	}
	_, err = u.profileRepo.CreateProfile(ctx, profile)
	if err != nil {
		return "", models.ErrProfileAlreadyExists
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
