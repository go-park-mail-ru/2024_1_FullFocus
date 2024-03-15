package repository

import (
	"context"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type Users interface {
	CreateUser(ctx context.Context, user models.User) (uint, error)
	GetUser(ctx context.Context, login string) (models.User, error)
}

type Sessions interface {
	CreateSession(ctx context.Context, userID uint) string
	GetUserIDBySessionID(ctx context.Context, sID string) (uint, error)
	SessionExists(ctx context.Context, sID string) bool
	DeleteSession(ctx context.Context, sID string) error
}

type Products interface {
	GetProducts(ctx context.Context, lastID, limit int) ([]models.Product, error)
}
