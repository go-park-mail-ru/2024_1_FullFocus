package models

type PromoProduct struct {
	ProductData  ProductCard
	BenefitType  string
	BenefitValue uint
}

type PromoData struct {
	ProductID    uint
	BenefitType  string
	BenefitValue uint
}
