package dto

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type ProductCard struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Price  uint   `json:"price"`
	ImgSrc string `json:"imgSrc"`
	Seller string `json:"seller"`
	Rating uint   `json:"rating"`
	InCart bool   `json:"inCart"`
}

func ConvertProductCardsToDTO(cards []models.ProductCard) []ProductCard {
	var productCards []ProductCard
	for _, card := range cards {
		productCards = append(productCards, ProductCard{
			ID:     card.ID,
			Name:   card.Name,
			Price:  card.Price,
			ImgSrc: card.ImgSrc,
			Seller: card.Seller,
			Rating: card.Rating,
			InCart: card.InCart,
		})
	}
	return productCards
}

type GetAllProductsPayload struct {
	ProductCards []ProductCard `json:"productCards"`
}

type Product struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       uint     `json:"price"`
	ImgSrc      string   `json:"imgSrc"`
	Seller      string   `json:"seller"`
	Rating      uint     `json:"rating"`
	InCart      bool     `json:"inCart"`
	Categories  []string `json:"categories"`
}

func ConvertProductToDTO(product models.Product) Product {
	return Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImgSrc:      product.ImgSrc,
		Seller:      product.Seller,
		Rating:      product.Rating,
		InCart:      product.InCart,
		Categories:  product.Categories,
	}
}

type GetProductsByCategoryIDPayload struct {
	CategoryName string        `json:"categoryName"`
	Products     []ProductCard `json:"productCards"`
}

func ConvertProductsByCategoryIDPayload(payload models.GetProductsByCategoryIDPayload) GetProductsByCategoryIDPayload {
	return GetProductsByCategoryIDPayload{
		CategoryName: payload.CategoryName,
		Products:     ConvertProductCardsToDTO(payload.Products),
	}
}

type GetProductsByQueryPayload struct {
	Query    string        `json:"query"`
	Products []ProductCard `json:"productCards"`
}
