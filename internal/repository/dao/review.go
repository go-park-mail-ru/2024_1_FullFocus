package dao

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type ProductReviewTable struct {
	ProfileName   string `db:"full_name"`
	ProfileAvatar string `db:"imgsrc"`
	Rating        uint   `db:"rating"`
	CreatedAt     string `db:"created_at"`
	Comment       string `db:"comments"`
	Advantages    string `db:"advantages"`
	Disadvantages string `db:"disadvantages"`
}

func ConvertReviewsToModels(tt []ProductReviewTable) []models.ProductReview {
	productReviews := make([]models.ProductReview, 0)
	for _, t := range tt {
		productReviews = append(productReviews, models.ProductReview{
			AuthorName:    t.ProfileName,
			AuthorAvatar:  t.ProfileAvatar,
			Rating:        t.Rating,
			CreatedAt:     t.CreatedAt,
			Comment:       t.Comment,
			Advanatages:   t.Advantages,
			Disadvantages: t.Disadvantages,
		})
	}
	return productReviews
}
