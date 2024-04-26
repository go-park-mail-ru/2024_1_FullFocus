package models

type ProductReview struct {
	ProductID     uint
	AuthorName    string
	AuthorAvatar  string
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
	Sorting      SortType
}
