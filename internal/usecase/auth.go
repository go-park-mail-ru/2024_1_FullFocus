package usecase

import (
	"context"

	authgrpc "github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/auth/grpc"
	profilegrpc "github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/profile/grpc"
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

const (
	_defaultEmail       = "yourawesome@mail.ru"
	_defaultPhoneNumber = "70000000000"
)

type AuthUsecase struct {
	authClient    *authgrpc.Client
	profileClient *profilegrpc.Client
}

func NewAuthUsecase(ac *authgrpc.Client, pc *profilegrpc.Client) *AuthUsecase {
	return &AuthUsecase{
		authClient:    ac,
		profileClient: pc,
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
	sID, err := u.authClient.Login(ctx, login, password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidField) {
			return "", helper.NewValidationError("invalid input",
				"Невалидные данные")
		}
		return "", err
	}
	return sID, nil
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
	uID, sID, err := u.authClient.Signup(ctx, login, password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidField) {
			return "", helper.NewValidationError("invalid input",
				"Невалидные данные")
		}
		return "", err
	}
	if err = u.profileClient.CreateProfile(ctx, models.Profile{
		ID:          uID,
		FullName:    login,
		Email:       _defaultEmail,
		PhoneNumber: _defaultPhoneNumber,
	}); err != nil {
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
