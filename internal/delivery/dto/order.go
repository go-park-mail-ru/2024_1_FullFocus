package dto

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type OrderProduct struct {
	ProductName string `json:"productName"`
	Price       uint   `json:"price"`
	Count       uint   `json:"count"`
	ImgSrc      string `json:"imgSrc"`
}

func ConvertOrderProductsToDTO(products []models.OrderProduct) []OrderProduct {
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

type Order struct {
	ID     uint   `json:"id"`
	Sum    uint   `json:"sum"`
	Status string `json:"status"`
}

func ConvertOrdersToDTO(orders []models.Order) []Order {
	orderProducts := make([]Order, 0, len(orders))
	for _, order := range orders {
		orderProducts = append(orderProducts, Order{
			ID:     order.ID,
			Sum:    order.Sum,
			Status: order.Status,
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

type GetOrderInput struct {
	OrderID uint `json:"orderID"`
}

type GetOrderPayload struct {
	Products []OrderProduct `json:"products"`
	Sum      uint           `json:"sum"`
	Status   string         `json:"status"`
}

type GetAllOrdersPayload struct {
	Orders []Order `json:"orders"`
}

// Update

type CancelOrderInput struct {
	OrderID uint `json:"orderID"`
}
