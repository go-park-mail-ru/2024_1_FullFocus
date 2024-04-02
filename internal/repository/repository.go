package repository

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source=repository.go -destination=./mocks/repository_mock.go
type (
	Users interface {
		CreateUser(ctx context.Context, user models.User) (uuid.UUID, error)
		GetUser(ctx context.Context, login string) (models.User, error)
	}

	Sessions interface {
		CreateSession(ctx context.Context, userID uuid.UUID) string
		GetUserIDBySessionID(ctx context.Context, sID string) (uuid.UUID, error)
		SessionExists(ctx context.Context, sID string) bool
		DeleteSession(ctx context.Context, sID string) error
	}

	Products interface {
		GetProducts(ctx context.Context, lastID, limit int) ([]models.Product, error)
	}

	Avatars interface {
		UploadAvatar(ctx context.Context, img models.Image) error
		DeleteAvatar(ctx context.Context, imageName string) error
	}
)
