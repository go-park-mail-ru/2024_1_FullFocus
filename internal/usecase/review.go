package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/profile"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/clients/review"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
)

const (
	_productReviewsLimit   = 5
	_defaultReviewSortType = 4
)

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
	sorting, err := validateReviewSorting(input.Sorting)
	if err != nil {
		return []models.ProductReview{}, err
	}
	reviewsData, err := u.reviewClient.GetProductReviews(ctx, models.GetProductReviewsDataInput{
		ProductID:    input.ProductID,
		PageSize:     input.PageSize,
		LastReviewID: input.LastReviewID,
		SortingQuery: sorting,
	})
	if err != nil {
		return nil, err
	}
	profileIDs := make([]uint, len(reviewsData))
	for i := range len(profileIDs) {
		profileIDs[i] = reviewsData[i].ProfileID
	}
	profileData, err := u.profileClient.GetProfileNamesAvatarsByIDs(ctx, profileIDs)
	if err != nil {
		return nil, err
	}
	reviews := make([]models.ProductReview, 0)
	for i, r := range reviewsData {
		reviews = append(reviews, models.ProductReview{
			ReviewID:      r.ReviewID,
			AuthorName:    profileData[i].FullName,
			AuthorAvatar:  profileData[i].AvatarName,
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

func validateReviewSorting(input models.SortType) (string, error) {
	if input.ID != 3 && input.ID != 4 {
		defaultSorting, err := helper.GetSortTypeByID(_defaultReviewSortType)
		if err != nil {
			return "", models.ErrInternal
		}
		return defaultSorting.QueryPart, nil
	}
	return input.QueryPart, nil
}
