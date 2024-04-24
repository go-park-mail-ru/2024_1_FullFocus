package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type GetProductReviewsInput struct {
	ProductID    uint `json:"productID"`
	LastReviewID uint `json:"lastReviewID"`
	PageSize     uint `json:"limit"`
}

func ConvertGetReviewInputToModel(d GetProductReviewsInput) models.GetProductReviewsInput {
	return models.GetProductReviewsInput{
		ProductID:    d.ProductID,
		LastReviewID: d.LastReviewID,
		PageSize:     d.PageSize,
	}
}

type ProductReview struct {
	ProfileName   string  `json:"profileName"`
	ProfileAvatar string  `json:"profileAvatar"`
	CreatedAt     string  `json:"createdAt"`
	Rating        float32 `json:"rating"`
	Advanatages   string  `json:"advanatages"`
	Disadvantages string  `json:"disadvantages"`
	Comment       string  `json:"comment"`
}

func ConvertReviewsToDto(mm []models.ProductReview) []ProductReview {
	reviews := make([]ProductReview, 0)
	for _, m := range mm {
		reviews = append(reviews, ProductReview{
			ProfileName:   m.AuthorName,
			ProfileAvatar: m.AuthorAvatar,
			Rating:        m.Rating,
			CreatedAt:     m.CreatedAt,
			Advanatages:   m.Advanatages,
			Disadvantages: m.Disadvantages,
			Comment:       m.Comment,
		})
	}
	return reviews
}

type CreateReviewInput struct {
	ProductID     uint    `json:"productID"`
	Rating        float32 `json:"rating"`
	Comment       string  `json:"comment"`
	Advantages    string  `json:"advantages"`
	Disadvantages string  `json:"disadvantages"`
}

func ConvertCreateReviewInputToModel(d CreateReviewInput) models.ProductReview {
	return models.ProductReview{
		ProductID:     d.ProductID,
		Rating:        d.Rating,
		Comment:       d.Comment,
		Advanatages:   d.Advantages,
		Disadvantages: d.Disadvantages,
	}
}
