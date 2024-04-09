package usecase

import (
	"context"
	"html"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
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

func (u *ProfileUsecase) UpdateProfile(ctx context.Context, uID uint, newProfile dto.ProfileData) error {
	err := helper.ValidateField(newProfile.FullName, _minLoginLength, _maxLoginLength)
	if err != nil {
		return helper.NewValidationError("invalid fullname input",
			"Имя должно содержать от 4 до 32 букв английского алфавита или цифр")
	}
	err = helper.ValidateEmail(newProfile.Email)
	if err != nil {
		return helper.NewValidationError("invalid email input",
			"Имеил должен содержать @ и .")
	}

	object := model.Profile{
		ID:          newProfile.ID,
		Email:       newProfile.Email,
		FullName:    newProfile.FullName,
		PhoneNumber: newProfile.PhoneNumber,
		ImgSrc:      newProfile.ImgSrc,
	}
	err = u.profileRepo.UpdateProfile(ctx, uID, object)
	if err != nil {
		return model.ErrNoProfile
	}

	return nil
}

func (u *ProfileUsecase) GetProfile(ctx context.Context, uID uint) (dto.ProfileData, error) {
	profile, err := u.profileRepo.GetProfile(ctx, uID)
	escapedFullName := html.EscapeString(profile.FullName)
	escapedImfSrc := html.EscapeString(profile.ImgSrc)
	object := dto.ProfileData{
		ID:          profile.ID,
		Email:       profile.Email,
		FullName:    escapedFullName,
		PhoneNumber: profile.PhoneNumber,
		ImgSrc:      escapedImfSrc,
	}
	if err != nil {
		return dto.ProfileData{}, model.ErrNoProfile
	}
	return object, nil
}

func (u *ProfileUsecase) CreateProfile(ctx context.Context, profile dto.ProfileData) (uint, error) {
	object := model.Profile{
		ID:          profile.ID,
		Email:       profile.Email,
		FullName:    profile.FullName,
		PhoneNumber: profile.PhoneNumber,
		ImgSrc:      profile.ImgSrc,
	}
	uID, err := u.profileRepo.CreateProfile(ctx, object)
	if err != nil {
		return 0, model.ErrProfileAlreadyExists
	}
	return uID, nil
}
