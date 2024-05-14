package models

type ProductCard struct {
	ID     uint
	Name   string
	Price  uint
	ImgSrc string
	Seller string
	Rating uint
	Amount uint
}

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       uint
	ImgSrc      string
	Seller      string
	Rating      uint
	Amount      uint
	Categories  []string
}

type GetAllProductsInput struct {
	ProfileID uint
	PageNum   uint
	PageSize  uint
	Sorting   SortType
}

type GetProductsByQueryInput struct {
	Query     string
	ProfileID uint
	PageNum   uint
	PageSize  uint
	Sorting   SortType
}

type GetProductsByCategoryIDInput struct {
	CategoryID uint
	ProfileID  uint
	PageNum    uint
	PageSize   uint
	Sorting    SortType
}

type GetProductsByCategoryIDPayload struct {
	CategoryName string
	Products     []ProductCard
}
