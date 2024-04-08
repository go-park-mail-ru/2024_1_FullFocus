package database

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type ProductCard struct {
	ID     uint   `db:"id"`
	Name   string `db:"product_name"`
	Price  uint   `db:"price"`
	ImgSrc string `db:"imgsrc"`
	Seller string `db:"seller"`
	Rating uint   `db:"rating"`
}

func ConvertProductCardsFromTable(cards []ProductCard) []models.ProductCard {
	var productCards []models.ProductCard
	for _, card := range cards {
		productCards = append(productCards, models.ProductCard{
			ID:     card.ID,
			Name:   card.Name,
			Price:  card.Price,
			ImgSrc: card.ImgSrc,
			Seller: card.Seller,
			Rating: card.Rating,
		})
	}
	return productCards
}

type Product struct {
	ID          uint   `db:"id"`
	Name        string `db:"product_name"`
	Description string `db:"product_description"`
	Price       uint   `db:"price"`
	ImgSrc      string `db:"imgsrc"`
	Seller      string `db:"seller"`
	Rating      uint   `db:"rating"`
}

func ConvertProductFromTable(categories []string, product Product) models.Product {
	return models.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImgSrc:      product.ImgSrc,
		Seller:      product.Seller,
		Rating:      product.Rating,
		Categories:  categories,
	}
}
