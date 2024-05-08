package usecase

import (
	"context"
	"html"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/models"
)

//go:generate mockgen -source=profile.go -destination=../repository/mocks/repository_mock.go
type Profile interface {
	CreateProfile(ctx context.Context, profile models.Profile) error
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
	if len(newProfile.FullName) == 0 || len(newProfile.Email) == 0 || len(newProfile.PhoneNumber) == 0 {
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
	return profile, nil
}

func (u *Usecase) GetProfileNamesByIDs(ctx context.Context, pIDs []uint) ([]string, error) {
	return u.repo.GetProfileNamesByIDs(ctx, pIDs)
}

func (u *Usecase) GetProfileMetaInfo(ctx context.Context, pID uint) (models.ProfileMetaInfo, error) {
	return u.repo.GetProfileMetaInfo(ctx, pID)
}

func (u *Usecase) GetProfileNamesAvatarsByIDs(ctx context.Context, pIDs []uint) ([]models.ProfileNameAvatar, error) {
	profiles, err := u.repo.GetProfileNamesAvatarsByIDs(ctx, pIDs)
	if err != nil {
		return nil, err
	}
	orderedProfiles := make([]models.ProfileNameAvatar, len(profiles))
	for i, pID := range pIDs {
		for _, p := range profiles {
			if pID == p.ID {
				orderedProfiles[i] = p
				break
			}
		}
	}
	return orderedProfiles, nil
}

func (u *Usecase) CreateProfile(ctx context.Context, profile models.Profile) error {
	return u.repo.CreateProfile(ctx, profile)
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
