package dao

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type ProductCard struct {
	ID     uint   `db:"id"`
	Name   string `db:"product_name"`
	Price  uint   `db:"price"`
	ImgSrc string `db:"imgsrc"`
	Seller string `db:"seller"`
	Rating uint   `db:"rating"`
	Amount uint   `db:"count"`
}

func ConvertProductCardToModel(card ProductCard) models.ProductCard {
	return models.ProductCard{
		ID:     card.ID,
		Name:   card.Name,
		Price:  card.Price,
		ImgSrc: card.ImgSrc,
		Seller: card.Seller,
		Rating: card.Rating,
		Amount: card.Amount,
	}
}

func ConvertProductCardsFromTable(cards []ProductCard) []models.ProductCard {
	productCards := make([]models.ProductCard, 0, len(cards))
	for _, card := range cards {
		productCards = append(productCards, ConvertProductCardToModel(card))
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
	Amount      uint   `db:"count"`
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
		Amount:      product.Amount,
		Categories:  categories,
	}
}

func ConvertProductsFromTables(categories [][]string, products []Product) []models.Product {
	res := make([]models.Product, 0, len(products))
	for i, p := range products {
		res = append(res, ConvertProductFromTable(categories[i], p))
	}
	return res
}
