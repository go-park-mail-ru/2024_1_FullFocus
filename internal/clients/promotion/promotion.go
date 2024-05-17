package promotion

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type PromotionClient interface {
	CreatePromoProductInfo(ctx context.Context, input models.PromoData) error
	GetPromoProductsInfo(ctx context.Context, amount uint) ([]models.PromoData, error)
	DeletePromoProductInfo(ctx context.Context, ID uint) error
}
