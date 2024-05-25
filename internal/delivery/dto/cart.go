package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type CartContent struct {
	Products   []CartProduct `json:"products"`
	TotalCount uint          `json:"totalCount"`
	TotalCost  uint          `json:"totalCost"`
}

func ConvertContentToDto(m models.CartContent) CartContent {
	return CartContent{
		Products:   ConvertProductsToDto(m.Products),
		TotalCount: m.TotalCount,
		TotalCost:  m.TotalCost,
	}
}

type CartProduct struct {
	ProductID    uint   `json:"productId"`
	Name         string `json:"name"`
	Price        uint   `json:"oldPrice"`
	Count        uint   `json:"count"`
	Img          string `json:"imgsrc"`
	Cost         uint   `json:"cost"`
	BenefitType  string `json:"benefitType"`
	BenefitValue uint   `json:"benefitValue"`
	NewPrice     uint   `json:"newPrice"`
}

func ConvertProductsToDto(mm []models.CartProduct) []CartProduct {
	cartProduct := make([]CartProduct, 0)
	for _, m := range mm {
		cartProduct = append(cartProduct, CartProduct{
			ProductID:    m.ProductID,
			Name:         m.Name,
			Price:        m.Price,
			Count:        m.Count,
			Img:          m.Img,
			Cost:         m.Cost,
			BenefitType:  m.BenefitType,
			BenefitValue: m.BenefitValue,
			NewPrice:     m.NewPrice,
		})
	}
	return cartProduct
}

type UpdateCartItemInput struct {
	ProductID uint `json:"productId"`
}

type UpdateCartItemPayload struct {
	Count uint `json:"count"`
}
