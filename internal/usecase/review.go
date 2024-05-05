package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/profile"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/review"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

const _productReviewsLimit = 5

type ReviewUsecase struct {
	profileClient profile.ProfileClient
	reviewClient  review.ReviewClient
}

func NewReviewUsecase(pc profile.ProfileClient, rc review.ReviewClient) *ReviewUsecase {
	return &ReviewUsecase{
		profileClient: pc,
		reviewClient:  rc,
	}
}

func (u *ReviewUsecase) GetProductReviews(ctx context.Context, input models.GetProductReviewsInput) ([]models.ProductReview, error) {
	if input.PageSize == 0 {
		input.PageSize = _productReviewsLimit
	}
	reviewsData, err := u.reviewClient.GetProductReviews(ctx, input)
	if err != nil {
		return nil, err
	}
	reviews := make([]models.ProductReview, 0)
	for _, r := range reviewsData {
		authorData, err := u.profileClient.GetProfileByID(ctx, r.ProfileID)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, models.ProductReview{
			ReviewID:      r.ReviewID,
			AuthorName:    authorData.FullName,
			AuthorAvatar:  authorData.AvatarName,
			Rating:        r.Rating,
			Advanatages:   r.Advanatages,
			Disadvantages: r.Disadvantages,
			Comment:       r.Comment,
			CreatedAt:     r.CreatedAt,
		})
	}
	return reviews, nil
}

func (u *ReviewUsecase) CreateProductReview(ctx context.Context, uID uint, input models.CreateProductReviewInput) error {
	return u.reviewClient.CreateProductReview(ctx, uID, input)
}
