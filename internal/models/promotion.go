package models

type PromoProduct struct {
	ProductData  ProductCard
	BenefitType  string
	BenefitValue uint
	NewPrice     uint
}

type PromoData struct {
	ProductID    uint
	BenefitType  string
	BenefitValue uint
}

type CachePromoProduct struct {
	Product PromoProduct
	Empty   bool
}
