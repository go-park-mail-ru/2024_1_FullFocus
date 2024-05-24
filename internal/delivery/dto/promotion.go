package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type PromoProductCard struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	OldPrice     uint   `json:"oldPrice"`
	ImgSrc       string `json:"imgSrc"`
	Seller       string `json:"seller"`
	Rating       uint   `json:"rating"`
	Amount       uint   `json:"amount"`
	BenefitType  string `json:"benefitType"`
	BenefitValue uint   `json:"benefitValue"`
	NewPrice     uint   `json:"newPrice"`
}

func ConvertPromoProductCardToDTO(m models.PromoProductCard) PromoProductCard {
	return PromoProductCard{
		ID:           m.ProductData.ID,
		Name:         m.ProductData.Name,
		OldPrice:     m.ProductData.Price,
		ImgSrc:       m.ProductData.ImgSrc,
		Seller:       m.ProductData.Seller,
		Rating:       m.ProductData.Rating,
		Amount:       m.ProductData.Amount,
		BenefitType:  m.BenefitType,
		BenefitValue: m.BenefitValue,
		NewPrice:     m.NewPrice,
	}
}

func ConvertPromoProductCardsToDTOs(mm []models.PromoProductCard) []PromoProductCard {
	data := make([]PromoProductCard, 0, len(mm))
	for _, m := range mm {
		data = append(data, ConvertPromoProductCardToDTO(m))
	}
	return data
}

type PromoProduct struct {
	ID           uint     `json:"id"`
	Name         string   `json:"name"`
	OldPrice     uint     `json:"oldPrice"`
	Description  string   `json:"description"`
	ImgSrc       string   `json:"imgSrc"`
	Seller       string   `json:"seller"`
	Rating       uint     `json:"rating"`
	Amount       uint     `json:"amount"`
	Categories   []string `json:"categories"`
	BenefitType  string   `json:"benefitType"`
	BenefitValue uint     `json:"benefitValue"`
	NewPrice     uint     `json:"newPrice"`
}

func ConvertPromoProductToDTO(m models.PromoProduct) PromoProduct {
	return PromoProduct{
		ID:           m.ProductData.ID,
		Name:         m.ProductData.Name,
		OldPrice:     m.ProductData.Price,
		Description:  m.ProductData.Description,
		ImgSrc:       m.ProductData.Description,
		Seller:       m.ProductData.Seller,
		Rating:       m.ProductData.Rating,
		Amount:       m.ProductData.Amount,
		Categories:   m.ProductData.Categories,
		BenefitType:  m.BenefitType,
		BenefitValue: m.BenefitValue,
		NewPrice:     m.NewPrice,
	}
}

type PromoData struct {
	ProductID    uint   `json:"productID"`
	BenefitType  string `json:"benefitType"`
	BenefitValue uint   `json:"benefitValue"`
}

func ConvertPromoDataToModel(d PromoData) models.PromoData {
	return models.PromoData{
		ProductID:    d.ProductID,
		BenefitType:  d.BenefitType,
		BenefitValue: d.BenefitValue,
	}
}

type DeletePromoProductInput struct {
	ProductID uint `json:"productID"`
}
