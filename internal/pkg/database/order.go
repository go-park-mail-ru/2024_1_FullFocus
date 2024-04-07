package database

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type Order struct {
	ID     uint   `json:"id"`
	Sum    uint   `json:"sum"`
	Status string `json:"status"`
}

func ConvertOrdersFromTable(orders []Order) []models.Order {
	orderModels := make([]models.Order, 0, len(orders))
	for _, order := range orders {
		orderModels = append(orderModels, models.Order{
			ID:     order.ID,
			Sum:    order.Sum,
			Status: order.Status,
		})
	}
	return orderModels
}

type OrderItem struct {
	OrderingID uint `db:"ordering_id"`
	ProductID  uint `db:"product_id"`
	Count      uint `db:"count"`
}

func ConvertOrderItemsToTables(orderID uint, items []models.OrderItem) []OrderItem {
	orderItems := make([]OrderItem, 0, len(items))
	for _, item := range items {
		orderItems = append(orderItems, OrderItem{
			OrderingID: orderID,
			ProductID:  item.ProductID,
			Count:      item.Count,
		})
	}
	return orderItems
}

type OrderProduct struct {
	ProductName string `db:"product_name"`
	Price       uint   `db:"price"`
	Count       uint   `db:"count"`
	ImgSrc      string `db:"imgsrc"`
}

func ConvertOrderProductsToModels(items []OrderProduct) []models.OrderProduct {
	orderProducts := make([]models.OrderProduct, 0, len(items))
	for _, item := range items {
		orderProducts = append(orderProducts, models.OrderProduct{
			Price:       item.Price,
			ProductName: item.ProductName,
			Count:       item.Count,
			ImgSrc:      item.ImgSrc,
		})
	}
	return orderProducts
}
