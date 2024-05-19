package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type PromoProduct struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	OldPrice     uint   `json:"oldPrice"`
	ImgSrc       string `json:"imgSrc"`
	Seller       string `json:"seller"`
	Rating       uint   `json:"rating"`
	BenefitType  string `json:"benefitType"`
	BenefitValue uint   `json:"benefitValue"`
	NewPrice     uint   `json:"newPrice"`
}

func ConvertPromoProductToDTO(m models.PromoProduct) PromoProduct {
	return PromoProduct{
		ID:           m.ProductData.ID,
		Name:         m.ProductData.Name,
		OldPrice:     m.ProductData.Price,
		ImgSrc:       m.ProductData.ImgSrc,
		Seller:       m.ProductData.Seller,
		Rating:       m.ProductData.Rating,
		BenefitType:  m.BenefitType,
		BenefitValue: m.BenefitValue,
		NewPrice:     m.NewPrice,
	}
}

func ConvertPromoProductsToDTOs(mm []models.PromoProduct) []PromoProduct {
	data := make([]PromoProduct, 0, len(mm))
	for _, m := range mm {
		data = append(data, ConvertPromoProductToDTO(m))
	}
	return data
}
