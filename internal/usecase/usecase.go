package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

//go:generate mockgen -source=usecase.go -destination=./mocks/usecase_mock.go
type (
	Auth interface {
		Login(ctx context.Context, login string, password string) (string, error)
		Signup(ctx context.Context, login string, password string) (string, string, error)
		GetUserIDBySessionID(ctx context.Context, sID string) (uint, error)
		Logout(ctx context.Context, sID string) error
		IsLoggedIn(ctx context.Context, isID string) bool
	}

	Products interface {
		GetProducts(ctx context.Context, lastID, limit int) ([]models.Product, error)
	}

	Avatars interface {
		UploadAvatar(ctx context.Context, img dto.Image) error
		DeleteAvatar(ctx context.Context) error
	}

	Profiles interface {
		UpdateProfile(ctx context.Context, username string, newProfile models.Profile) error
		GetProfile(ctx context.Context, username string) (models.Profile, error)
		CreateProfile(ctx context.Context, profile models.Profile) (uint, error)
	}
)
