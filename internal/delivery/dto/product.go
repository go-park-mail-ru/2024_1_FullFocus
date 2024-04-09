package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

// TODO: inCart

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

// TODO: inCart
// TODO: how to handle ErrNoRows

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
