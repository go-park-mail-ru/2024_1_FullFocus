package models

type ProductCard struct {
	ID     uint
	Name   string
	Price  uint
	ImgSrc string
	Seller string
	Rating uint
}

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       uint
	ImgSrc      string
	Seller      string
	Rating      uint
	Categories  []string
}
