package models

type CartItem struct {
	ProductID uint
	Count     uint
	Price     uint
	OnSale    bool
}

type CartProduct struct {
	ProductID    uint
	Name         string
	Price        uint
	Count        uint
	Img          string
	Cost         uint
	OnSale       bool
	BenefitType  string
	BenefitValue uint
	NewPrice     uint
}

type CartContent struct {
	Products   []CartProduct
	TotalCount uint
	TotalCost  uint
}
