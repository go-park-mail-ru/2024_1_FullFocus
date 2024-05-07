package usecase

import (
	"context"
	"html"

	auth "github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/auth"
	profile "github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/profile"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
	"github.com/pkg/errors"
)

type ProfileUsecase struct {
	profileClient profile.ProfileClient
	authClient    auth.AuthClient
	cartRepo      repository.Carts
}

func NewProfileUsecase(pc profile.ProfileClient, ac auth.AuthClient, cr repository.Carts) *ProfileUsecase {
	return &ProfileUsecase{
		profileClient: pc,
		authClient:    ac,
		cartRepo:      cr,
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

func (u *ProfileUsecase) GetProfile(ctx context.Context, uID uint) (models.FullProfile, error) {
	profile, err := u.profileClient.GetProfileByID(ctx, uID)
	if err != nil {
		return models.FullProfile{}, err
	}
	profile.FullName = html.EscapeString(profile.FullName)
	login, err := u.authClient.GetUserLoginByUserID(ctx, uID)
	if err != nil {
		return models.FullProfile{}, err
	}
	return models.FullProfile{
		ProfileData: profile,
		Login:       login,
	}, nil
}

func (u *ProfileUsecase) GetProfileMetaInfo(ctx context.Context, uID uint) (models.ProfileMetaInfo, error) {
	info, err := u.profileClient.GetProfileMetaInfo(ctx, uID)
	if err != nil {
		return models.ProfileMetaInfo{}, err
	}
	amount, err := u.cartRepo.GetCartItemsAmount(ctx, uID)
	if err != nil {
		return models.ProfileMetaInfo{}, err
	}
	info.CartItemsAmount = amount
	return info, nil
}

func (u *ProfileUsecase) CreateProfile(ctx context.Context, profile models.Profile) error {
	return u.profileClient.CreateProfile(ctx, profile)
}
