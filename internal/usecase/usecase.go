package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

//go:generate mockgen -source=usecase.go -destination=./mocks/usecase_mock.go
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
		UploadAvatar(ctx context.Context, img dto.Image, profileID uint) error
		DeleteAvatar(ctx context.Context, profileID uint) error
	}

	Orders interface {
		Create(ctx context.Context, input models.CreateOrderInput) (uint, error)
		GetOrderByID(ctx context.Context, profileID uint, orderingID uint) (models.GetOrderPayload, error)
		GetAllOrders(ctx context.Context, profileID uint) ([]models.Order, error)
		Delete(ctx context.Context, profileID uint, orderingID uint) error
	}
)
