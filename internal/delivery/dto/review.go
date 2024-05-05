package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type ProductReview struct {
	ID            uint   `json:"reviewID"`
	ProfileName   string `json:"profileName"`
	ProfileAvatar string `json:"profileAvatar"`
	CreatedAt     string `json:"createdAt"`
	Rating        uint   `json:"rating"`
	Advanatages   string `json:"advanatages"`
	Disadvantages string `json:"disadvantages"`
	Comment       string `json:"comment"`
}

func ConvertReviewsToDto(mm []models.ProductReview) []ProductReview {
	reviews := make([]ProductReview, 0)
	for _, m := range mm {
		reviews = append(reviews, ProductReview{
			ID:            m.ReviewID,
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
	ProductID     uint   `json:"productID"`
	Rating        uint   `json:"rating"`
	Comment       string `json:"comment"`
	Advantages    string `json:"advantages"`
	Disadvantages string `json:"disadvantages"`
}

func ConvertCreateReviewInputToModel(d CreateReviewInput) models.CreateProductReviewInput {
	return models.CreateProductReviewInput{
		ProductID:     d.ProductID,
		Rating:        d.Rating,
		Comment:       d.Comment,
		Advanatages:   d.Advantages,
		Disadvantages: d.Disadvantages,
	}
}
