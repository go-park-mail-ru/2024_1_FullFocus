package dao

import "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/review/models"

type ProductReviewTable struct {
	ReviewID      uint   `db:"id"`
	ProfileID     uint   `db:"profile_id"`
	Rating        uint   `db:"rating"`
	CreatedAt     string `db:"created_at"`
	Comment       string `db:"comments"`
	Advantages    string `db:"advantages"`
	Disadvantages string `db:"disadvantages"`
}

func ConvertReviewsToModels(tt []ProductReviewTable) []models.ProductReviewData {
	productReviews := make([]models.ProductReviewData, 0)
	for _, t := range tt {
		productReviews = append(productReviews, models.ProductReviewData{
			ReviewID:      t.ReviewID,
			ProfileID:     t.ProfileID,
			Rating:        t.Rating,
			CreatedAt:     t.CreatedAt,
			Comment:       t.Comment,
			Advanatages:   t.Advantages,
			Disadvantages: t.Disadvantages,
		})
	}
	return productReviews
}
