package usecase

import (
	"context"
	"html"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/models"
)

//go:generate mockgen -source=profile.go -destination=../repository/mocks/repository_mock.go
type Profile interface {
	CreateProfile(ctx context.Context, pID uint) error
	GetProfile(ctx context.Context, uID uint) (models.Profile, error)
	GetProfileNamesByIDs(ctx context.Context, pIDs []uint) ([]string, error)
	GetProfileMetaInfo(ctx context.Context, pID uint) (models.ProfileMetaInfo, error)
	GetProfileNamesAvatarsByIDs(ctx context.Context, pIDs []uint) ([]models.ProfileNameAvatar, error)
	UpdateProfile(ctx context.Context, uID uint, profileNew models.ProfileUpdateInput) error
	UpdateAvatarByProfileID(ctx context.Context, uID uint, imgSrc string) (string, error)
	GetAvatarByProfileID(ctx context.Context, uID uint) (string, error)
	DeleteAvatarByProfileID(ctx context.Context, uID uint) (string, error)
}

type Usecase struct {
	repo Profile
}

func NewUsecase(pr Profile) *Usecase {
	return &Usecase{
		repo: pr,
	}
}

func (u *Usecase) UpdateProfile(ctx context.Context, uID uint, newProfile models.ProfileUpdateInput) error {
	if len(newProfile.FullName) == 0 {
		return models.ErrInvalidInput
	}
	return u.repo.UpdateProfile(ctx, uID, newProfile)
}

func (u *Usecase) GetProfile(ctx context.Context, uID uint) (models.Profile, error) {
	profile, err := u.repo.GetProfile(ctx, uID)
	if err != nil {
		return models.Profile{}, err
	}
	profile.FullName = html.EscapeString(profile.FullName)
	profile.Address = html.EscapeString(profile.Address)
	profile.PhoneNumber = html.EscapeString(profile.PhoneNumber)
	return profile, nil
}

func (u *Usecase) GetProfileNamesByIDs(ctx context.Context, pIDs []uint) ([]string, error) {
	names, err := u.repo.GetProfileNamesByIDs(ctx, pIDs)
	if err != nil {
		return nil, err
	}
	for i, name := range names {
		names[i] = html.EscapeString(name)
	}
	return names, nil
}

func (u *Usecase) GetProfileMetaInfo(ctx context.Context, pID uint) (models.ProfileMetaInfo, error) {
	info, err := u.repo.GetProfileMetaInfo(ctx, pID)
	if err != nil {
		return models.ProfileMetaInfo{}, err
	}
	info.FullName = html.EscapeString(info.FullName)
	return info, nil
}

func (u *Usecase) GetProfileNamesAvatarsByIDs(ctx context.Context, pIDs []uint) ([]models.ProfileNameAvatar, error) {
	data, err := u.repo.GetProfileNamesAvatarsByIDs(ctx, pIDs)
	if err != nil {
		return nil, err
	}
	for i, d := range data {
		data[i].FullName = html.EscapeString(d.FullName)
	}
	return data, nil
}

func (u *Usecase) CreateProfile(ctx context.Context, pID uint) error {
	return u.repo.CreateProfile(ctx, pID)
}

func (u *Usecase) UpdateAvatarByProfileID(ctx context.Context, uID uint, imgSrc string) (string, error) {
	return u.repo.UpdateAvatarByProfileID(ctx, uID, imgSrc)
}

func (u *Usecase) GetAvatarByProfileID(ctx context.Context, uID uint) (string, error) {
	return u.repo.GetAvatarByProfileID(ctx, uID)
}

func (u *Usecase) DeleteAvatarByProfileID(ctx context.Context, uID uint) (string, error) {
	return u.repo.DeleteAvatarByProfileID(ctx, uID)
}
