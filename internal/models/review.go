package models

type ProductReview struct {
	ReviewID      uint
	ProductID     uint
	AuthorName    string
	AuthorAvatar  string
	CreatedAt     string
	Rating        uint
	Advanatages   string
	Disadvantages string
	Comment       string
}

type ProductReviewData struct {
	ReviewID      uint
	ProfileID     uint
	CreatedAt     string
	Rating        uint
	Advanatages   string
	Disadvantages string
	Comment       string
}

type GetProductReviewsInput struct {
	ProductID    uint
	LastReviewID uint
	PageSize     uint
}

type CreateProductReviewInput struct {
	ProductID     uint
	Rating        uint
	Advanatages   string
	Disadvantages string
	Comment       string
}
