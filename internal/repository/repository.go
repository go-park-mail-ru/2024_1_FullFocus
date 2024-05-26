//go:generate mockgen -source=repository.go -destination=./mocks/repository_mock.go
package repository

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type (
	Avatars interface {
		GetAvatar(ctx context.Context, fileName string) (models.Avatar, error)
		UploadAvatar(ctx context.Context, fileName string, img models.Avatar) error
		DeleteAvatar(ctx context.Context, imageName string) error
	}

	Carts interface {
		GetAllCartItems(ctx context.Context, uID uint) ([]models.CartProduct, error)
		GetAllCartItemsInfo(ctx context.Context, uID uint) ([]models.CartItem, error)
		GetCartItemsAmount(ctx context.Context, uID uint) (uint, error)
		UpdateCartItem(ctx context.Context, uID, prID uint) (uint, error)
		DeleteCartItem(ctx context.Context, uID, orID uint) (uint, error)
		DeleteAllCartItems(ctx context.Context, uID uint) error
	}

	Categories interface {
		GetAllCategories(ctx context.Context) ([]models.Category, error)
		GetCategoryNameById(ctx context.Context, categoryID uint) (string, error)
	}

	Notifications interface {
		SendNotification(ctx context.Context, profileID uint, payload string) error
		CreateNotification(ctx context.Context, profileID uint, input models.CreateNotificationInput) error
		GetAllNotifications(ctx context.Context, profileID uint) ([]models.Notification, error)
		GetNotificationsAmount(ctx context.Context, profileID uint) (uint, error)
		MarkNotificationRead(ctx context.Context, id uint) error
	}

	Orders interface {
		Create(ctx context.Context, userID uint, sum uint, orderItems []models.OrderItem) (uint, error)
		GetOrderByID(ctx context.Context, orderID uint) (models.GetOrderPayload, error)
		GetAllOrders(ctx context.Context, profileID uint) ([]models.Order, error)
		GetProfileIDByOrderID(ctx context.Context, orderID uint) (uint, error)
		UpdateStatus(ctx context.Context, orderID uint, newStatus string) (string, error)
		Delete(ctx context.Context, orderID uint) error
	}

	Products interface {
		GetAllProductCards(ctx context.Context, input models.GetAllProductsInput) ([]models.ProductCard, error)
		GetProductByID(ctx context.Context, profileID uint, productID uint) (models.Product, error)
		GetTotalPrice(ctx context.Context, items []models.OrderItem) (uint, error)
		GetProductsByCategoryID(ctx context.Context, input models.GetProductsByCategoryIDInput) ([]models.ProductCard, error)
		GetProductsByQuery(ctx context.Context, input models.GetProductsByQueryInput) ([]models.ProductCard, error)
		GetProductsByIDs(ctx context.Context, profileID uint, IDs []uint) ([]models.Product, error)
		GetProductPriceByID(ctx context.Context, ID uint) (uint, error)
		MarkProduct(ctx context.Context, ID uint, promo bool) error
	}

	Promocodes interface {
		CreatePromocodeItem(ctx context.Context, info models.CreatePromocodeItemInput) error
		GetNewPromocode(ctx context.Context, sum uint) (uint, error)
		GetPromocodeItemByActivationCode(ctx context.Context, pID uint, code string) (models.PromocodeActivationTerms, error)
		GetAvailablePromocodes(ctx context.Context, profileID uint) ([]models.PromocodeItem, error)
		GetPromocodeByID(ctx context.Context, promocodeID uint) (models.Promocode, error)
		GetPromocodesAmount(ctx context.Context, profileID uint) (uint, error)
		ApplyPromocode(ctx context.Context, input models.ApplyPromocodeInput) (uint, error)
		DeleteUsedPromocode(ctx context.Context, id uint) error
	}

	Suggests interface {
		GetCategorySuggests(ctx context.Context, query string) ([]models.CategorySuggest, error)
		GetProductSuggests(ctx context.Context, query string) ([]string, error)
	}
)
