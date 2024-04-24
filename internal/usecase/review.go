package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

const _productReviewsLimit = 5

type ReviewUsecase struct {
	reviewRepo repository.Reviews
}

func NewReviewUsecase(r repository.Reviews) *ReviewUsecase {
	return &ReviewUsecase{
		reviewRepo: r,
	}
}

func (u *ReviewUsecase) GetProductReviews(ctx context.Context, input models.GetProductReviewsInput) ([]models.ProductReview, error) {
	if input.PageSize == 0 {
		input.PageSize = _productReviewsLimit
	}
	return u.reviewRepo.GetProductReviews(ctx, input)
}
