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

const (
	_notSet int = iota
	_male
	_female
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
	if len(newProfile.PhoneNumber) != 0 {
		if err := helper.ValidateNumber(newProfile.PhoneNumber, 5); err != nil {
			return helper.NewValidationError("invalid phone number",
				"Неверный номер телефона")
		}
	}
	if len(newProfile.Address) != 0 {
		if err := helper.ValidateAddress(newProfile.Address); err != nil {
			return helper.NewValidationError("invalid address",
				"Некорректный адрес")
		}
	}
	if newProfile.Gender != uint(_notSet) && newProfile.Gender != uint(_male) && newProfile.Gender != uint(_female) {
		return helper.NewValidationError("invalid gender",
			"Некорректный пол")
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
	profile.Address = html.EscapeString(profile.Address)
	profile.PhoneNum = html.EscapeString(profile.PhoneNum)
	email, err := u.authClient.GetUserEmailByUserID(ctx, uID)
	if err != nil {
		return models.FullProfile{}, err
	}
	email = html.EscapeString(email)
	return models.FullProfile{
		ProfileData: profile,
		Email:       email,
	}, nil
}

func (u *ProfileUsecase) GetProfileMetaInfo(ctx context.Context, uID uint) (models.ProfileMetaInfo, error) {
	info, err := u.profileClient.GetProfileMetaInfo(ctx, uID)
	if err != nil {
		return models.ProfileMetaInfo{}, err
	}
	info.FullName = html.EscapeString(info.FullName)
	amount, err := u.cartRepo.GetCartItemsAmount(ctx, uID)
	if err != nil {
		return models.ProfileMetaInfo{}, err
	}
	info.CartItemsAmount = amount
	return info, nil
}

func (u *ProfileUsecase) CreateProfile(ctx context.Context, pID uint) error {
	return u.profileClient.CreateProfile(ctx, pID)
}
