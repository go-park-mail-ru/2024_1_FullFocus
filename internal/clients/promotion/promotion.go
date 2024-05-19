package promotion

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type PromotionClient interface {
	CreatePromoProductInfo(ctx context.Context, input models.PromoData) error
	GetAllPromoProductsIDs(ctx context.Context) ([]uint, error)
	GetPromoProductsInfoByIDs(ctx context.Context, prIDs []uint) ([]models.PromoData, error)
	DeletePromoProductInfo(ctx context.Context, ID uint) error
}
