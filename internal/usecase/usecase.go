//go:generate mockgen -source=usecase.go -destination=./mocks/usecase_mock.go
package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type (
	Auth interface {
		Login(ctx context.Context, login string, password string) (string, error)
		Signup(ctx context.Context, login string, password string) (string, error)
		GetUserIDBySessionID(ctx context.Context, sID string) (uint, error)
		Logout(ctx context.Context, sID string) error
		IsLoggedIn(ctx context.Context, isID string) bool
		UpdatePassword(ctx context.Context, userID uint, password string, newPassword string) error
	}

	Products interface {
		GetAllProductCards(ctx context.Context, input models.GetAllProductsInput) ([]models.ProductCard, error)
		GetProductByID(ctx context.Context, profileID uint, productID uint) (models.Product, error)
		GetProductsByCategoryID(ctx context.Context, input models.GetProductsByCategoryIDInput) (models.GetProductsByCategoryIDPayload, error)
		GetProductsByQuery(ctx context.Context, input models.GetProductsByQueryInput) ([]models.ProductCard, error)
	}

	Avatars interface {
		GetAvatar(ctx context.Context, profileID uint) (models.Avatar, error)
		UploadAvatar(ctx context.Context, profileID uint, img models.Avatar) error
		DeleteAvatar(ctx context.Context, uID uint) error
	}

	Orders interface {
		Create(ctx context.Context, input models.CreateOrderInput) (uint, error)
		GetOrderByID(ctx context.Context, profileID uint, orderingID uint) (models.GetOrderPayload, error)
		GetAllOrders(ctx context.Context, profileID uint) ([]models.Order, error)
		Delete(ctx context.Context, profileID uint, orderingID uint) error
	}

	Profiles interface {
		UpdateProfile(ctx context.Context, uID uint, newProfile models.ProfileUpdateInput) error
		GetProfile(ctx context.Context, uID uint) (models.Profile, error)
		GetProfileMetaInfo(ctx context.Context, uID uint) (models.ProfileMetaInfo, error)
		CreateProfile(ctx context.Context, profile models.Profile) error
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

	Reviews interface {
		GetProductReviews(ctx context.Context, input models.GetProductReviewsInput) ([]models.ProductReview, error)
		CreateProductReview(ctx context.Context, uID uint, input models.ProductReview) error
	}

	Suggests interface {
		GetSuggestions(ctx context.Context, query string) (models.Suggestion, error)
	}
)
