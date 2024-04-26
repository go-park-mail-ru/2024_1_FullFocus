package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository"
)

const (
	_productReviewsLimit   = 5
	_defaultReviewSortType = 4
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
	if input.PageSize == 0 {
		input.PageSize = _productReviewsLimit
	}
	if input.Sorting.ID != 3 && input.Sorting.ID != 4 {
		defaultSorting, err := helper.GetSortTypeByID(_defaultReviewSortType)
		if err != nil {
			return []models.ProductReview{}, models.ErrInternal
		}
		input.Sorting = defaultSorting
	}
	return u.reviewRepo.GetProductReviews(ctx, input)
}

func (u *ReviewUsecase) CreateProductReview(ctx context.Context, uID uint, input models.ProductReview) error {
	return u.reviewRepo.CreateProductReview(ctx, uID, input)
}
