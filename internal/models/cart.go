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
}
