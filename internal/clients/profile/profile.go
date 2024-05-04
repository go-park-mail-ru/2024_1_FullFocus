package profile

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type ProfileClient interface {
	CreateProfile(ctx context.Context, profile models.Profile) error
	DeleteAvatarByProfileID(ctx context.Context, pID uint) (string, error)
	GetAvatarByID(ctx context.Context, pID uint) (string, error)
	GetProfileByID(ctx context.Context, pID uint) (models.Profile, error)
	GetProfileNamesByIDs(ctx context.Context, pIDs []uint) ([]string, error)
	UpdateAvatarByProfileID(ctx context.Context, pID uint, avatarName string) (string, error)
	UpdateProfile(ctx context.Context, pID uint, newProfile models.ProfileUpdateInput) error
	GetProfileMetaInfo(ctx context.Context, uID uint) (models.ProfileMetaInfo, error)
}
