package usecase

import (
	"context"
	"html"

	"github.com/go-park-mail-ru/2024_1_FullFocus/microservices/review/models"
	commonError "github.com/go-park-mail-ru/2024_1_FullFocus/pkg/error"
)

type Review interface {
	CreateProductReview(ctx context.Context, input models.CreateProductReviewInput) error
	GetProductReviews(ctx context.Context, input models.GetProductReviewsInput) ([]models.ProductReviewData, error)
}

type Usecase struct {
	repo Review
}

func NewUsecase(r Review) *Usecase {
	return &Usecase{
		repo: r,
	}
}

func (u *Usecase) CreateProductReview(ctx context.Context, input models.CreateProductReviewInput) error {
	if input.Rating > 5 {
		return commonError.ErrInvalidInput
	}
	return u.repo.CreateProductReview(ctx, input)
}

func (u *Usecase) GetProductReviews(ctx context.Context, input models.GetProductReviewsInput) ([]models.ProductReviewData, error) {
	reviews, err := u.repo.GetProductReviews(ctx, input)
	if err != nil {
		return nil, err
	}
	for i := range reviews {
		reviews[i].Advanatages = html.EscapeString(reviews[i].Advanatages)
		reviews[i].Disadvantages = html.EscapeString(reviews[i].Disadvantages)
		reviews[i].Comment = html.EscapeString(reviews[i].Comment)
	}
	return reviews, nil
}
