package review

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type ReviewClient interface {
	CreateProductReview(ctx context.Context, uID uint, input models.CreateProductReviewInput) error
	GetProductReviews(ctx context.Context, input models.GetProductReviewsDataInput) ([]models.ProductReviewData, error)
}
