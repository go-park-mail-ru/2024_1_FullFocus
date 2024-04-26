//go:generate mockgen -source=repository.go -destination=./mocks/repository_mock.go
package repository

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type (
	Users interface {
		CreateUser(ctx context.Context, user models.User) (uint, error)
		GetUser(ctx context.Context, login string) (models.User, error)
	}

	Sessions interface {
		CreateSession(ctx context.Context, userID uint) string
		GetUserIDBySessionID(ctx context.Context, sID string) (uint, error)
		SessionExists(ctx context.Context, sID string) bool
		DeleteSession(ctx context.Context, sID string) error
	}

	Products interface {
		GetAllProductCards(ctx context.Context, input models.GetAllProductsInput) ([]models.ProductCard, error)
		GetProductByID(ctx context.Context, profileID uint, productID uint) (models.Product, error)
		GetProductsByCategoryID(ctx context.Context, input models.GetProductsByCategoryIDInput) ([]models.ProductCard, error)
		GetProductsByQuery(ctx context.Context, input models.GetProductsByQueryInput) ([]models.ProductCard, error)
	}

	Avatars interface {
		GetAvatar(ctx context.Context, fileName string) (models.Avatar, error)
		UploadAvatar(ctx context.Context, fileName string, img models.Avatar) error
		DeleteAvatar(ctx context.Context, imageName string) error
	}

	Orders interface {
		Create(ctx context.Context, userID uint, orderItems []models.OrderItem) (uint, error)
		GetOrderByID(ctx context.Context, orderID uint) (models.GetOrderPayload, error)
		GetAllOrders(ctx context.Context, profileID uint) ([]models.Order, error)
		GetProfileIDByOrderID(ctx context.Context, orderID uint) (uint, error)
		Delete(ctx context.Context, orderID uint) error
	}

	Profiles interface {
		UpdateProfile(ctx context.Context, uID uint, profileNew models.ProfileUpdateInput) error
		GetProfile(ctx context.Context, uID uint) (models.Profile, error)
		CreateProfile(ctx context.Context, profile models.Profile) (uint, error)
		UpdateAvatarByProfileID(ctx context.Context, uID uint, imgSrc string) (string, error)
		GetAvatarByProfileID(ctx context.Context, uID uint) (string, error)
		DeleteAvatarByProfileID(ctx context.Context, uID uint) (string, error)
	}

	Carts interface {
		GetAllCartItems(ctx context.Context, uID uint) ([]models.CartProduct, error)
		GetAllCartItemsID(ctx context.Context, uID uint) ([]models.CartItem, error)
		UpdateCartItem(ctx context.Context, uID, prID uint) (uint, error)
		DeleteCartItem(ctx context.Context, uID, orID uint) (uint, error)
		DeleteAllCartItems(ctx context.Context, uID uint) error
	}

	Categories interface {
		GetAllCategories(ctx context.Context) ([]models.Category, error)
		GetCategoryNameById(ctx context.Context, categoryID uint) (string, error)
	}

	Suggests interface {
		GetCategorySuggests(ctx context.Context, query string) ([]models.CategorySuggest, error)
		GetProductSuggests(ctx context.Context, query string) ([]string, error)
	}
)
