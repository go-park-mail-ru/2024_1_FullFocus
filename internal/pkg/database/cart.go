package database

import (
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
)

type CartItemTable struct {
	PrID  uint `db:"product_id"`
	Count uint `db:"count"`
}

func ConvertTableToCartItem(t CartItemTable) models.CartItem {
	return models.CartItem{
		ProductID: t.PrID,
		Count:     t.Count,
	}
}

func ConvertTablesToCartItems(tt []CartItemTable) []models.CartItem {
	cartItems := make([]models.CartItem, 0)
	for _, t := range tt {
		cartItem := models.CartItem{
			ProductID: t.PrID,
			Count:     t.Count,
		}
		cartItems = append(cartItems, cartItem)
	}
	return cartItems
}
