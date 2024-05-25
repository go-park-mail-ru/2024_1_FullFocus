package dao

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

type CartProductTable struct {
	PrID   uint   `db:"id"`
	Name   string `db:"product_name"`
	Price  uint   `db:"price"`
	Count  uint   `db:"count"`
	Img    string `db:"imgsrc"`
	OnSale bool   `db:"on_sale"`
}

func ConvertTablesToCartProducts(tt []CartProductTable) []models.CartProduct {
	cartProducts := make([]models.CartProduct, 0)
	for _, t := range tt {
		cartProduct := models.CartProduct{
			ProductID: t.PrID,
			Name:      t.Name,
			Price:     t.Price,
			Count:     t.Count,
			Img:       t.Img,
			OnSale:    t.OnSale,
		}
		cartProducts = append(cartProducts, cartProduct)
	}
	return cartProducts
}
