package usecase

import (
	"context"
	"html"

	profilegrpc "github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/profile/grpc"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/pkg/errors"
)

type ProfileUsecase struct {
	profileClient *profilegrpc.Client
}

func NewProfileUsecase(pr *profilegrpc.Client) *ProfileUsecase {
	return &ProfileUsecase{
		profileClient: pr,
	}
}

func (u *ProfileUsecase) UpdateProfile(ctx context.Context, uID uint, newProfile models.ProfileUpdateInput) error {
	if newProfile.FullName == "" {
		return helper.NewValidationError("invalid name input",
			"Имя не может быть пустым")
	}
	if err := helper.ValidateEmail(newProfile.Email); err != nil {
		return helper.NewValidationError("invalid email input",
			"Email должен содержать @ и .")
	}
	if err := helper.ValidateNumber(newProfile.PhoneNumber, 5); err != nil {
		return helper.NewValidationError("invalid phone number",
			"Неверный номер телефона")
	}
	if err := u.profileClient.UpdateProfile(ctx, uID, newProfile); err != nil {
		if errors.Is(err, models.ErrInvalidField) {
			return helper.NewValidationError("invalid input",
				"Невалидные данные")
		}
		return err
	}
	return nil
}

func (u *ProfileUsecase) GetProfile(ctx context.Context, uID uint) (models.Profile, error) {
	profile, err := u.profileClient.GetProfileByID(ctx, uID)
	if err != nil {
		return models.Profile{}, err
	}
	profile.FullName = html.EscapeString(profile.FullName)
	return profile, nil
}

func (u *ProfileUsecase) CreateProfile(ctx context.Context, profile models.Profile) error {
	return u.profileClient.CreateProfile(ctx, profile)
}
