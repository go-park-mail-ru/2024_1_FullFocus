package usecase

import (
	"context"
	model "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type ProfileUsecase struct {
	profileRepo repository.Profiles
}

func NewProfileUsecase(pr repository.Profiles) *ProfileUsecase {
	return &ProfileUsecase{
		profileRepo: pr,
	}
}

func (u *ProfileUsecase) UpdateProfile(ctx context.Context, username string, newProfile model.Profile) error {
	err := helper.ValidateField(newProfile.User.Username, _minLoginLength, _maxLoginLength)
	if err != nil {
		return model.NewValidationError("invalid login input",
			"Логин должен содержать от 4 до 32 букв английского алфавита или цифр")
	}

	err = helper.ValidateField(newProfile.User.Password, _minPasswordLength, _maxPasswordLength)
	if err != nil {
		return model.NewValidationError("invalid password input",
			"Пароль должен содержать от 8 до 32 букв английского алфавита или цифр")
	}

	if newProfile.Image.PayloadSize > 4e+7 {
		return model.NewValidationError("invalid image input",
			"Аватарка не может превышать размер 40 мегабайт")
	}

	err = u.profileRepo.UpdateProfile(ctx, username, newProfile)
	if err != nil {
		return model.ErrNoProfile
	}

	return nil
}

func (u *ProfileUsecase) GetProfile(ctx context.Context, username string) (model.Profile, error) {
	// XSS проверка username
	profile, err := u.profileRepo.GetProfile(ctx, username)
	if err != nil {
		return model.Profile{}, model.ErrNoProfile
	}
	return profile, nil
}

func (u *ProfileUsecase) CreateProfile(ctx context.Context, profile model.Profile) (uint, error) {
	uID, err := u.profileRepo.CreateProfile(ctx, profile)
	if err != nil {
		return 0, model.ErrProfileAlreadyExists
	}
	return uID, nil
}
