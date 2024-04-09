package models

type CartItem struct {
	ProductID uint
	Count     uint
}

type CartProduct struct {
	ProductID uint
	Name      string
	Price     uint
	Count     uint
	Img       string
	Cost      uint
}

type CartContent struct {
	Products   []CartProduct
	TotalCount uint
	TotalCost  uint
}
