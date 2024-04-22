package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

type ReviewUsecase struct {
	reviewRepo repository.Reviews
}

func NewReviewUsecase(r repository.Reviews) *ReviewUsecase {
	return &ReviewUsecase{
		reviewRepo: r,
	}
}

func (u *ReviewUsecase) GetProductReviews(ctx context.Context, input models.GetProductReviewsInput) ([]models.ProductReview, error) {
	return u.reviewRepo.GetProductReviews(ctx, input)
}
