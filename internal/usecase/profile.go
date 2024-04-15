package usecase

import (
	"context"
	"html"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
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

func (u *ProfileUsecase) UpdateProfile(ctx context.Context, uID uint, newProfile models.ProfileUpdateInput) error {
	if err := helper.ValidateField(newProfile.FullName, _minLoginLength, _maxLoginLength); err != nil {
		return helper.NewValidationError("invalid fullname input",
			"Имя должно содержать от 4 до 32 букв английского алфавита или цифр")
	}
	if err := helper.ValidateEmail(newProfile.Email); err != nil {
		return helper.NewValidationError("invalid email input",
			"Email должен содержать @ и .")
	}
	if err := u.profileRepo.UpdateProfile(ctx, uID, newProfile); err != nil {
		return models.ErrNoProfile
	}
	return nil
}

func (u *ProfileUsecase) GetProfile(ctx context.Context, uID uint) (models.Profile, error) {
	profile, err := u.profileRepo.GetProfile(ctx, uID)
	if err != nil {
		return models.Profile{}, err
	}
	profile.FullName = html.EscapeString(profile.FullName)
	profile.ImgSrc = html.EscapeString(profile.ImgSrc)
	return profile, nil
}

func (u *ProfileUsecase) CreateProfile(ctx context.Context, profile models.Profile) (uint, error) {
	uID, err := u.profileRepo.CreateProfile(ctx, profile)
	if err != nil {
		return 0, err
	}
	return uID, nil
}
