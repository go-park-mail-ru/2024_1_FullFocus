package models

type OrderProduct struct {
	ProductName string
	Count       uint
	ImgSrc      string
}

type OrderItem struct {
	ProductID uint
	Count     uint
}

type CreateOrderInput struct {
	UserID   uint
	Items    []OrderItem
	FromCart bool
}
