package models

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
	Limit        uint
	SortingQuery string
}

type CreateProductReviewInput struct {
	ProductID     uint
	ProfileID     uint
	Rating        uint
	Advanatages   string
	Disadvantages string
	Comment       string
}
