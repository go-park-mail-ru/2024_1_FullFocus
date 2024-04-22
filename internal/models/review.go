package models

type ProductReview struct {
	AuthorName    string
	AuthorAvatar  string
	Date          string
	Rating        float32
	Advanatages   string
	Disadvantages string
	Comment       string
}

type GetProductReviewsInput struct {
	ProductID    uint
	LastReviewID uint
	PageSize     uint
}
