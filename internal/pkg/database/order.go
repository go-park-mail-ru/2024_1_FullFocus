package database

import "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"

type OrderItem struct {
	OrderingID uint `db:"ordering_id"`
	ProductID  uint `db:"product_id"`
	Count      uint `db:"count"`
}

func ConvertOrderItemsToTables(orderingID uint, items []models.OrderItem) []OrderItem {
	orderItems := make([]OrderItem, 0, len(items))
	for _, item := range items {
		orderItems = append(orderItems, OrderItem{
			OrderingID: orderingID,
			ProductID:  item.ProductID,
			Count:      item.Count,
		})
	}
	return orderItems
}

type OrderProduct struct {
	ProductName string `db:"product_name"`
	Count       uint   `db:"count"`
	ImgSrc      string `db:"imgsrc"`
}

func ConvertOrderProductsToModels(items []OrderProduct) []models.OrderProduct {
	orderProducts := make([]models.OrderProduct, 0, len(items))
	for _, item := range items {
		orderProducts = append(orderProducts, models.OrderProduct{
			ProductName: item.ProductName,
			Count:       item.Count,
			ImgSrc:      item.ImgSrc,
		})
	}
	return orderProducts
}
