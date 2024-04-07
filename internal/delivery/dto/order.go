package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type OrderProduct struct {
	ProductName string `json:"productName"`
	Count       uint   `json:"count"`
	ImgSrc      string `json:"imgSrc"`
}

func ConvertProductsToDTO(products []models.OrderProduct) []OrderProduct {
	orderProducts := make([]OrderProduct, 0, len(products))
	for _, product := range products {
		orderProducts = append(orderProducts, OrderProduct{
			ProductName: product.ProductName,
			Count:       product.Count,
			ImgSrc:      product.ImgSrc,
		})
	}
	return orderProducts
}

func ConvertOrderToModels(products []OrderProduct) []models.OrderProduct {
	orderProducts := make([]models.OrderProduct, 0, len(products))
	for _, product := range products {
		orderProducts = append(orderProducts, models.OrderProduct{
			ProductName: product.ProductName,
			Count:       product.Count,
			ImgSrc:      product.ImgSrc,
		})
	}
	return orderProducts
}

type OrderItem struct {
	ProductID uint `json:"productID"`
	Count     uint `json:"count"`
}

// Create

type CreateOrderInput struct {
	UserID   uint        `json:"userID"`
	Items    []OrderItem `json:"items"`
	FromCart bool        `json:"fromCart"`
}

type CreateOrderPayload struct {
	OrderID uint `json:"orderID"`
}

// Get

type GetOrderProductsInput struct {
	OrderID uint `json:"orderID"`
}

// Update

type CancelOrderInput struct {
	OrderID uint `json:"orderID"`
}
