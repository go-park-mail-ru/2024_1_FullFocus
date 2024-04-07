package models

type Product struct {
	ID          uint
	ProductName string
	Description string
	Price       uint
	ImgSrc      string
	Seller      string
	Rating      uint
	Categories  []string
}
