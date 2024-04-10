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
		GetAllProductCards(ctx context.Context, input models.GetAllProductsInput) ([]models.ProductCard, error)
		GetProductById(ctx context.Context, profileID uint, productID uint) (models.Product, error)
		GetProductsByCategoryId(ctx context.Context, input models.GetProductsByCategoryIDInput) ([]models.ProductCard, error)
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

	Profiles interface {
		UpdateProfile(ctx context.Context, uID uint, newProfile dto.ProfileData) error
		GetProfile(ctx context.Context, uID uint) (dto.ProfileData, error)
		CreateProfile(ctx context.Context, profile dto.ProfileData) (uint, error)
	}
	Carts interface {
		GetAllCartItems(ctx context.Context, uID uint) (models.CartContent, error)
		UpdateCartItem(ctx context.Context, uID, prID uint) (uint, error)
		DeleteCartItem(ctx context.Context, uID, prID uint) (uint, error)
		DeleteAllCartItems(ctx context.Context, uID uint) error
	}

	Categories interface {
		GetAllCategories(ctx context.Context) ([]models.Category, error)
	}
)
