package dto

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type OrderProduct struct {
	ID          uint   `json:"id"`
	ProductName string `json:"productName"`
	Price       uint   `json:"price"`
	Count       uint   `json:"count"`
	ImgSrc      string `json:"imgSrc"`
}

func ConvertOrderProductsToDTO(products []models.OrderProduct) []OrderProduct {
	orderProducts := make([]OrderProduct, 0, len(products))
	for _, product := range products {
		orderProducts = append(orderProducts, OrderProduct{
			ID:          product.ID,
			ProductName: product.ProductName,
			Count:       product.Count,
			ImgSrc:      product.ImgSrc,
		})
	}
	return orderProducts
}

type Order struct {
	ID         uint   `json:"id"`
	Sum        uint   `json:"sum"`
	Status     string `json:"status"`
	ItemsCount uint   `json:"itemsCount"`
	CreatedAt  string `json:"createdAt"`
}

func ConvertOrdersToDTO(orders []models.Order) []Order {
	orderProducts := make([]Order, 0, len(orders))
	for _, order := range orders {
		orderProducts = append(orderProducts, Order{
			ID:         order.ID,
			Sum:        order.Sum,
			Status:     order.Status,
			ItemsCount: order.ItemsCount,
			CreatedAt:  order.CreatedAt,
		})
	}
	return orderProducts
}

// Create

type OrderItem struct {
	ProductID uint `json:"productID"`
	Count     uint `json:"count"`
}

type CreateOrderInput struct {
	Items    []OrderItem `json:"items"`
	FromCart bool        `json:"fromCart"`
}

func ConvertCreateOrderInputToModel(userID uint, input CreateOrderInput) models.CreateOrderInput {
	createInput := models.CreateOrderInput{
		UserID:   userID,
		FromCart: input.FromCart,
	}
	for _, item := range input.Items {
		createInput.Items = append(createInput.Items, models.OrderItem{
			ProductID: item.ProductID,
			Count:     item.Count,
		})
	}
	return createInput
}

type CreateOrderPayload struct {
	OrderID uint `json:"orderID"`
}

// Get

type GetOrderPayload struct {
	Products   []OrderProduct `json:"products"`
	Sum        uint           `json:"sum"`
	Status     string         `json:"status"`
	ItemsCount uint           `json:"itemsCount"`
	CreatedAt  string         `json:"createdAt"`
}

type GetAllOrdersPayload struct {
	Orders []Order `json:"orders"`
}

// Update

type CancelOrderInput struct {
	OrderID uint `json:"orderID"`
}
