//go:generate mockgen -source=repository.go -destination=./mocks/repository_mock.go
package repository

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type (
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

	Carts interface {
		GetAllCartItems(ctx context.Context, uID uint) ([]models.CartProduct, error)
		GetAllCartItemsID(ctx context.Context, uID uint) ([]models.CartItem, error)
		GetCartItemsAmount(ctx context.Context, uID uint) (uint, error)
		UpdateCartItem(ctx context.Context, uID uint, prID uint) (uint, error)
		DeleteCartItem(ctx context.Context, uID uint, orID uint) (uint, error)
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

	Reviews interface {
		GetProductReviews(ctx context.Context, input models.GetProductReviewsInput) ([]models.ProductReview, error)
		CreateProductReview(ctx context.Context, uID uint, input models.ProductReview) error
	}
)
