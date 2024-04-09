//go:generate mockgen -source=usecase.go -destination=./mocks/usecase_mock.go
package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type (
	Auth interface {
		Login(ctx context.Context, login string, password string) (string, error)
		Signup(ctx context.Context, login string, password string) (string, error)
		GetUserIDBySessionID(ctx context.Context, sID string) (uint, error)
		Logout(ctx context.Context, sID string) error
		IsLoggedIn(ctx context.Context, isID string) bool
	}

	Products interface {
		GetProducts(ctx context.Context, lastID, limit int) ([]models.Product, error)
	}

	Avatars interface {
		UploadAvatar(ctx context.Context, img dto.Image, uID uint) error
		DeleteAvatar(ctx context.Context, uID uint) error
	}

	Carts interface {
		GetAllCartItems(ctx context.Context, uID uint) (models.CartContent, error)
		UpdateCartItem(ctx context.Context, uID, prID uint) (uint, error)
		DeleteCartItem(ctx context.Context, uID, prID uint) (uint, error)
		DeleteAllCartItems(ctx context.Context, uID uint) error
	}
)
