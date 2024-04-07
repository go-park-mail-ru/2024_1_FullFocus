package models

type OrderProduct struct {
	ProductName string
	Price       uint
	Count       uint
	ImgSrc      string
}

type OrderItem struct {
	ProductID uint
	Count     uint
}

type Order struct {
	ID     uint
	Sum    uint
	Status string
}

// Create

type CreateOrderInput struct {
	UserID   uint
	Items    []OrderItem
	FromCart bool
}

// Get

type GetOrderPayload struct {
	Products []OrderProduct
	Sum      uint
	Status   string
}
