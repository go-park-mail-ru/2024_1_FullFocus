package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type Auth interface {
	Login(ctx context.Context, login string, password string) (string, error)
	Signup(ctx context.Context, login string, password string) (string, string, error)
	GetUserIDBySessionID(ctx context.Context, sID string) (uint, error)
	Logout(ctx context.Context, sID string) error
	IsLoggedIn(ctx context.Context, isID string) bool
}

type Products interface {
	GetProducts(ctx context.Context, lastID, limit int) ([]models.Product, error)
}
