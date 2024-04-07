package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type Product struct {
	ID          uint     `json:"id"`
	ProductName string   `json:"name"`
	Description string   `json:"description"`
	Price       uint     `json:"price"`
	ImgSrc      string   `json:"imgSrc"`
	Seller      string   `json:"seller"`
	Rating      uint     `json:"rating"`
	Categories  []string `json:"categories"`
}

func ConvertProductsToDTO(products []Product) []models.Product {
	productModels := make([]models.Product, 0, len(products))
	for _, product := range products {
		productModels = append(productModels, models.Product{
			ID:          product.ID,
			ProductName: product.ProductName,
			Description: product.Description,
			Price:       product.Price,
			ImgSrc:      product.ImgSrc,
			Seller:      product.Seller,
			Rating:      product.Rating,
			Categories:  product.Categories,
		})
	}
	return productModels
}
