package models

type PromoProduct struct {
	ProductData  Product
	BenefitType  string
	BenefitValue uint
	NewPrice     uint
}

type PromoProductCard struct {
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

func ConvertPromoProductToCard(data PromoProduct) PromoProductCard {
	return PromoProductCard{
		ProductData:  ConvertProductToCard(data.ProductData),
		BenefitType:  data.BenefitType,
		BenefitValue: data.BenefitValue,
		NewPrice:     data.NewPrice,
	}
}
