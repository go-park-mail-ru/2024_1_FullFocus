package promotion

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type PromotionClient interface {
	GetPromoProducts(ctx context.Context, amount uint) ([]models.PromoData, error)
}
