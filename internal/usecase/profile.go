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

func (u *ProfileUsecase) UpdateProfile(ctx context.Context, uID uint, newProfile model.Profile) error {
	err := helper.ValidateField(newProfile.FullName, _minLoginLength, _maxLoginLength)
	if err != nil {
		return model.NewValidationError("invalid fullname input",
			"Имя должно содержать от 4 до 32 букв английского алфавита или цифр")
	}

	err = u.profileRepo.UpdateProfile(ctx, uID, newProfile)
	if err != nil {
		return model.ErrNoProfile
	}

	return nil
}

func (u *ProfileUsecase) GetProfile(ctx context.Context, uID uint) (model.Profile, error) {
	// XSS проверка username
	profile, err := u.profileRepo.GetProfile(ctx, uID)
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
