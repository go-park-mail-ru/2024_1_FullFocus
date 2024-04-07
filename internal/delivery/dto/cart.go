package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type CartProduct struct {
	ProductID uint   `json:"productId"`
	Name      string `json:"name"`
	Price     uint   `json:"price"`
	Count     uint   `json:"count"`
	Img       string `json:"imgsrc"`
}

func ConvertProductsToDto(mm []models.CartProduct) []CartProduct {
	cartProduct := make([]CartProduct, 0)
	for _, m := range mm {
		cartProduct = append(cartProduct, CartProduct{
			ProductID: m.ProductID,
			Name:      m.Name,
			Price:     m.Price,
			Count:     m.Count,
			Img:       m.Img,
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
