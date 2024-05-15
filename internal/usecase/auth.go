package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/auth"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/profile"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/pkg/errors"
)

const (
	_minLoginLength    = 4
	_maxLoginLength    = 32
	_minPasswordLength = 8
	_maxPasswordLength = 32
)

type AuthUsecase struct {
	authClient    auth.AuthClient
	profileClient profile.ProfileClient
}

func NewAuthUsecase(ac auth.AuthClient, pc profile.ProfileClient) *AuthUsecase {
	return &AuthUsecase{
		authClient:    ac,
		profileClient: pc,
	}
}

func (u *AuthUsecase) Login(ctx context.Context, email, password string) (string, error) {
	if err := helper.ValidateEmail(email); err != nil {
		return "", helper.NewValidationError("invalid email input",
			"Email должен содержать @ и .")
	}
	if err := helper.ValidateField(password, _minPasswordLength, _maxPasswordLength); err != nil {
		return "", helper.NewValidationError("invalid password input",
			"Пароль должен содержать от 8 до 32 букв английского алфавита или цифр")
	}
	sID, err := u.authClient.Login(ctx, email, password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidField) {
			return "", helper.NewValidationError("invalid input",
				"Невалидные данные")
		}
		return "", err
	}
	return sID, nil
}

func (u *AuthUsecase) Signup(ctx context.Context, input models.SignupData) (string, error) {
	if err := helper.ValidateField(input.Password, _minPasswordLength, _maxPasswordLength); err != nil {
		return "", helper.NewValidationError("invalid password input",
			"Пароль должен содержать от 8 до 32 букв английского алфавита или цифр")
	}
	if err := helper.ValidateEmail(input.Email); err != nil {
		return "", helper.NewValidationError("invalid email input",
			"Email должен содержать @ и .")
	}
	uID, sID, err := u.authClient.Signup(ctx, input.Email, input.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidField) {
			return "", helper.NewValidationError("invalid input",
				"Невалидные данные")
		}
		return "", err
	}
	if err = u.profileClient.CreateProfile(ctx, uID); err != nil {
		return "", err
	}
	return sID, nil
}

func (u *AuthUsecase) GetUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	uID, err := u.authClient.GetUserIDBySessionID(ctx, sID)
	if err != nil {
		return 0, err
	}
	return uID, nil
}

func (u *AuthUsecase) Logout(ctx context.Context, sID string) error {
	return u.authClient.Logout(ctx, sID)
}

func (u *AuthUsecase) IsLoggedIn(ctx context.Context, sID string) bool {
	return u.authClient.CheckAuth(ctx, sID)
}

func (u *AuthUsecase) UpdatePassword(ctx context.Context, userID uint, password string, newPassword string) error {
	if err := helper.ValidateField(newPassword, _minPasswordLength, _maxPasswordLength); err != nil {
		return helper.NewValidationError("invalid new password input",
			"Пароль должен содержать от 8 до 32 букв английского алфавита или цифр")
	}
	if err := u.authClient.UpdatePassword(ctx, userID, password, newPassword); err != nil {
		if errors.Is(err, models.ErrInvalidField) {
			return helper.NewValidationError("invalid input",
				"Невалидные данные")
		}
		return err
	}
	return nil
}
