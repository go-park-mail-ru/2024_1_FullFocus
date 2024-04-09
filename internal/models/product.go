package models

type ProductCard struct {
	ID     uint
	Name   string
	Price  uint
	ImgSrc string
	Seller string
	Rating uint
	InCart bool
}

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       uint
	ImgSrc      string
	Seller      string
	Rating      uint
	InCart      bool
	Categories  []string
}

type GetAllProductsInput struct {
	ProfileID uint
	PageNum   uint
	PageSize  uint
}

type GetProductsByCategoryIDInput struct {
	CategoryID uint
	ProfileID  uint
	PageNum    uint
	PageSize   uint
}
