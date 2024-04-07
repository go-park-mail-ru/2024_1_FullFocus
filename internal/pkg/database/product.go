package database

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type Product struct {
	ID          uint   `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"product_description"`
	Price       uint   `json:"price"`
	ImgSrc      string `json:"imgsrc"`
	Seller      string `json:"seller"`
	Rating      uint   `json:"rating"`
}

func ConvertProductsFromTable(categories []string, product []Product) []models.Product {
	productModels := make([]models.Product, 0, len(product))
	for _, order := range product {
		productModels = append(productModels, models.Product{
			ID:          order.ID,
			ProductName: order.ProductName,
			Description: order.Description,
			Price:       order.Price,
			ImgSrc:      order.ImgSrc,
			Seller:      order.Seller,
			Rating:      order.Rating,
			Categories:  categories,
		})
	}
	return productModels
}
